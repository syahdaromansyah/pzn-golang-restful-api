package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/middleware"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
)

var configPaths []string

func main() {
	parseFlag()

	appConfig := config.NewAppConfig(configPaths)

	pool := config.NewPgxPool(appConfig)
	defer pool.Close()

	logFile := createOrOpenLogFile(appConfig)
	defer logFile.Close()

	logger := setupLogger(appConfig, logFile)

	router := httprouter.New()

	routeConfig := InitializeController(pool, logger, router)
	routeConfig.Setup()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", appConfig.Server.Host, appConfig.Server.Port),
		Handler: setupMiddleware(appConfig, router, logger),
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	shutdownDoneChan := make(chan bool, 1)

	go func() {
		signal := <-sigChan

		logger.Warnf("received signal %s", signal)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)

		logger.Warn("the server is shutting down...")

		if err != nil {
			logger.WithError(err).Error("the server failed to shut down gracefully")
		} else {
			logger.Info("the server has shut down gracefully")
		}

		shutdownDoneChan <- true
	}()

	logger.Infof("the server is listening on %s", server.Addr)

	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		logger.WithError(err).Error("the server has been closed")

		<-shutdownDoneChan
	} else {
		logger.WithError(err).Error("the server failed to listen and serve")
	}
}

func parseFlag() {
	if !flag.Parsed() {
		var configPathsFlag string

		flag.StringVar(&configPathsFlag, "configPaths", "./", "Multiple config paths separated with commas")

		flag.Parse()

		configPaths = strings.Split(configPathsFlag, ",")
	}
}

func createOrOpenLogFile(appConfig *config.AppConfig) *os.File {
	logFile, err := os.OpenFile(
		appConfig.Log.FilePath,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	helper.LogStdPanicIfError(err)
	return logFile
}

func setupLogger(appConfig *config.AppConfig, logFile *os.File) *logrus.Logger {
	logger := config.NewLogrus(appConfig)

	if appConfig.Log.Output == "file" {
		logger.SetOutput(logFile)
	}

	return logger
}

func setupMiddleware(appConfig *config.AppConfig, router *httprouter.Router, logger *logrus.Logger) middleware.HttpMiddleware {
	authMiddleware := middleware.NewHttpAuthMiddleware(appConfig, router)
	panicMiddleware := middleware.NewHttpPanicMiddleware(logger, authMiddleware)

	return panicMiddleware
}

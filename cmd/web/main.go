package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/middleware"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
)

func createOrOpenLogFile(vp *viper.Viper) *os.File {
	logFile, err := os.OpenFile(
		vp.GetString("log.filepath"),
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0644,
	)
	helper.LogStdPanicIfError(err)
	return logFile
}

func setupLogger(vp *viper.Viper, logFile *os.File) *logrus.Logger {
	logger := config.NewLogrus(vp)

	if vp.GetString("log.output") == "file" {
		logger.SetOutput(logFile)
	}

	return logger
}

func setupMiddleware(vp *viper.Viper, router *httprouter.Router, logger *logrus.Logger) middleware.HttpMiddleware {
	authMiddleware := middleware.NewHttpAuthMiddleware(vp, router)
	panicMiddleware := middleware.NewHttpPanicMiddleware(logger, authMiddleware)

	return panicMiddleware
}

func main() {
	vp := config.NewViper([]string{".", "./.."})

	pool := config.NewPgxPool(vp)
	defer pool.Close()

	logFile := createOrOpenLogFile(vp)
	defer logFile.Close()

	logger := setupLogger(vp, logFile)

	router := httprouter.New()

	routeConfig := InitializeController(vp, pool, logger, router)
	routeConfig.Setup()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", vp.GetString("server.host"), vp.GetInt("server.port")),
		Handler: setupMiddleware(vp, router, logger),
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

	<-shutdownDoneChan

	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("the server has been shut down")
	} else {
		logger.WithError(err).Error("the server failed to listen and serve")
	}
}

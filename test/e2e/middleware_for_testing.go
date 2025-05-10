package e2e

import (
	"github.com/julienschmidt/httprouter"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/middleware"
)

func setupMiddleware(appConfig *config.AppConfig) middleware.HttpMiddleware {
	pool := config.NewPgxPool(appConfig)
	logger := config.NewLogrus(appConfig)
	router := httprouter.New()

	routeConfig := InitializeControllerForTesting(appConfig, pool, logger, router)
	routeConfig.Setup()

	return middleware.NewHttpPanicMiddleware(
		logger,
		middleware.NewHttpAuthMiddleware(
			appConfig,
			router,
		),
	)
}

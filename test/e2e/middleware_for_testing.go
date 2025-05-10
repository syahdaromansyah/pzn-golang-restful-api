package e2e

import (
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/middleware"
)

func setupMiddleware(vp *viper.Viper) middleware.HttpMiddleware {
	pool := config.NewPgxPool(vp)
	logger := config.NewLogrus(vp)
	router := httprouter.New()

	routeConfig := InitializeControllerForTesting(vp, pool, logger, router)
	routeConfig.Setup()

	return middleware.NewHttpPanicMiddleware(
		logger,
		middleware.NewHttpAuthMiddleware(
			vp,
			router,
		),
	)
}

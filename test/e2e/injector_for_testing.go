//go:build wireinject
// +build wireinject

package e2e

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/route"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/db"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/repository"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/security"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/usecase"
)

var repositorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
)

var useCaseSet = wire.NewSet(
	usecase.NewCategoryUseCaseImpl,
)

var controllerSet = wire.NewSet(
	http.NewCategoryControllerImpl,
)

func InitializeControllerForTesting(vp *viper.Viper, database db.PgxPool, logger *logrus.Logger, router *httprouter.Router) route.RouteConfig {
	wire.Build(
		security.NewIdGenImpl,
		security.NewValidationImpl,
		repositorySet,
		useCaseSet,
		controllerSet,
		route.NewRouteConfigHttpRouter,
	)

	return nil
}

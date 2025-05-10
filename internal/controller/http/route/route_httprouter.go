package route

import (
	go_http "net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http"
)

type RouteConfigHttpRouter struct {
	Router             *httprouter.Router
	CategoryController http.CategoryController
}

func NewRouteConfigHttpRouter(router *httprouter.Router, categoryController http.CategoryController) RouteConfig {
	return &RouteConfigHttpRouter{
		Router:             router,
		CategoryController: categoryController,
	}
}

func (r *RouteConfigHttpRouter) Setup() {
	// Category Endpoints
	r.Router.GET("/api/categories", r.CategoryController.FindAll)
	r.Router.GET("/api/categories/:categoryId", r.CategoryController.FindById)
	r.Router.POST("/api/categories", r.CategoryController.Create)
	r.Router.PUT("/api/categories/:categoryId", r.CategoryController.Update)
	r.Router.DELETE("/api/categories/:categoryId", r.CategoryController.Delete)

	// Panic Endpoint
	r.Router.PanicHandler = func(w go_http.ResponseWriter, r *go_http.Request, err any) {
		panic(err)
	}
}

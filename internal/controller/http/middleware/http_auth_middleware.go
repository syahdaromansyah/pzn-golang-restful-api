package middleware

import (
	"errors"
	"net/http"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"
)

type httpAuthMiddleware struct {
	AppConfig *config.AppConfig
	Handler   http.Handler
}

func NewHttpAuthMiddleware(appConfig *config.AppConfig, handler http.Handler) HttpMiddleware {
	return &httpAuthMiddleware{
		AppConfig: appConfig,
		Handler:   handler,
	}
}

func (m *httpAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.AppConfig.Server.ApiKey == r.Header.Get("X-API-Key") {
		m.Handler.ServeHTTP(w, r)
	} else {
		panic(exception.NewErrorClientRequest(errors.New("unauthorized"), http.StatusUnauthorized, "unauthorized"))
	}
}

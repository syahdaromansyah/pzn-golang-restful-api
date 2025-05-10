package middleware

import (
	"errors"
	"net/http"

	"github.com/spf13/viper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"
)

type httpAuthMiddleware struct {
	Viper   *viper.Viper
	Handler http.Handler
}

func NewHttpAuthMiddleware(vp *viper.Viper, handler http.Handler) HttpMiddleware {
	return &httpAuthMiddleware{
		Viper:   vp,
		Handler: handler,
	}
}

func (m *httpAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Viper.GetString("server.api_key") == r.Header.Get("X-API-Key") {
		m.Handler.ServeHTTP(w, r)
	} else {
		panic(exception.NewErrorClientRequest(errors.New("unauthorized"), http.StatusUnauthorized, "unauthorized"))
	}
}

package middleware

import "net/http"

type HttpMiddleware interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

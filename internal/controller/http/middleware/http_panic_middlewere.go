package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
)

type httpPanicMiddleware struct {
	Logger  *logrus.Logger
	Handler http.Handler
}

func NewHttpPanicMiddleware(logger *logrus.Logger, handler http.Handler) HttpMiddleware {
	return &httpPanicMiddleware{
		Logger:  logger,
		Handler: handler,
	}
}

func (m *httpPanicMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer m.recoverError(w, r)
	m.Handler.ServeHTTP(w, r)
}

func (m *httpPanicMiddleware) recoverError(w http.ResponseWriter, r *http.Request) {
	errRecover := recover()

	if errRecover != nil {
		w.Header().Set("content-type", "application/json")

		if m.clientError(w, r, errRecover) {
			return
		}

		m.internalServerError(w, r, errRecover)
	}
}

func (m *httpPanicMiddleware) clientError(w http.ResponseWriter, _ *http.Request, err any) bool {
	if exception, ok := err.(*exception.ErrorClientRequest); ok {
		if exception.StatusCode == http.StatusBadRequest {
			w.WriteHeader(http.StatusBadRequest)

			webResponse := &model.WebResponseMessage{
				Code:    http.StatusBadRequest,
				Status:  "BAD REQUEST",
				Message: exception.GetDetailError(),
			}

			helper.WriteToResponseBody(w, webResponse)
		}

		if exception.StatusCode == http.StatusUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := &model.WebResponseMessage{
				Code:    http.StatusUnauthorized,
				Status:  "UNAUTHORIZED",
				Message: exception.GetDetailError(),
			}

			helper.WriteToResponseBody(w, webResponse)
		}

		if exception.StatusCode == http.StatusNotFound {
			w.WriteHeader(http.StatusNotFound)

			webResponse := &model.WebResponseMessage{
				Code:    http.StatusNotFound,
				Status:  "NOT FOUND",
				Message: exception.GetDetailError(),
			}

			helper.WriteToResponseBody(w, webResponse)
		}

		return true
	}

	if exception, ok := err.(validator.ValidationErrors); ok {
		w.WriteHeader(http.StatusBadRequest)

		webResponse := &model.WebResponseMessage{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)

		return true
	}

	return false
}

func (m *httpPanicMiddleware) internalServerError(w http.ResponseWriter, _ *http.Request, err any) {
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := &model.WebResponseMessage{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL SERVER ERROR",
		Message: "something went wrong",
	}

	helper.WriteToResponseBody(w, webResponse)

	exception, ok := err.(*exception.ErrorInternalServer)

	if ok {
		m.Logger.WithError(exception).WithField("detail_error", exception.DetailError()).Error("internal server error")
	} else {
		m.Logger.WithError(exception).Error("internal server error")
	}
}

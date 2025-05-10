package test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/middleware"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
)

var logger = config.NewLogrus(
	config.NewAppConfig([]string{"./../../../../.."}),
)

type panicHandler struct {
	Error error
}

func (h *panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic(h.Error)
}

func Test400Handler(t *testing.T) {
	t.Run("Validation Error", func(t *testing.T) {
		errorValidation := validator.ValidationErrors{}

		recorder := httptest.NewRecorder()

		// ---SUT (Subject Under Test)
		middleware.NewHttpPanicMiddleware(logger, &panicHandler{
			Error: errorValidation,
		}).ServeHTTP(recorder, nil)
		// ---------------------------

		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusBadRequest, recorderResponse.StatusCode)

		requestBodyBytes, err := io.ReadAll(recorderResponse.Body)
		helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(requestBodyBytes, webResponse)
		helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD REQUEST", webResponse.Status)
	})

	t.Run("Other Bad Request Error", func(t *testing.T) {
		errorOther400 := exception.NewErrorClientRequest(errors.New("other 400 error request"), http.StatusBadRequest, "other 400 error request")

		recorder := httptest.NewRecorder()

		// ---SUT (Subject Under Test)
		middleware.NewHttpPanicMiddleware(logger, &panicHandler{
			Error: errorOther400,
		}).ServeHTTP(recorder, nil)
		// ---------------------------

		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusBadRequest, recorderResponse.StatusCode)

		requestBodyBytes, err := io.ReadAll(recorderResponse.Body)
		helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(requestBodyBytes, webResponse)
		helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD REQUEST", webResponse.Status)
		assert.Equal(t, "other 400 error request", webResponse.Data)
	})
}

func Test401Handler(t *testing.T) {
	error401 := exception.NewErrorClientRequest(errors.New("401 error request"), http.StatusUnauthorized, "401 error request")

	recorder := httptest.NewRecorder()

	// ---SUT (Subject Under Test)
	middleware.NewHttpPanicMiddleware(logger, &panicHandler{
		Error: error401,
	}).ServeHTTP(recorder, nil)
	// ---------------------------

	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusUnauthorized, recorderResponse.StatusCode)

	requestBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[string])

	err = json.Unmarshal(requestBodyBytes, webResponse)
	helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusUnauthorized, webResponse.Code)
	assert.Equal(t, "UNAUTHORIZED", webResponse.Status)
	assert.Equal(t, "401 error request", webResponse.Data)
}

func Test404Handler(t *testing.T) {
	error404 := exception.NewErrorClientRequest(errors.New("404 error request"), http.StatusNotFound, "404 error request")

	recorder := httptest.NewRecorder()

	// ---SUT (Subject Under Test)
	middleware.NewHttpPanicMiddleware(logger, &panicHandler{
		Error: error404,
	}).ServeHTTP(recorder, nil)
	// ---------------------------

	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusNotFound, recorderResponse.StatusCode)

	requestBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[string])

	err = json.Unmarshal(requestBodyBytes, webResponse)
	helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusNotFound, webResponse.Code)
	assert.Equal(t, "NOT FOUND", webResponse.Status)
	assert.Equal(t, "404 error request", webResponse.Data)
}

func Test500Handler(t *testing.T) {
	error500 := exception.NewErrorInternalServer(errors.New("internal server error"), "internal server error")

	recorder := httptest.NewRecorder()

	// ---SUT (Subject Under Test)
	middleware.NewHttpPanicMiddleware(logger, &panicHandler{
		Error: error500,
	}).ServeHTTP(recorder, nil)
	// ---------------------------

	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusInternalServerError, recorderResponse.StatusCode)

	requestBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[struct{}])

	err = json.Unmarshal(requestBodyBytes, webResponse)
	helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusInternalServerError, webResponse.Code)
	assert.Equal(t, "INTERNAL SERVER ERROR", webResponse.Status)
	assert.Equal(t, struct{}{}, webResponse.Data)
}

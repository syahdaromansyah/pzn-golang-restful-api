package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http/middleware"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
)

var vp = config.NewViper([]string{"./../../../../.."})

type controllerHandler struct{}

func (h *controllerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "response from controllerHandler")
}

func TestFailed(t *testing.T) {
	testRequest := httptest.NewRequest("", "/", nil)
	recorder := httptest.NewRecorder()

	assert.Panics(t, func() {
		// ---SUT (Subject Under Test)
		middleware.NewHttpAuthMiddleware(vp, new(controllerHandler)).ServeHTTP(recorder, testRequest)
		// ---------------------------
	})
}

func TestSuccess(t *testing.T) {
	testRequest := httptest.NewRequest("", "/", nil)
	testRequest.Header.Set("X-API-Key", "test_key")

	recorder := httptest.NewRecorder()

	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		middleware.NewHttpAuthMiddleware(vp, new(controllerHandler)).ServeHTTP(recorder, testRequest)
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.LogStdPanicIfError(err)

	assert.Equal(t, "response from controllerHandler", string(responseBodyBytes))
}

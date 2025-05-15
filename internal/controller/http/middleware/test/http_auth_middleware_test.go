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

func setupAppTestConfig() *config.AppConfig {
	return &config.AppConfig{
		Server: &config.Server{
			ApiKey: "test_key",
		},
	}
}

var appTestConfig = setupAppTestConfig()

type controllerHandler struct{}

func (h *controllerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "response from controllerHandler")
}

func TestFailed(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest("", "/", nil)
	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.Panics(t, func() {
		// ---SUT (Subject Under Test)
		middleware.NewHttpAuthMiddleware(appTestConfig, new(controllerHandler)).ServeHTTP(recorder, testRequest)
		// ---------------------------
	})
}

func TestSuccess(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest("", "/", nil)
	testRequest.Header.Set("X-API-Key", "test_key")

	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		middleware.NewHttpAuthMiddleware(appTestConfig, new(controllerHandler)).ServeHTTP(recorder, testRequest)
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.LogStdPanicIfError(err)

	assert.Equal(t, "response from controllerHandler", string(responseBodyBytes))
}

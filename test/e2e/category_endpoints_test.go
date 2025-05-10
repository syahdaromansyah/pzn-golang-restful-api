package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	internal_helper "github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
	"github.com/syahdaromansyah/pzn-golang-restful-api/test/helper"
)

var baseUrl = "http://localhost:3000"
var vp = config.NewViper([]string{"./../.."})
var categoriesDbTableHelper = helper.NewCategoriesDbTable(vp)

func TestUnauthorized(t *testing.T) {
	// Arrange

	// Assume it is the category FindAll request
	testRequest := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/categories", baseUrl), nil)

	recorder := httptest.NewRecorder()

	middlewareTesting := setupMiddleware(vp)

	// Action
	middlewareTesting.ServeHTTP(recorder, testRequest)

	// Assert
	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusUnauthorized, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	internal_helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[string])

	err = json.Unmarshal(responseBodyBytes, webResponse)
	internal_helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusUnauthorized, webResponse.Code)
	assert.Equal(t, "UNAUTHORIZED", webResponse.Status)
	assert.Equal(t, "unauthorized", webResponse.Data)
}

func TestCreateFailed(t *testing.T) {
	t.Run("400 - Malformed Request Body", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		requestBody := strings.NewReader("")

		testRequest := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/categories", baseUrl), requestBody)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusBadRequest, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD REQUEST", webResponse.Status)
	})

	t.Run("400 - Field Name - Chars Min", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		requestBody := strings.NewReader(`{"name":"F"}`)

		testRequest := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/categories", baseUrl), requestBody)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusBadRequest, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD REQUEST", webResponse.Status)
	})
}

func TestCreateSuccess(t *testing.T) {
	// Arrange
	defer categoriesDbTableHelper.DeleteAll()

	requestBody := strings.NewReader(`{"name":"Fashions"}`)

	testRequest := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/categories", baseUrl), requestBody)

	testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

	recorder := httptest.NewRecorder()

	middlewareTesting := setupMiddleware(vp)

	// Action
	middlewareTesting.ServeHTTP(recorder, testRequest)

	// Assert
	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusCreated, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	internal_helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[*model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, webResponse)
	internal_helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusCreated, webResponse.Code)
	assert.Equal(t, "CREATED", webResponse.Status)
	assert.Equal(t, 36, len(webResponse.Data.Id))
	assert.Equal(t, "Fashions", webResponse.Data.Name)
}

func TestUpdateFailed(t *testing.T) {
	t.Run("400 - Malformed Request Body", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		requestBody := strings.NewReader("")

		testRequest := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), requestBody)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusBadRequest, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD REQUEST", webResponse.Status)
	})

	t.Run("404 - Category is Not Found", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		requestBody := strings.NewReader(`{"name":"Electronics"}`)

		testRequest := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), requestBody)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusNotFound, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusNotFound, webResponse.Code)
		assert.Equal(t, "NOT FOUND", webResponse.Status)
		assert.Equal(t, "category is not found", webResponse.Data)
	})

	t.Run("400 - Field Name - Chars Min", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		requestBody := strings.NewReader(`{"name":"E"}`)

		testRequest := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), requestBody)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusBadRequest, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD REQUEST", webResponse.Status)
	})
}

func TestUpdateSuccess(t *testing.T) {
	// Arrange
	defer categoriesDbTableHelper.DeleteAll()

	// Insert dummy data to DB
	categoriesDbTableHelper.Add(&entity.Category{
		Id:   "CAT-1",
		Name: "Tools",
	})
	// ------------------------

	requestBody := strings.NewReader(`{"name":"Electronics"}`)

	testRequest := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), requestBody)

	testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

	recorder := httptest.NewRecorder()

	middlewareTesting := setupMiddleware(vp)

	// Action
	middlewareTesting.ServeHTTP(recorder, testRequest)

	// Assert
	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	internal_helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[*model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, webResponse)
	internal_helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, "CAT-1", webResponse.Data.Id)
	assert.Equal(t, "Electronics", webResponse.Data.Name)
}

func TestDeleteFailed(t *testing.T) {
	t.Run("404 - Category is Not Found", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		testRequest := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), nil)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusNotFound, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusNotFound, webResponse.Code)
		assert.Equal(t, "NOT FOUND", webResponse.Status)
		assert.Equal(t, "category is not found", webResponse.Data)
	})
}

func TestDeleteSuccess(t *testing.T) {
	// Arrange
	defer categoriesDbTableHelper.DeleteAll()

	// Insert dummy data to DB
	categoriesDbTableHelper.Add(&entity.Category{
		Id:   "CAT-1",
		Name: "Tools",
	})
	// ------------------------

	testRequest := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), nil)

	testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

	recorder := httptest.NewRecorder()

	middlewareTesting := setupMiddleware(vp)

	// Action
	middlewareTesting.ServeHTTP(recorder, testRequest)

	// Assert
	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	internal_helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[struct{}])

	err = json.Unmarshal(responseBodyBytes, webResponse)
	internal_helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, struct{}{}, webResponse.Data)

	assert.Equal(t, 0, len(categoriesDbTableHelper.FindAll()))
}

func TestFindByIdFailed(t *testing.T) {
	t.Run("404 - Category is Not Found", func(t *testing.T) {
		// Arrange
		defer categoriesDbTableHelper.DeleteAll()

		testRequest := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), nil)

		testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

		recorder := httptest.NewRecorder()

		middlewareTesting := setupMiddleware(vp)

		// Action
		middlewareTesting.ServeHTTP(recorder, testRequest)

		// Assert
		recorderResponse := recorder.Result()

		assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
		assert.Equal(t, http.StatusNotFound, recorderResponse.StatusCode)

		responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
		internal_helper.LogStdPanicIfError(err)

		webResponse := new(model.WebResponse[string])

		err = json.Unmarshal(responseBodyBytes, webResponse)
		internal_helper.LogStdPanicIfError(err)

		assert.Equal(t, http.StatusNotFound, webResponse.Code)
		assert.Equal(t, "NOT FOUND", webResponse.Status)
		assert.Equal(t, "category is not found", webResponse.Data)
	})
}

func TestFindByIdSuccess(t *testing.T) {
	// Arrange
	defer categoriesDbTableHelper.DeleteAll()

	// Insert dummy data to DB
	categoriesDbTableHelper.Add(&entity.Category{
		Id:   "CAT-1",
		Name: "Tools",
	})
	// ------------------------

	testRequest := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/categories/CAT-1", baseUrl), nil)

	testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

	recorder := httptest.NewRecorder()

	middlewareTesting := setupMiddleware(vp)

	// Action
	middlewareTesting.ServeHTTP(recorder, testRequest)

	// Assert
	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	internal_helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[*model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, webResponse)
	internal_helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, &model.CategoryResponse{
		Id:   "CAT-1",
		Name: "Tools",
	}, webResponse.Data)
}

func TestFindAllSuccess(t *testing.T) {
	// Arrange
	defer categoriesDbTableHelper.DeleteAll()

	// Insert dummy data to DB
	categoriesDbTableHelper.AddMany([]entity.Category{
		{Id: "CAT-1", Name: "Tools"},
		{Id: "CAT-2", Name: "Foods"},
		{Id: "CAT-3", Name: "Drinks"},
	})
	// ------------------------

	testRequest := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/categories", baseUrl), nil)

	testRequest.Header.Set("X-API-Key", vp.GetString("server.api_key"))

	recorder := httptest.NewRecorder()

	middlewareTesting := setupMiddleware(vp)

	// Action
	middlewareTesting.ServeHTTP(recorder, testRequest)

	// Assert
	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))
	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	internal_helper.LogStdPanicIfError(err)

	webResponse := new(model.WebResponse[[]model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, webResponse)
	internal_helper.LogStdPanicIfError(err)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, []model.CategoryResponse{
		{Id: "CAT-1", Name: "Tools"},
		{Id: "CAT-2", Name: "Foods"},
		{Id: "CAT-3", Name: "Drinks"},
	}, webResponse.Data)
}

package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	internal_controller_http "github.com/syahdaromansyah/pzn-golang-restful-api/internal/controller/http"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
	internal_usecase_mock "github.com/syahdaromansyah/pzn-golang-restful-api/internal/usecase/mock"
)

func TestCreateFailed(t *testing.T) {
	t.Run("Malformed Request Body", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader(""))

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.Panics(t, func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.NewCategoryControllerImpl(categoryUseCase).Create(recorder, testRequest, nil)
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("UseCase Create Method Panic", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader(`{}`))

		testRequest.Header.Add("content-type", "application/json")

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		categoryUseCase.Mock.On("Create", mock.Anything, mock.Anything).Panic("usecase Create method panic")

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.PanicsWithValue(t, "usecase Create method panic", func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.NewCategoryControllerImpl(categoryUseCase).Create(recorder, testRequest, nil)
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "Create", 1)
	})
}

func TestCreateSuccess(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader(`{}`))

	categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

	categoryUseCase.Mock.On("Create", mock.Anything, mock.Anything).Return(&model.CategoryResponse{
		Id:   "CAT-1",
		Name: "Fashions",
	}).Times(1)

	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		internal_controller_http.NewCategoryControllerImpl(categoryUseCase).Create(recorder, testRequest, nil)
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	assert.Equal(t, "application/json", recorderResponse.Header.Get("content-type"))

	assert.Equal(t, http.StatusCreated, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.PanicIfError(err)

	bodyResponse := new(model.WebResponse[*model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, bodyResponse)
	helper.PanicIfError(err)

	assert.Equal(t, &model.WebResponse[*model.CategoryResponse]{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data: &model.CategoryResponse{
			Id:   "CAT-1",
			Name: "Fashions",
		},
	}, bodyResponse)

	categoryUseCase.Mock.AssertExpectations(t)
	categoryUseCase.Mock.AssertNumberOfCalls(t, "Create", 1)
}

func TestUpdateFailed(t *testing.T) {
	t.Run("Malformed Request Body", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodPut, "http://localhost/", strings.NewReader(""))

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.Panics(t, func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.
				NewCategoryControllerImpl(categoryUseCase).
				Update(
					recorder,
					testRequest,
					httprouter.Params{{Key: "categoryId", Value: "CAT-1"}},
				)
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "Update", 0)
	})

	t.Run("UseCase Update Method Panic", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodPut, "http://localhost/", strings.NewReader(`{}`))

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		categoryUseCase.Mock.
			On("Update", mock.Anything, "CAT-1", mock.Anything).
			Panic("usecase Update method panic")

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.PanicsWithValue(t, "usecase Update method panic", func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.
				NewCategoryControllerImpl(categoryUseCase).
				Update(
					recorder,
					testRequest,
					httprouter.Params{{Key: "categoryId", Value: "CAT-1"}},
				)
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "Update", 1)
	})
}

func TestUpdateSuccess(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest(http.MethodPut, "http://localhost/", strings.NewReader(`{}`))

	testRequest.Header.Add("content-type", "application/json")

	categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

	categoryUseCase.Mock.
		On("Update", mock.Anything, "CAT-1", mock.Anything).
		Return(&model.CategoryResponse{
			Id:   "CAT-1",
			Name: "Foods",
		}).Times(1)

	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		internal_controller_http.
			NewCategoryControllerImpl(categoryUseCase).
			Update(
				recorder,
				testRequest,
				httprouter.Params{{Key: "categoryId", Value: "CAT-1"}},
			)
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	assert.Equal(
		t,
		"application/json",
		recorderResponse.Header.Get("content-type"),
	)

	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.PanicIfError(err)

	bodyResponse := new(model.WebResponse[*model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, bodyResponse)
	helper.PanicIfError(err)

	assert.Equal(t, &model.WebResponse[*model.CategoryResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &model.CategoryResponse{
			Id:   "CAT-1",
			Name: "Foods",
		},
	}, bodyResponse)

	categoryUseCase.Mock.AssertExpectations(t)
	categoryUseCase.Mock.AssertNumberOfCalls(t, "Update", 1)
}

func TestDeleteFailed(t *testing.T) {
	t.Run("UseCase Delete Method Panic", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodDelete, "http://localhost/", nil)

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		categoryUseCase.Mock.On("Delete", mock.Anything, "CAT-5").Panic("usecase Delete method panic")

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.PanicsWithValue(t, "usecase Delete method panic", func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.NewCategoryControllerImpl(categoryUseCase).Delete(recorder, testRequest, httprouter.Params{{Key: "categoryId", Value: "CAT-5"}})
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "Delete", 1)
	})
}

func TestDeleteSuccess(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest(http.MethodDelete, "http://localhost/", nil)

	categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

	categoryUseCase.Mock.On("Delete", mock.Anything, "CAT-5").Times(1)

	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		internal_controller_http.NewCategoryControllerImpl(categoryUseCase).Delete(recorder, testRequest, httprouter.Params{{Key: "categoryId", Value: "CAT-5"}})
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.PanicIfError(err)

	bodyResponse := new(model.WebResponse[struct{}])

	err = json.Unmarshal(responseBodyBytes, bodyResponse)
	helper.PanicIfError(err)

	assert.Equal(t, &model.WebResponse[struct{}]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   struct{}{},
	}, bodyResponse)

	categoryUseCase.Mock.AssertExpectations(t)
	categoryUseCase.Mock.AssertNumberOfCalls(t, "Delete", 1)
}

func TestFindByIdFailed(t *testing.T) {
	t.Run("UseCase FindById Method Panic", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		categoryUseCase.Mock.On("FindById", mock.Anything, "CAT-5").Panic("usecase FindById method panic")

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.PanicsWithValue(t, "usecase FindById method panic", func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.NewCategoryControllerImpl(categoryUseCase).FindById(recorder, testRequest, httprouter.Params{{Key: "categoryId", Value: "CAT-5"}})
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "FindById", 1)
	})
}

func TestFindByIdSuccess(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)

	categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

	categoryUseCase.Mock.On("FindById", mock.Anything, "CAT-5").Return(&model.CategoryResponse{
		Id:   "CAT-5",
		Name: "Drinks",
	}).Times(1)

	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		internal_controller_http.NewCategoryControllerImpl(categoryUseCase).FindById(recorder, testRequest, httprouter.Params{{Key: "categoryId", Value: "CAT-5"}})
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	assert.Equal(
		t,
		"application/json",
		recorderResponse.Header.Get("content-type"),
	)

	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.PanicIfError(err)

	bodyResponse := new(model.WebResponse[*model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, bodyResponse)
	helper.PanicIfError(err)

	assert.Equal(t, &model.WebResponse[*model.CategoryResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &model.CategoryResponse{
			Id:   "CAT-5",
			Name: "Drinks",
		},
	}, bodyResponse)

	categoryUseCase.Mock.AssertExpectations(t)
	categoryUseCase.Mock.AssertNumberOfCalls(t, "FindById", 1)
}

func TestFindAllFailed(t *testing.T) {
	t.Run("UseCase FindAll Method Panic", func(t *testing.T) {
		// Arrange
		testRequest := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)

		categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

		categoryUseCase.Mock.On("FindAll", mock.Anything).Panic("usecase FindAll method panic")

		recorder := httptest.NewRecorder()

		// Action & Assert
		assert.PanicsWithValue(t, "usecase FindAll method panic", func() {
			// ---SUT (Subject Under Test)
			internal_controller_http.NewCategoryControllerImpl(categoryUseCase).FindAll(recorder, testRequest, nil)
			// ---------------------------
		})

		categoryUseCase.Mock.AssertExpectations(t)
		categoryUseCase.Mock.AssertNumberOfCalls(t, "FindAll", 1)
	})
}

func TestFindAllSuccess(t *testing.T) {
	// Arrange
	testRequest := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)

	categoryUseCase := internal_usecase_mock.NewCategoryUseCaseMock()

	categoryUseCase.Mock.On("FindAll", mock.Anything).Return([]model.CategoryResponse{
		{Id: "CAT-5", Name: "Drinks"},
		{Id: "CAT-6", Name: "Foods"},
		{Id: "CAT-7", Name: "Vegetables"},
		{Id: "CAT-8", Name: "Meats"},
	}).Times(1)

	recorder := httptest.NewRecorder()

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		internal_controller_http.NewCategoryControllerImpl(categoryUseCase).FindAll(recorder, testRequest, nil)
		// ---------------------------
	})

	recorderResponse := recorder.Result()

	assert.Equal(
		t,
		"application/json",
		recorderResponse.Header.Get("content-type"),
	)

	assert.Equal(t, http.StatusOK, recorderResponse.StatusCode)

	responseBodyBytes, err := io.ReadAll(recorderResponse.Body)
	helper.PanicIfError(err)

	bodyResponse := new(model.WebResponse[[]model.CategoryResponse])

	err = json.Unmarshal(responseBodyBytes, bodyResponse)
	helper.PanicIfError(err)

	assert.Equal(t, &model.WebResponse[[]model.CategoryResponse]{
		Code:   http.StatusOK,
		Status: "OK",
		Data: []model.CategoryResponse{
			{Id: "CAT-5", Name: "Drinks"},
			{Id: "CAT-6", Name: "Foods"},
			{Id: "CAT-7", Name: "Vegetables"},
			{Id: "CAT-8", Name: "Meats"},
		},
	}, bodyResponse)

	categoryUseCase.Mock.AssertExpectations(t)
	categoryUseCase.Mock.AssertNumberOfCalls(t, "FindAll", 1)
}

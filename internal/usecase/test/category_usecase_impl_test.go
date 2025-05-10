package usecase

import (
	"testing"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
	internal_repository_mock "github.com/syahdaromansyah/pzn-golang-restful-api/internal/repository/mock"
	internal_security_mock "github.com/syahdaromansyah/pzn-golang-restful-api/internal/security/mock"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateFailed(t *testing.T) {
	t.Run("Repository Save Method Panic", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		validate := internal_security_mock.NewValidationMock()

		validate.Mock.On("Struct", mock.Anything).Return(validator.New().Struct(&model.CreateCategoryRequest{
			Name: "Fashions",
		}))

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("Save", mock.Anything, mock.Anything, mock.Anything).Panic("repository save method panic")

		// Action & Assert
		assert.PanicsWithValue(t, "repository save method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, validate, categoryRepository).Create(t.Context(), &model.CreateCategoryRequest{
				Name: "Fashions",
			})
			// ---------------------------
		})

		validate.Mock.AssertExpectations(t)
		validate.Mock.AssertNumberOfCalls(t, "Struct", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "Save", 1)
	})
}

func TestCreateSuccess(t *testing.T) {
	// Arrange
	pool, err := pgxmock.NewPool()
	helper.PanicIfError(err)

	defer pool.Close()

	pool.ExpectBegin()
	pool.ExpectCommit()

	validate := internal_security_mock.NewValidationMock()

	validate.Mock.On("Struct", mock.Anything).Return(validator.New().Struct(&model.CreateCategoryRequest{
		Name: "Fashions",
	}))

	categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

	categoryRepository.Mock.On("Save", mock.Anything, mock.Anything, mock.Anything).Return(&entity.Category{
		Id:   "CAT-1",
		Name: "Fashions",
	}).Times(1)

	var result *model.CategoryResponse

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = usecase.NewCategoryUseCaseImpl(pool, validate, categoryRepository).Create(t.Context(), &model.CreateCategoryRequest{
			Name: "Fashions",
		})
		// ---------------------------
	})

	assert.Equal(t, &model.CategoryResponse{
		Id:   "CAT-1",
		Name: "Fashions",
	}, result)

	validate.Mock.AssertExpectations(t)
	validate.Mock.AssertNumberOfCalls(t, "Struct", 1)

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "Save", 1)
}

func TestUpdateFailed(t *testing.T) {
	t.Run("Repository FindById Method Panic", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		validate := internal_security_mock.NewValidationMock()

		validate.Mock.On("Struct", mock.Anything).Return(validator.New().Struct(&model.UpdateCategoryRequest{
			Name: "Electronics",
		}))

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Panic("repository FindById method panic")

		// Action & Assert
		assert.PanicsWithValue(t, "repository FindById method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, validate, categoryRepository).Update(t.Context(), "CAT-1", &model.UpdateCategoryRequest{
				Name: "Electronics",
			})
			// ---------------------------
		})

		validate.Mock.AssertExpectations(t)
		validate.Mock.AssertNumberOfCalls(t, "Struct", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "Update", 0)
	})

	t.Run("Repository Update Method Panic", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		validate := internal_security_mock.NewValidationMock()

		validate.Mock.On("Struct", mock.Anything).Return(validator.New().Struct(&model.UpdateCategoryRequest{
			Name: "Electronics",
		}))

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Return(new(entity.Category))

		categoryRepository.Mock.On("Update", mock.Anything, mock.Anything, mock.Anything).Panic("repository Update method panic")

		// Action & Assert
		assert.PanicsWithValue(t, "repository Update method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, validate, categoryRepository).Update(t.Context(), "CAT-1", &model.UpdateCategoryRequest{
				Name: "Electronics",
			})
			// ---------------------------
		})

		validate.Mock.AssertExpectations(t)
		validate.Mock.AssertNumberOfCalls(t, "Struct", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "Update", 1)
	})
}

func TestUpdateSuccess(t *testing.T) {
	// Arrange
	pool, err := pgxmock.NewPool()
	helper.PanicIfError(err)

	defer pool.Close()

	pool.ExpectBegin()
	pool.ExpectCommit()

	validate := internal_security_mock.NewValidationMock()

	requestBody := &model.UpdateCategoryRequest{
		Name: "Electronics",
	}

	validate.Mock.On("Struct", requestBody).Return(validator.New().Struct(requestBody))

	categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

	categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Return(&entity.Category{
		Id:   "CAT-1",
		Name: "Electronics",
	}).Times(1)

	categoryRepository.Mock.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&entity.Category{
		Id:   "CAT-1",
		Name: "Electronics",
	}).Times(1)

	var result *model.CategoryResponse

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = usecase.NewCategoryUseCaseImpl(pool, validate, categoryRepository).Update(t.Context(), "CAT-1", requestBody)
		// ---------------------------
	})

	assert.Equal(t, &model.CategoryResponse{
		Id:   "CAT-1",
		Name: "Electronics",
	}, result)

	validate.Mock.AssertExpectations(t)
	validate.Mock.AssertNumberOfCalls(t, "Struct", 1)

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "Update", 1)
}

func TestDeleteFailed(t *testing.T) {
	t.Run("Repository FindById Method Panic", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Panic("repository Delete method panic").Times(1)

		// Action & Assert
		assert.PanicsWithValue(t, "repository Delete method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).Delete(t.Context(), "CAT-1")
			// ---------------------------
		})

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "Delete", 0)
	})

	t.Run("Repository Delete Method Panic", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Return(new(entity.Category)).Times(1)

		categoryRepository.Mock.On("Delete", mock.Anything, mock.Anything, "CAT-1").Panic("repository Delete method panic").Times(1)

		// Action & Assert
		assert.PanicsWithValue(t, "repository Delete method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).Delete(t.Context(), "CAT-1")
			// ---------------------------
		})

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "Delete", 1)
	})
}

func TestDeleteSuccess(t *testing.T) {
	// Arrange
	pool, err := pgxmock.NewPool()
	helper.PanicIfError(err)

	defer pool.Close()

	pool.ExpectBegin()
	pool.ExpectCommit()

	categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

	categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Return(new(entity.Category)).Times(1)
	categoryRepository.Mock.On("Delete", mock.Anything, mock.Anything, "CAT-1").Times(1)

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).Delete(t.Context(), "CAT-1")
		// ---------------------------
	})

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "Delete", 1)
}

func TestFindByIdFailed(t *testing.T) {
	t.Run("Repository FindById Method Failed", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Panic("repository FindById method panic")

		// Action & Assert
		assert.PanicsWithValue(t, "repository FindById method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).FindById(t.Context(), "CAT-1")
			// ---------------------------
		})

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)
	})
}

func TestFindByIdSuccess(t *testing.T) {
	// Arrange
	pool, err := pgxmock.NewPool()
	helper.PanicIfError(err)

	defer pool.Close()

	pool.ExpectBegin()
	pool.ExpectCommit()

	categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

	categoryRepository.Mock.On("FindById", mock.Anything, mock.Anything, "CAT-1").Return(&entity.Category{
		Id:   "CAT-1",
		Name: "Drinks",
	}).Times(1)

	var result *model.CategoryResponse

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).FindById(t.Context(), "CAT-1")
		// ---------------------------
	})

	assert.Equal(t, &model.CategoryResponse{
		Id:   "CAT-1",
		Name: "Drinks",
	}, result)

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "FindById", 1)
}

func TestFindAllFailed(t *testing.T) {
	t.Run("Repository FindAll Method Panic", func(t *testing.T) {
		// Arrange
		pool, err := pgxmock.NewPool()
		helper.PanicIfError(err)

		defer pool.Close()

		pool.ExpectBegin()
		pool.ExpectRollback()

		categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

		categoryRepository.Mock.On("FindAll", mock.Anything, mock.Anything).Panic("repository FindAll method panic")

		// Action & Assert
		assert.PanicsWithValue(t, "repository FindAll method panic", func() {
			// ---SUT (Subject Under Test)
			usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).FindAll(t.Context())
			// ---------------------------
		})

		categoryRepository.Mock.AssertExpectations(t)
		categoryRepository.Mock.AssertNumberOfCalls(t, "FindAll", 1)
	})
}

func TestFindAllSuccess(t *testing.T) {
	// Arrange
	pool, err := pgxmock.NewPool()
	helper.PanicIfError(err)

	defer pool.Close()

	pool.ExpectBegin()
	pool.ExpectCommit()

	categoryRepository := internal_repository_mock.NewCategoryRepositoryMock()

	categoryRepository.Mock.On("FindAll", mock.Anything, mock.Anything).Return([]entity.Category{
		{Id: "CAT-1", Name: "Drinks"},
		{Id: "CAT-2", Name: "Foods"},
		{Id: "CAT-3", Name: "Furniture"},
	}).Times(1)

	var result []model.CategoryResponse

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = usecase.NewCategoryUseCaseImpl(pool, nil, categoryRepository).FindAll(t.Context())
		// ---------------------------
	})

	assert.Equal(t, []model.CategoryResponse{
		{Id: "CAT-1", Name: "Drinks"},
		{Id: "CAT-2", Name: "Foods"},
		{Id: "CAT-3", Name: "Furniture"},
	}, result)

	categoryRepository.Mock.AssertExpectations(t)
	categoryRepository.Mock.AssertNumberOfCalls(t, "FindAll", 1)
}

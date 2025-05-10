package repository

import (
	"context"
	"testing"
	"time"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/repository"
	internal_security_mock "github.com/syahdaromansyah/pzn-golang-restful-api/internal/security/mock"
	test_helper "github.com/syahdaromansyah/pzn-golang-restful-api/test/helper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var vp = config.NewViper([]string{"./../../.."})

func TestSaveFailed(t *testing.T) {
	t.Run("Field 'name' | Constraint Check | Chars Min", func(t *testing.T) {
		// Arrange
		pool := config.NewPgxPool(vp)
		defer pool.Close()

		ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
		defer cancel()

		tx, err := pool.Begin(ctx)
		helper.PanicIfError(err)

		defer helper.TxRollbackIfPanic(ctx, tx)

		idGen := internal_security_mock.NewIdGenMock()

		idGen.Mock.On("Generate", 36).Return("CAT-1", nil).Times(1)

		// Action & Assert
		if assert.Panics(t, func() {
			// ---SUT (Subject Under Test)
			repository.NewCategoryRepositoryImpl(idGen).
				Save(ctx, tx, &entity.Category{
					Name: "A",
				})
			// ---------------------------
		}) {
			err := tx.Rollback(ctx)
			helper.PanicIfError(err)
		}

		idGen.Mock.AssertExpectations(t)
		idGen.Mock.AssertNumberOfCalls(t, "Generate", 1)
	})
}

func TestSaveSuccess(t *testing.T) {
	t.Run("Create Uncreated Category", func(t *testing.T) {
		// Arrange
		pool := config.NewPgxPool(vp)
		defer pool.Close()

		ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
		defer cancel()

		tx, err := pool.Begin(ctx)
		helper.PanicIfError(err)

		defer helper.TxRollbackIfPanic(ctx, tx)

		idGen := internal_security_mock.NewIdGenMock()

		idGen.Mock.On("Generate", 36).Return("CAT-1", nil).Times(1)

		var result *entity.Category

		// Action & Assert
		assert.NotPanics(t, func() {
			// ---SUT (Subject Under Test)
			result = repository.NewCategoryRepositoryImpl(idGen).Save(ctx, tx, &entity.Category{
				Name: "Beverages",
			})
			// ---------------------------
		})

		helper.TxCommit(ctx, tx)

		assert.Equal(t, &entity.Category{
			Id:   "CAT-1",
			Name: "Beverages",
		}, result)

		idGen.Mock.AssertExpectations(t)
		idGen.Mock.AssertNumberOfCalls(t, "Generate", 1)

		dbHelper := test_helper.NewCategoriesDbTable(vp)
		defer dbHelper.DeleteAll()

		assert.Equal(t, 1, len(dbHelper.FindAll()))

		assert.Equal(t, &entity.Category{
			Id:   "CAT-1",
			Name: "Beverages",
		}, dbHelper.FindById("CAT-1"))
	})

	t.Run("Create Category with Another Id", func(t *testing.T) {
		// Arrange

		// --- Insert dummy data to DB
		dbHelper := test_helper.NewCategoriesDbTable(vp)
		defer dbHelper.DeleteAll()

		dbHelper.Add(&entity.Category{
			Id:   "CAT-1",
			Name: "Medicines",
		})
		// --- END

		pool := config.NewPgxPool(vp)
		defer pool.Close()

		ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
		defer cancel()

		tx, err := pool.Begin(ctx)
		helper.PanicIfError(err)

		defer helper.TxRollbackIfPanic(ctx, tx)

		idGen := internal_security_mock.NewIdGenMock()

		mock.InOrder(
			idGen.Mock.On("Generate", 36).Return("CAT-1", nil).Times(1),
			idGen.Mock.On("Generate", 36).Return("CAT-5", nil).Times(1),
		)

		var result *entity.Category

		// Action & Assert
		assert.NotPanics(t, func() {
			// ---SUT (Subject Under Test)
			result = repository.NewCategoryRepositoryImpl(idGen).Save(ctx, tx, &entity.Category{
				Name: "Beverages",
			})
			// ---------------------------
		})

		helper.TxCommit(ctx, tx)

		assert.Equal(t, &entity.Category{
			Id:   "CAT-5",
			Name: "Beverages",
		}, result)

		idGen.Mock.AssertExpectations(t)
		idGen.Mock.AssertNumberOfCalls(t, "Generate", 2)

		assert.Equal(t, 2, len(dbHelper.FindAll()))

		assert.Equal(t, &entity.Category{
			Id:   "CAT-5",
			Name: "Beverages",
		}, dbHelper.FindById("CAT-5"))
	})
}

func TestUpdateFailed(t *testing.T) {
	t.Run("Field 'name' | Constraint Check | Chars Min", func(t *testing.T) {
		// Arrange

		// --- Insert dummy data to DB
		dbHelper := test_helper.NewCategoriesDbTable(vp)
		defer dbHelper.DeleteAll()

		dbHelper.Add(&entity.Category{
			Id:   "CAT-1",
			Name: "Medicines",
		})
		// --- END

		pool := config.NewPgxPool(vp)
		defer pool.Close()

		ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
		defer cancel()

		tx, err := pool.Begin(ctx)
		helper.PanicIfError(err)

		defer helper.TxRollbackIfPanic(ctx, tx)

		// Action & Assert
		if assert.Panics(t, func() {
			// ---SUT (Subject Under Test)
			repository.NewCategoryRepositoryImpl(nil).Update(ctx, tx, &entity.Category{
				Id:   "CAT-1",
				Name: "A",
			})
			// ---------------------------
		}) {
			err := tx.Rollback(ctx)
			helper.PanicIfError(err)
		}
	})
}

func TestUpdateSuccess(t *testing.T) {
	// Arrange

	// --- Insert dummy data to DB
	dbHelper := test_helper.NewCategoriesDbTable(vp)
	defer dbHelper.DeleteAll()

	dbHelper.Add(&entity.Category{
		Id:   "CAT-1",
		Name: "Medicines",
	})
	// --- END

	pool := config.NewPgxPool(vp)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.TxRollbackIfPanic(ctx, tx)

	var result *entity.Category

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = repository.NewCategoryRepositoryImpl(nil).Update(ctx, tx, &entity.Category{
			Id:   "CAT-1",
			Name: "Fashions",
		})
		// ---------------------------
	})

	helper.TxCommit(ctx, tx)

	assert.Equal(t, &entity.Category{
		Id:   "CAT-1",
		Name: "Fashions",
	}, result)

	assert.Equal(t, &entity.Category{
		Id:   "CAT-1",
		Name: "Fashions",
	}, dbHelper.FindById("CAT-1"))
}

func TestDeleteSuccess(t *testing.T) {
	// Arrange

	// --- Insert dummy data to DB
	dbHelper := test_helper.NewCategoriesDbTable(vp)
	defer dbHelper.DeleteAll()

	dbHelper.Add(&entity.Category{
		Id:   "CAT-1",
		Name: "Medicines",
	})
	// --- END

	pool := config.NewPgxPool(vp)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.TxRollbackIfPanic(ctx, tx)

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		repository.NewCategoryRepositoryImpl(nil).Delete(ctx, tx, "CAT-1")
		// ---------------------------
	})

	helper.TxCommit(ctx, tx)

	assert.Equal(t, 0, len(dbHelper.FindAll()))
}

func TestFindByIdFailed(t *testing.T) {
	t.Run("Category is Not Found", func(t *testing.T) {
		// Arrange
		pool := config.NewPgxPool(vp)
		defer pool.Close()

		ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
		defer cancel()

		tx, err := pool.Begin(ctx)
		helper.PanicIfError(err)

		defer helper.TxRollbackIfPanic(ctx, tx)

		// Action & Assert
		if assert.Panics(t, func() {
			// ---SUT (Subject Under Test)
			repository.NewCategoryRepositoryImpl(nil).FindById(context.Background(), tx, "CAT-1")
			// ---------------------------
		}) {
			err := tx.Rollback(ctx)
			helper.PanicIfError(err)
		}
	})
}

func TestFindByIdSuccess(t *testing.T) {
	// Arrange

	// --- Insert dummy data to DB
	dbHelper := test_helper.NewCategoriesDbTable(vp)
	defer dbHelper.DeleteAll()

	dbHelper.Add(&entity.Category{
		Id:   "CAT-1",
		Name: "Medicines",
	})
	// --- END

	pool := config.NewPgxPool(vp)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.TxRollbackIfPanic(ctx, tx)

	var result *entity.Category

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = repository.NewCategoryRepositoryImpl(nil).FindById(context.Background(), tx, "CAT-1")
		// ---------------------------
	})

	helper.TxCommit(ctx, tx)

	assert.Equal(t, &entity.Category{
		Id:   "CAT-1",
		Name: "Medicines",
	}, result)
}

func TestFindAllSuccess(t *testing.T) {
	// Arrange

	// --- Insert dummy data to DB
	dbHelper := test_helper.NewCategoriesDbTable(vp)
	defer dbHelper.DeleteAll()

	dbHelper.AddMany([]entity.Category{
		{Id: "CAT-1", Name: "Medicines"},
		{Id: "C-2", Name: "Fashions"},
		{Id: "C-3", Name: "Toys"},
	})
	// --- END

	pool := config.NewPgxPool(vp)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), vp.GetDuration("test.timeout")*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	defer helper.TxRollbackIfPanic(ctx, tx)

	var result []entity.Category

	// Action & Assert
	assert.NotPanics(t, func() {
		// ---SUT (Subject Under Test)
		result = repository.NewCategoryRepositoryImpl(nil).FindAll(context.Background(), tx)
		// ---------------------------
	})

	helper.TxCommit(ctx, tx)

	assert.Equal(t, []entity.Category{
		{Id: "CAT-1", Name: "Medicines"},
		{Id: "C-2", Name: "Fashions"},
		{Id: "C-3", Name: "Toys"},
	}, result)
}

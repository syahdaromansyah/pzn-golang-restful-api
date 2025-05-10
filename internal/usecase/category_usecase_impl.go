package usecase

import (
	"context"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/db"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model/converter"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/repository"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/security"
)

type categoryUseCaseImpl struct {
	DB                 db.PgxPool
	Validator          security.Validation
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUseCaseImpl(db db.PgxPool, validate security.Validation, categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCaseImpl{
		DB:                 db,
		Validator:          validate,
		CategoryRepository: categoryRepository,
	}
}

func (u *categoryUseCaseImpl) Create(ctx context.Context, requestBody *model.CreateCategoryRequest) *model.CategoryResponse {
	err := u.Validator.Struct(requestBody)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin(ctx)
	helper.InternalServerPanicIfError(err, "category > usecase > Create")

	defer helper.TxCommitRollback(ctx, tx)

	category := &entity.Category{
		Name: requestBody.Name,
	}

	category = u.CategoryRepository.Save(ctx, tx, category)

	return converter.CategoryToResponse(category)
}

func (u *categoryUseCaseImpl) Update(ctx context.Context, categoryId string, requestBody *model.UpdateCategoryRequest) *model.CategoryResponse {
	err := u.Validator.Struct(requestBody)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin(ctx)
	helper.InternalServerPanicIfError(err, "category > usecase > Update")

	defer helper.TxCommitRollback(ctx, tx)

	u.CategoryRepository.FindById(ctx, tx, categoryId)

	category := &entity.Category{
		Id:   categoryId,
		Name: requestBody.Name,
	}

	category = u.CategoryRepository.Update(ctx, tx, category)

	return converter.CategoryToResponse(category)
}

func (u *categoryUseCaseImpl) Delete(ctx context.Context, categoryId string) {
	tx, err := u.DB.Begin(ctx)
	helper.InternalServerPanicIfError(err, "category > usecase > Delete")

	defer helper.TxCommitRollback(ctx, tx)

	u.CategoryRepository.FindById(ctx, tx, categoryId)
	u.CategoryRepository.Delete(ctx, tx, categoryId)
}

func (u *categoryUseCaseImpl) FindById(ctx context.Context, categoryId string) *model.CategoryResponse {
	tx, err := u.DB.Begin(ctx)
	helper.InternalServerPanicIfError(err, "category > usecase > FindById")

	defer helper.TxCommitRollback(ctx, tx)

	result := u.CategoryRepository.FindById(ctx, tx, categoryId)

	return converter.CategoryToResponse(result)
}

func (u *categoryUseCaseImpl) FindAll(ctx context.Context) []model.CategoryResponse {
	tx, err := u.DB.Begin(ctx)
	helper.InternalServerPanicIfError(err, "category > usecase > FindAll")

	defer helper.TxCommitRollback(ctx, tx)

	result := u.CategoryRepository.FindAll(ctx, tx)

	return converter.CategoriesToResponse(result)
}

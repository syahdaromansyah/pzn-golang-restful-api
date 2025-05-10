package usecase

import (
	"context"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
)

type CategoryUseCase interface {
	Create(ctx context.Context, requestBody *model.CreateCategoryRequest) *model.CategoryResponse
	Update(ctx context.Context, categoryId string, requestBody *model.UpdateCategoryRequest) *model.CategoryResponse
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) *model.CategoryResponse
	FindAll(ctx context.Context) []model.CategoryResponse
}

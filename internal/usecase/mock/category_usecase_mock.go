package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"
)

type categoryUseCaseMock struct {
	Mock *mock.Mock
}

func NewCategoryUseCaseMock() *categoryUseCaseMock {
	return &categoryUseCaseMock{
		Mock: new(mock.Mock),
	}
}

func (u *categoryUseCaseMock) Create(ctx context.Context, requestBody *model.CreateCategoryRequest) *model.CategoryResponse {
	args := u.Mock.Called(ctx, requestBody)
	return args.Get(0).(*model.CategoryResponse)
}

func (u *categoryUseCaseMock) Update(ctx context.Context, categoryId string, requestBody *model.UpdateCategoryRequest) *model.CategoryResponse {
	args := u.Mock.Called(ctx, categoryId, requestBody)
	return args.Get(0).(*model.CategoryResponse)
}

func (u *categoryUseCaseMock) Delete(ctx context.Context, categoryId string) {
	u.Mock.Called(ctx, categoryId)
}

func (u *categoryUseCaseMock) FindById(ctx context.Context, categoryId string) *model.CategoryResponse {
	args := u.Mock.Called(ctx, categoryId)
	return args.Get(0).(*model.CategoryResponse)
}

func (u *categoryUseCaseMock) FindAll(ctx context.Context) []model.CategoryResponse {
	args := u.Mock.Called(ctx)
	return args.Get(0).([]model.CategoryResponse)
}

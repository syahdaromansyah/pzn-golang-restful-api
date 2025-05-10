package repository

import (
	"context"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type categoryRepositoryMock struct {
	Mock *mock.Mock
}

func NewCategoryRepositoryMock() *categoryRepositoryMock {
	return &categoryRepositoryMock{
		Mock: new(mock.Mock),
	}
}

func (r *categoryRepositoryMock) Save(ctx context.Context, tx pgx.Tx, category *entity.Category) *entity.Category {
	args := r.Mock.Called(ctx, tx, category)
	return args.Get(0).(*entity.Category)
}

func (r *categoryRepositoryMock) Update(ctx context.Context, tx pgx.Tx, category *entity.Category) *entity.Category {
	args := r.Mock.Called(ctx, tx, category)
	return args.Get(0).(*entity.Category)
}

func (r *categoryRepositoryMock) Delete(ctx context.Context, tx pgx.Tx, categoryId string) {
	r.Mock.Called(ctx, tx, categoryId)
}

func (r *categoryRepositoryMock) FindById(ctx context.Context, tx pgx.Tx, categoryId string) *entity.Category {
	args := r.Mock.Called(ctx, tx, categoryId)
	return args.Get(0).(*entity.Category)
}

func (r *categoryRepositoryMock) FindAll(ctx context.Context, tx pgx.Tx) []entity.Category {
	args := r.Mock.Called(ctx, tx)
	return args.Get(0).([]entity.Category)
}

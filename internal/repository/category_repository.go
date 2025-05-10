package repository

import (
	"context"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"

	"github.com/jackc/pgx/v5"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx pgx.Tx, category *entity.Category) *entity.Category
	Update(ctx context.Context, tx pgx.Tx, category *entity.Category) *entity.Category
	Delete(ctx context.Context, tx pgx.Tx, categoryId string)
	FindById(ctx context.Context, tx pgx.Tx, categoryId string) *entity.Category
	FindAll(ctx context.Context, tx pgx.Tx) []entity.Category
}

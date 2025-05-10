package repository

import (
	"context"
	"net/http"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/exception"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/security"

	"github.com/jackc/pgx/v5"
)

type categoryRepositoryImpl struct {
	IdGenerator security.IdGenerator
}

func NewCategoryRepositoryImpl(idGenerator security.IdGenerator) CategoryRepository {
	return &categoryRepositoryImpl{
		IdGenerator: idGenerator,
	}
}

func (r *categoryRepositoryImpl) Save(ctx context.Context, tx pgx.Tx, category *entity.Category) *entity.Category {
	for {
		generatedId, err := r.IdGenerator.Generate(36)
		helper.InternalServerPanicIfError(err, "category > repository > Save")

		rows, err := tx.Query(ctx, "SELECT id FROM categories WHERE id = $1 LIMIT 1", generatedId)
		helper.InternalServerPanicIfError(err, "category > repository > Save")

		categoryIds, err := pgx.CollectRows(rows, pgx.RowTo[string])
		helper.InternalServerPanicIfError(err, "category > repository > Save")

		if len(categoryIds) == 0 {
			_, err := tx.Exec(ctx, "INSERT INTO categories (id, name) VALUES ($1, $2)", generatedId, category.Name)
			helper.InternalServerPanicIfError(err, "category > repository > Save")

			category.Id = generatedId
			break
		}
	}

	return category
}

func (r *categoryRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, category *entity.Category) *entity.Category {
	_, err := tx.Exec(ctx, "UPDATE categories SET name = $1 WHERE id = $2", category.Name, category.Id)
	helper.InternalServerPanicIfError(err, "category > repository > Update")

	return category
}

func (r *categoryRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, categoryId string) {
	_, err := tx.Exec(ctx, "DELETE FROM categories WHERE id = $1", categoryId)
	helper.InternalServerPanicIfError(err, "category > repository > Delete")
}

func (r *categoryRepositoryImpl) FindById(ctx context.Context, tx pgx.Tx, categoryId string) *entity.Category {
	result := new(entity.Category)

	err := tx.QueryRow(ctx, "SELECT id, name FROM categories WHERE id = $1", categoryId).Scan(&result.Id, &result.Name)
	helper.ClientPanicIfError(err, exception.NewErrorClientRequest(err, http.StatusNotFound, "category is not found"))

	return result
}

func (r *categoryRepositoryImpl) FindAll(ctx context.Context, tx pgx.Tx) []entity.Category {
	rows, err := tx.Query(ctx, "SELECT id, name FROM categories")
	helper.InternalServerPanicIfError(err, "category > repository > FindAll")

	defer rows.Close()

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Category])
	helper.InternalServerPanicIfError(err, "category > repository > FindAll")

	return result
}

package helper

import (
	"context"
	"time"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/config"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/entity"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"

	"github.com/jackc/pgx/v5"
)

type categoriesDbTable struct {
	AppConfig *config.AppConfig
}

func NewCategoriesDbTable(appConfig *config.AppConfig) *categoriesDbTable {
	return &categoriesDbTable{
		AppConfig: appConfig,
	}
}

func (d *categoriesDbTable) Add(data *entity.Category) {
	pool := config.NewPgxPool(d.AppConfig)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.AppConfig.Test.Timeout*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.LogStdPanicIfError(err)

	_, err = tx.Exec(ctx, "INSERT INTO categories (id, name) VALUES ($1, $2)", data.Id, data.Name)
	helper.TxRollbackIfError(ctx, tx, err)

	helper.TxCommit(ctx, tx)
}

func (d *categoriesDbTable) AddMany(data []entity.Category) {
	pool := config.NewPgxPool(d.AppConfig)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.AppConfig.Test.Timeout*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.LogStdPanicIfError(err)

	for _, eachData := range data {
		_, err = tx.Exec(ctx, "INSERT INTO categories (id, name) VALUES ($1, $2)", eachData.Id, eachData.Name)
		helper.TxRollbackIfError(ctx, tx, err)
	}

	helper.TxCommit(ctx, tx)
}

func (d *categoriesDbTable) DeleteAll() {
	pool := config.NewPgxPool(d.AppConfig)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.AppConfig.Test.Timeout*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.LogStdPanicIfError(err)

	_, err = tx.Exec(ctx, "DELETE FROM categories")
	helper.TxRollbackIfError(ctx, tx, err)

	helper.TxCommit(ctx, tx)
}

func (d *categoriesDbTable) FindById(categoryId string) *entity.Category {
	pool := config.NewPgxPool(d.AppConfig)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.AppConfig.Test.Timeout*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.LogStdPanicIfError(err)

	result := new(entity.Category)

	err = tx.QueryRow(ctx, "SELECT id, name FROM categories WHERE id = $1 LIMIT 1", categoryId).Scan(&result.Id, &result.Name)
	helper.TxRollbackIfError(ctx, tx, err)

	helper.TxCommit(ctx, tx)

	return result
}

func (d *categoriesDbTable) FindAll() []entity.Category {
	pool := config.NewPgxPool(d.AppConfig)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.AppConfig.Test.Timeout*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	helper.LogStdPanicIfError(err)

	rows, err := tx.Query(ctx, "SELECT id, name FROM categories")
	helper.TxRollbackIfError(ctx, tx, err)

	defer rows.Close()

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Category])
	helper.TxRollbackIfError(ctx, tx, err)

	helper.TxCommit(ctx, tx)

	return result
}

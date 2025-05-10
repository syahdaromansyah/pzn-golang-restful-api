package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgxPool interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Close()
}

package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func TxCommitRollback(ctx context.Context, tx pgx.Tx) {
	errRecover := recover()

	if errRecover != nil {
		errRollback := tx.Rollback(ctx)
		PanicIfError(errRollback)

		panic(errRecover)
	} else {
		errCommit := tx.Commit(ctx)
		PanicIfError(errCommit)
	}
}

func TxRollbackIfPanic(ctx context.Context, tx pgx.Tx) {
	errRecover := recover()

	if errRecover != nil {
		err := tx.Rollback(ctx)
		PanicIfError(err)
		panic(errRecover)
	}
}

func TxRollbackIfError(ctx context.Context, tx pgx.Tx, err error) {
	if err != nil {
		errRollback := tx.Rollback(ctx)
		PanicIfError(errRollback)
		panic(err)
	}
}

func TxCommit(ctx context.Context, tx pgx.Tx) {
	errCommit := tx.Commit(ctx)
	PanicIfError(errCommit)
}

package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/db"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
)

func NewPgxPool(appConfig *AppConfig) db.PgxPool {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.DBName)

	pgxPoolCfg, err := pgxpool.ParseConfig(connString)
	helper.LogStdPanicIfError(err)

	pgxPoolCfg.MinConns = int32(appConfig.Database.MinConns)
	pgxPoolCfg.MaxConns = int32(appConfig.Database.MaxConns)
	pgxPoolCfg.MaxConnLifetime = appConfig.Database.MaxConnLifeTime * time.Minute
	pgxPoolCfg.MaxConnIdleTime = appConfig.Database.MaxConnIdleTime * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pgxPool, err := pgxpool.NewWithConfig(ctx, pgxPoolCfg)
	helper.LogStdPanicIfError(err)

	return pgxPool
}

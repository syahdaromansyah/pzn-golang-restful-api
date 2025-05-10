package config

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/db"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
)

func NewPgxPool(vp *viper.Viper) db.PgxPool {
	pgxPoolCfg, err := pgxpool.ParseConfig(vp.GetString("database.connstring"))
	helper.LogStdPanicIfError(err)

	pgxPoolCfg.MinConns = int32(vp.GetInt("database.dbconfig.minconn"))
	pgxPoolCfg.MaxConns = int32(vp.GetInt("database.dbconfig.maxconn"))
	pgxPoolCfg.MaxConnLifetime = time.Duration(vp.GetInt("database.dbconfig.maxconn_lifetime")) * time.Minute
	pgxPoolCfg.MaxConnIdleTime = time.Duration(vp.GetInt("database.dbconfig.maxconn_idletime")) * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pgxPool, err := pgxpool.NewWithConfig(ctx, pgxPoolCfg)
	helper.LogStdPanicIfError(err)

	return pgxPool
}

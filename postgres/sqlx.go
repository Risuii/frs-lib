package frsPostgres

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSQLX(ctx context.Context, cfg PostgresConfig) (*sqlx.DB, error) {
	xDB, err := sqlx.Connect("postgres", cfg.ConnectionUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = xDB.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	xDB.SetMaxOpenConns(cfg.MaxPoolSize)
	xDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	xDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	xDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	return xDB, nil
}

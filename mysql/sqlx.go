package frsMySQL

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitSQLX(ctx context.Context, cfg MySQLConfig) (*sqlx.DB, error) {
	xDB, err := sqlx.Connect("mysql", cfg.ConnectionUrl)
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

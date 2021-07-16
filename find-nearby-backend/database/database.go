package database

import (
	"find-nearby-backend/config"
	"find-nearby-backend/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // init PG driver as required by sqlx
)

func New(cfg config.Config, logger logger.Logger) (*sqlx.DB, error) {
	db, err := newDBHandle(cfg.DatabaseConnectionURL(), cfg.DatabaseMaxPoolSize(), cfg.DatabaseMaxIdleConn())
	if err != nil {
		return db, err
	}
	logger.Infof("database connection initialized with pool size %d", cfg.DatabaseMaxPoolSize())
	return db, err
}

func newDBHandle(conStr string, maxPoolSize int, maxIdleConns int) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxPoolSize)
	db.SetMaxIdleConns(maxIdleConns)
	return db, nil
}

package server

import (
	"find-nearby-backend/config"
	"find-nearby-backend/database"
	"find-nearby-backend/logger"
)

// Start starts the app
func Start(cfg config.Config) {
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())
	db, err := database.New(cfg, log)
	if err != nil {
		log.Panicf(err.Error())
	}
	srv := NewServer(cfg.Addr(), db, log)
	srv.Start()
}

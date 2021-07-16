package config

import (
	"time"

	"github.com/spf13/viper"
)

type databaseConfig struct {
	host                    string
	port                    int
	username                string
	password                string
	name                    string
	maxPoolSize             int
	maxIdleConn             int
	maxOpenConn             int
	maxConnLifetimeDuration time.Duration
}

func newDatabaseConfig(vp *viper.Viper) *databaseConfig {
	return &databaseConfig{
		host:                    vp.GetString("DB_HOST"),
		port:                    vp.GetInt("DB_PORT"),
		name:                    vp.GetString("DB_NAME"),
		username:                vp.GetString("DB_USER"),
		password:                vp.GetString("DB_PASS"),
		maxPoolSize:             10,
		maxIdleConn:             vp.GetInt("DB_MAX_IDLE_CONN"),
		maxOpenConn:             vp.GetInt("DB_MAX_OPEN_CONN"),
		maxConnLifetimeDuration: vp.GetDuration("DB_CONN_MAX_LIFETIME"),
	}
}

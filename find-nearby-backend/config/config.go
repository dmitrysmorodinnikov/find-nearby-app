package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config interface {
	Addr() string
	DatabaseMaxPoolSize() int
	DatabaseConnectionURL() string
	DatabaseMaxIdleConn() int
	DatabaseMaxOpenConn() int
	DatabaseConnMaxLifetime() time.Duration
	LogLevel() string
	LogFormat() string
}

type config struct {
	appHost   string
	appPort   int
	dbConfig  *databaseConfig
	logLevel  string
	logFormat string
}

func LoadConfig() Config {
	vp := newWithViper()
	return config{
		appHost:   vp.GetString("APP_HOST"),
		appPort:   vp.GetInt("APP_PORT"),
		dbConfig:  newDatabaseConfig(vp),
		logLevel:  vp.GetString("LOG_LEVEL"),
		logFormat: vp.GetString("LOG_FORMAT"),
	}
}

// Addr returns the address of the web service
func (c config) Addr() string {
	return fmt.Sprintf("%s:%d", c.appHost, c.appPort)
}

// DatabaseMaxPoolSize returns max pool size for DB
func (c config) DatabaseMaxPoolSize() int {
	return c.dbConfig.maxPoolSize
}

// DatabaseConnectionURL returns connection url for DB
func (c config) DatabaseConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.dbConfig.username, c.dbConfig.password, c.dbConfig.host, c.dbConfig.port, c.dbConfig.name)
}

// DatabaseMaxIdleConn returns max idle connections for DB
func (c config) DatabaseMaxIdleConn() int {
	return c.dbConfig.maxIdleConn
}

// DatabaseMaxOpenConn return max open conn for DB
func (c config) DatabaseMaxOpenConn() int {
	return c.dbConfig.maxOpenConn
}

// DatabaseConnMaxLifetime returns max connection lifetime duration for DB
func (c config) DatabaseConnMaxLifetime() time.Duration {
	return c.dbConfig.maxConnLifetimeDuration
}

// LogLevel return log level config
func (c config) LogLevel() string {
	return c.logLevel
}

// LogFormat returns log format config
func (c config) LogFormat() string {
	return c.logFormat
}

func newWithViper() *viper.Viper {
	vp := viper.New()
	vp.AutomaticEnv()
	vp.SetConfigName("application")
	vp.AddConfigPath("./")
	vp.AddConfigPath("../")
	vp.AddConfigPath("../../")
	vp.ReadInConfig()
	return vp
}

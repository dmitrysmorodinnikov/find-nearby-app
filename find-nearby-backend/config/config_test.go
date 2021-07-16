package config_test

import (
	"find-nearby-backend/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	c := config.LoadConfig()
	assert.Equal(t, "debug", c.LogLevel())
	assert.Equal(t, "plaintext", c.LogFormat())
	assert.Equal(t, "localhost:3333", c.Addr())
	assert.Equal(t, 10, c.DatabaseMaxIdleConn())
	assert.Equal(t, 200, c.DatabaseMaxOpenConn())
	assert.Equal(t, 10, c.DatabaseMaxPoolSize())
	assert.Equal(t, "postgres://postgres:postgres@localhost:5430/find_nearby_dev?sslmode=disable", c.DatabaseConnectionURL())
}

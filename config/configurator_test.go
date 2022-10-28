package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigurator(t *testing.T) {
	cfg := NewConfigurator("test.yaml", ".")
	assert.NotEmpty(t, cfg)

	assert.Equal(t, cfg.Environment, "test")

	assert.Equal(t, cfg.Server.Host, "127.0.0.1")
	assert.Equal(t, cfg.Server.Port, "8080")

	assert.Equal(t, cfg.Database.Host, "host")
	assert.Equal(t, cfg.Database.User, "user")
	assert.Equal(t, cfg.Database.Password, "password")
	assert.Equal(t, cfg.Database.Name, "database")
	assert.Equal(t, cfg.Database.Port, "5432")
	assert.Equal(t, cfg.Database.Ssl, "disable")
	assert.Equal(t, cfg.Database.Timezone, "Europe/Berlin")
	assert.Equal(t, cfg.Database.Network, "tcp")
	assert.Equal(t, cfg.Database.Charset, "utf8mb4")
	assert.False(t, cfg.Database.Migrate)
}

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDBWithoutConfig(t *testing.T) {
	_, err := InitDB()
	assert.NotNil(t, err)
}

func TestInitDBWithConfig(t *testing.T) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "gang")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_SCHEMA", "simple_rest_la")
	db, err := InitDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestCreateRouter(t *testing.T) {
	router := CreateRouter()
	assert.NotNil(t, router)
}

package config

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

// Config represents the app's base configuration.
type Config struct {
	Port   int
	Env    string
	Cors   []string
	Cookie struct {
		Name     string
		SameSite string
	}
	DB struct {
		DSN string
	}
	JWT []byte
}

// AppConfig represents the shared application configuration.
type AppConfig struct {
	Version         string
	ServerStartTime time.Time
	Config          Config
	Logger          *log.Logger
	Validator       *validator.Validate
}

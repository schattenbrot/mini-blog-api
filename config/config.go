package config

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// Config represents the app's base configuration.
type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
	JWT []byte
}

// AppConfig represents the shared application configuration.
type AppConfig struct {
	Version   string
	Config    Config
	Logger    *log.Logger
	Validator *validator.Validate
}

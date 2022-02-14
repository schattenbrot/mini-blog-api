package config

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
	JWT []byte
}

type AppConfig struct {
	Version   string
	Config    Config
	Logger    *log.Logger
	Validator *validator.Validate
}

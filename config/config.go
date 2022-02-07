package config

import (
	"log"
)

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
}

type AppConfig struct {
	Version string
	Config  Config
	Logger  *log.Logger
}

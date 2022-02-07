package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func LoadConfig(cfg *Config) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	portString, ok := viper.Get("PORT").(string)
	if !ok {
		log.Fatal("could not find port")
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("couldn't convert port to int")
	}
	cfg.Port = port

	env, ok := viper.Get("ENVIRONMENT").(string)
	if !ok {
		log.Fatal("could not find environment")
	}
	cfg.Env = env

	dsn, ok := viper.Get("DSN").(string)
	if !ok {
		log.Fatal("could not find dsn")
	}
	cfg.DB.DSN = dsn
}

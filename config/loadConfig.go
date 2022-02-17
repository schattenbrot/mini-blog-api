package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

// LoadConfig loads the .env file and fills the base configuration.
func LoadConfig(cfg *Config) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error loading config file. Loading defaults instead")
	}

	portString, ok := viper.Get("PORT").(string)
	if !ok {
		portString = "4000"
		log.Println("could not find port. Defaulting to port 4000.")
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		port = 4000
		log.Println("could not convert port to int. Defaulting to port 4000.")
	}
	cfg.Port = port

	env, ok := viper.Get("ENVIRONMENT").(string)
	if !ok {
		env = "development"
		log.Println("could not find environment. Defaulting to 'development'")
	}
	cfg.Env = env

	dsn, ok := viper.Get("DSN").(string)
	if !ok {
		dsn = "mongodb://localhost:27017"
		log.Println("could not find dsn. Defaulting to 'mongodb://localhost:27017'")
	}
	cfg.DB.DSN = dsn

	jwt, ok := viper.Get("JWT_TOKEN_SECRET").(string)
	if !ok {
		jwt = "wonderfulsecretphrase"
		log.Println("could not find jwt token secret.",
			"Defaulting to 'wonderfulsecretphrase'")
	}
	cfg.JWT = []byte(jwt)
}

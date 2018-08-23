package main

import (
	"log"

	"github.com/caarlos0/env"
)

// Config struct for config values.
type Config struct {
	AuthToken string `env:"AUTH_TOKEN"`
	Port      string `env:"PORT" envDefault:"7000"`
}

// getConfigFromEnv sets up config by reading environment variables.
func getConfigFromEnv() Config {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

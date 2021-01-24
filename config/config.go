package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Config - app confiuration
type Config struct {
	Redis struct {
		Addr string `envconfig:"REDIS_ADDR"`
	}
	Auth struct {
		JWTKey string `envconfig:"JWT_KEY"`
	}
}

// New config
func New() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}

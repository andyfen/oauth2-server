package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config - app confiuration
type Config struct {
	Domain struct {
		URL string `envconfig:"DOMAIN_URL"`
	}
	Server struct {
		HOST string `envconfig:"SERVER_HOST"`
		PORT string `envconfig:"SERVER_PORT"`
	}
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

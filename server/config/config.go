package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config - app confiuration
type Config struct {
	DomainURL string `json:"domain_url"`
	HOST      string `json:"host"`
	PORT      string `json:"port"`
	RedisAddr string `json:"redis_url"`
	JWTKey    string `json:"jwt_key"`
}

// New config
func New() *Config {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := Config{
		DomainURL: os.Getenv("DOMAIN_URL"),
		HOST:      os.Getenv("HOST"),
		PORT:      os.Getenv("PORT"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		JWTKey:    os.Getenv("JWT_KEY"),
	}

	printRespJSON(cfg)

	return &cfg
}

func printRespJSON(resp interface{}) {
	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return
	}

	fmt.Println(string(b))
}

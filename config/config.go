package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiToken string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return &Config{
		ApiToken: getRequiredStringEnv("API_TOKEN"),
	}
}

func getRequiredStringEnv(key string) string {
	keyValue, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Обязательная переменная %s не найдена", key)
	}
	return keyValue
}

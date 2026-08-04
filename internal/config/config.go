package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	DSN       string
	JWT_SECRET string
}

func LoadEnv() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PORT:      os.Getenv("PORT"),
		DSN:       os.Getenv("DSN"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}
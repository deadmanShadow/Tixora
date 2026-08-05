package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Dsn       string
	JwtSecret string
}

func LoadEnv() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found or couldn't be loaded, reading system environment variables")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = os.Getenv("Dsn")
	}
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}
	if dsn == "" {
		dsn = os.Getenv("POSTGRES_URL")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = os.Getenv("JwtSecret")
	}
	if jwtSecret == "" {
		jwtSecret = os.Getenv("jwt_secret")
	}
	if jwtSecret == "" {
		log.Println(" JWT_SECRET is not set in environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("Port")
	}
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:      port,
		Dsn:       dsn,
		JwtSecret: jwtSecret,
	}
}
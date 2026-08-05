package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg *Config) *gorm.DB {
	dsn := cfg.Dsn
	if dsn == "" {
		log.Fatal("❌ Database DSN/DATABASE_URL is empty. Please check your .env file or environment variables.")
	}

	// Configure GORM logger with higher SlowThreshold (2s) to prevent noisy log spam for cloud databases (NeonDB)
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             2 * time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("Database connection successful")
	return db
}
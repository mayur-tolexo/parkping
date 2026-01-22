package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"parkping/internal/model"
)

func Connect() *gorm.DB {
	env := os.Getenv("APP_ENV")

	var (
		database *gorm.DB
		err      error
	)

	switch env {
	case "dev", "test":
		log.Println("Using in-memory SQLite database")
		database, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	default:
		log.Println("Using PostgreSQL database")
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_SSLMODE"),
		)
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if env == "dev" || env == "test" {
		sqlDB, _ := database.DB()
		sqlDB.SetMaxOpenConns(1)
	}

	// Auto-migrate schema
	err = database.AutoMigrate(
		&model.User{},
		&model.Vehicle{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	return database
}

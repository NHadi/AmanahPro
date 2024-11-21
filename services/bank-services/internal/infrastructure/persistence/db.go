package persistence

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// InitializeDB initializes the database connection using GORM and returns the DB instance.
func InitializeDB() (*gorm.DB, error) {
	// Check if running in Docker (using an environment variable)
	envFilePath := ".env" // Default path
	if _, isInDocker := os.LookupEnv("DOCKER_ENV"); isInDocker {
		envFilePath = "/app/.env" // Path for Docker container
	}

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the database URL from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Open database connection
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	return db, nil
}

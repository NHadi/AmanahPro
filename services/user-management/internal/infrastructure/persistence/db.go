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
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
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

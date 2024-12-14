package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	DatabaseURL      string
	RabbitMQURL      string
	RedisURL         string
	RedisPassword    string
	ElasticsearchURL string
	JWTSecret        string
	APP_ENV          string
	CONCURRENCY      string
	PREFETCH_COUNT   string
}

func LoadConfig() (*Config, error) {
	// Determine the runtime environment
	envFilePath := determineEnvFilePath("../../.env.local")

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize configuration
	return &Config{
		Port:             getEnv("BA_SERVICES_PORT", "8087"),
		DatabaseURL:      getEnv("DATABASE_BA_URL", ""),
		RabbitMQURL:      getEnv("RABBITMQ_URL", ""),
		RedisURL:         getEnv("REDIS_URL", ""),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		ElasticsearchURL: getEnv("ELASTICSEARCH_URL", ""),
		JWTSecret:        getEnv("JWT_SECRET", ""),
		APP_ENV:          getEnv("APP_ENV", ""),
		CONCURRENCY:      getEnv("CONCURRENCY", ""),
		PREFETCH_COUNT:   getEnv("PREFETCH_COUNT", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// determineEnvFilePath determines the correct .env file path based on runtime environment
func determineEnvFilePath(defaultpath string) string {
	// Default to production environment
	if fileExists("/root/AmanahPro/.env") {
		return "/root/AmanahPro/.env"
	}

	// Default to Docker environment if Docker-specific path exists
	if fileExists("/app/.env") {
		return "/app/.env"
	}

	// Development fallback
	return defaultpath
}

// fileExists checks if a file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

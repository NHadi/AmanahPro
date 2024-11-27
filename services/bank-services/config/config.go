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
		Port:             getEnv("BANK_SERVICES_PORT", "8082"),
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		RabbitMQURL:      getEnv("RABBITMQ_URL", ""),
		RedisURL:         getEnv("REDIS_URL", ""),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		ElasticsearchURL: getEnv("ELASTICSEARCH_URL", ""),
		JWTSecret:        getEnv("JWT_SECRET", ""),
		APP_ENV:          getEnv("APP_ENV", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// determineEnvFilePath determines the correct .env file path based on runtime environment
func determineEnvFilePath(localEnvPath string) string {
	// Check for Docker environment
	if isDockerEnvironment() {
		return "/app/.env" // Docker container path
	}

	if fileExists(localEnvPath) {
		return localEnvPath
	}

	// Fallback to production path
	return "/root/AmanahPro/.env" // Production path (e.g., VM)
}

// isDockerEnvironment checks if the application is running inside Docker
func isDockerEnvironment() bool {
	// Docker containers usually have a cgroup file
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	// Alternatively, check for specific Docker files
	cgroupPath := "/proc/1/cgroup"
	if fileExists(cgroupPath) {
		return true
	}
	return false
}

// fileExists checks if a file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

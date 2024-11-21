package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/NHadi/AmanahPro-common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// ServiceConfig holds the mapping of paths to service URLs
type ServiceConfig struct {
	ServiceMap map[string]string
}

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	// Check if running in Docker (using an environment variable)
	envFilePath := ".env" // Default path
	if _, isInDocker := os.LookupEnv("DOCKER_ENV"); isInDocker {
		envFilePath = "/app/.env" // Path for Docker container
	}

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Main, Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// Define service configuration
	serviceConfig := ServiceConfig{
		ServiceMap: map[string]string{
			"/user-management": os.Getenv("SERVICES_USER_MANAGEMENT"),
			"/bank_services":   os.Getenv("K"),
			// Add other services here as needed
		},
	}

	// Apply JWT Authentication Middleware, with exclusion for the login route
	r.Use(func(c *gin.Context) {
		// Bypass JWT middleware for /user-management/login
		if c.Request.URL.Path == "/user-management/login" && c.Request.Method == http.MethodPost {
			c.Next()
			return
		}

		// Apply JWTAuthMiddleware for other routes
		middleware.JWTAuthMiddleware(jwtSecret)(c)

	})

	// Route for dynamic proxying
	r.Any("/*proxyPath", func(c *gin.Context) {
		proxyPath := c.Param("proxyPath")
		log.Printf("Received request for: %s", proxyPath)

		// Check for the service in the configuration
		for route, serviceURL := range serviceConfig.ServiceMap {
			if strings.HasPrefix(proxyPath, route) {
				// Log the service URL to ensure it's loaded correctly
				log.Printf("Using service URL: %s for route: %s", serviceURL, route)

				// Parse the service URL
				url, err := url.Parse(serviceURL)
				if err != nil {
					log.Printf("Error parsing service URL: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Service not available"})
					return
				}

				// Create the reverse proxy
				proxy := httputil.NewSingleHostReverseProxy(url)

				// Update the request URL and host
				c.Request.URL.Host = url.Host
				c.Request.URL.Scheme = url.Scheme
				c.Request.URL.Path = strings.TrimPrefix(proxyPath, route)
				c.Request.Host = url.Host

				// Log the forwarding request
				log.Printf("Forwarding request to %s%s", url.String(), c.Request.URL.Path)

				// Serve the request
				proxy.ServeHTTP(c.Writer, c.Request)
				return
			}
		}

		// If no matching route found, return a 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
	})

	// Start the API Gateway
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

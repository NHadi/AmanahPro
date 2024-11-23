package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	middleware "github.com/NHadi/AmanahPro-common/middleware"

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
	envFilePath := "../.env.local" // Default path
	if _, isInDocker := os.LookupEnv("DOCKER_ENV"); isInDocker {
		envFilePath = "/app/.env" // Path for Docker container
	}

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// Define service configuration
	serviceConfig := ServiceConfig{
		ServiceMap: map[string]string{
			"/user-management": os.Getenv("SERVICES_USER_MANAGEMENT"),
			"/bank_services":   os.Getenv("SERVICES_BANK"),
		},
	}

	// Log the loaded service configuration
	log.Printf("Service Configuration: %+v", serviceConfig.ServiceMap)

	// Apply JWT Authentication Middleware, excluding specific routes
	r.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/user-management/login" && c.Request.Method == http.MethodPost {
			c.Next()
			return
		}

		middleware.JWTAuthMiddleware(jwtSecret)(c)
	})

	// Route for dynamic proxying
	r.Any("/*proxyPath", func(c *gin.Context) {
		proxyPath := c.Param("proxyPath")
		log.Printf("Received request for: %s", proxyPath)

		// Match the request path with the service map
		for route, serviceURL := range serviceConfig.ServiceMap {
			if strings.HasPrefix(proxyPath, route) {
				log.Printf("Using service URL: %s for route: %s", serviceURL, route)

				// Parse the service URL
				parsedURL, err := url.Parse(serviceURL)
				if err != nil {
					log.Printf("Error parsing service URL: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Service not available"})
					return
				}

				// Create and configure the reverse proxy
				proxy := httputil.NewSingleHostReverseProxy(parsedURL)
				proxy.Director = func(req *http.Request) {
					req.URL.Scheme = parsedURL.Scheme
					req.URL.Host = parsedURL.Host
					req.URL.Path = strings.TrimPrefix(proxyPath, route)
					req.Host = parsedURL.Host
					log.Printf("Proxying request: %s %s", req.Method, req.URL.String())
				}

				proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
					log.Printf("Error during proxying: %v", err)
					rw.WriteHeader(http.StatusBadGateway)
					rw.Write([]byte("Bad Gateway"))
				}

				// Serve the proxied request
				proxy.ServeHTTP(c.Writer, c.Request)
				return
			}
		}

		// Return 404 if no matching route is found
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
	})

	// Start the API Gateway
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

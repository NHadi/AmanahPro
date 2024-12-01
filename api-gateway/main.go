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

	// Determine the runtime environment
	envFilePath := determineEnvFilePath("../.env.local")

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	r := gin.Default()

	// Handle CORS directly in the router
	r.Use(func(c *gin.Context) {
		// List of allowed origins
		allowedOrigins := []string{
			"https://amanahpro.pilarasamandiri.com",
			"http://localhost:8090",
		}

		origin := c.Request.Header.Get("Origin")
		isAllowed := false

		// Check if the origin is in the allowed list
		for _, o := range allowedOrigins {
			if origin == o {
				isAllowed = true
				break
			}
		}

		// Set CORS headers if the origin is allowed
		if isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Define service configuration
	serviceConfig := ServiceConfig{
		ServiceMap: map[string]string{
			"/user-management":               os.Getenv("SERVICES_USER_MANAGEMENT"),
			"/bank_services":                 os.Getenv("SERVICES_BANK"),
			"/project_management_services":   os.Getenv("PROJECT_MANAGEMNET_SERVICES"),
			"/breakdown_management_services": os.Getenv("BREAKDOWN_MANAGEMNET_SERVICES"),
			"/sph_services":                  os.Getenv("SPH_SERVICES"),
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
	// Start server
	port := os.Getenv("GATEWAY_PORT")
	if os.Getenv("APP_ENV") == "PRODUCTION" {
		// Serve HTTPS
		certFile := "/etc/letsencrypt/live/amanahpro.pilarasamandiri.com/fullchain.pem"
		keyFile := "/etc/letsencrypt/live/amanahpro.pilarasamandiri.com/privkey.pem"

		if err := r.RunTLS(":"+port, certFile, keyFile); err != nil {
			log.Fatalf("Failed to start HTTPS server: %v", err)
		}
	} else {

		log.Printf("Server running at http://localhost:%s", port)
		log.Fatal(r.Run(":" + port))
	}

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

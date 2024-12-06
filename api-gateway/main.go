package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"runtime/debug"
	"strings"
	"time"

	middleware "github.com/NHadi/AmanahPro-common/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// ServiceConfig holds the mapping of paths to service URLs
type ServiceConfig struct {
	ServiceMap map[string]string
}

const (
	defaultPort = "8080"
	logDir      = "log"
)

const TraceIDHeader = "X-Trace-ID"

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

	defer recoverFromPanic()

	configureLogging()
	log.Println("Starting API Gateway...")

	r := gin.Default()

	// Apply Trace ID middleware
	r.Use(TraceIDMiddleware())

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
			"/spk_services":                  os.Getenv("SPK_SERVICES"),
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
		traceID := c.GetString(TraceIDHeader) // Retrieve Trace ID
		proxyPath := c.Param("proxyPath")
		log.Printf("TraceID: %s - Received request for: %s", traceID, proxyPath)

		// Match the request path with the service map
		for route, serviceURL := range serviceConfig.ServiceMap {
			if strings.HasPrefix(proxyPath, route) {
				log.Printf("TraceID: %s - Using service URL: %s for route: %s", traceID, serviceURL, route)

				// Parse the service URL
				parsedURL, err := url.Parse(serviceURL)
				if err != nil {
					log.Printf("TraceID: %s - Error parsing service URL: %v", traceID, err)
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
					// Propagate the trace ID
					req.Header.Set(TraceIDHeader, traceID)
					log.Printf("TraceID: %s - Proxying request: %s %s", traceID, req.Method, req.URL.String())
				}

				proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
					log.Printf("TraceID: %s - Error during proxying: %v", traceID, err)
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

// TraceIDMiddleware ensures every request has a trace ID
func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(TraceIDHeader)
		if traceID == "" {
			// Generate a new trace ID if not provided
			traceID = uuid.New().String()
		}
		// Set the trace ID in the context and response headers
		c.Set(TraceIDHeader, traceID)
		c.Writer.Header().Set(TraceIDHeader, traceID)
		c.Next()
	}
}

// determineEnvFilePath determines the correct .env file path based on runtime environment
func determineEnvFilePath(localEnvPath string) string {
	if isDockerEnvironment() {
		return "/app/.env" // Docker container path
	}
	if fileExists(localEnvPath) {
		return localEnvPath
	}
	return "/root/AmanahPro/.env" // Production path
}

// isDockerEnvironment checks if the application is running inside Docker
func isDockerEnvironment() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	cgroupPath := "/proc/1/cgroup"
	return fileExists(cgroupPath)
}

// fileExists checks if a file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// configureLogging sets up daily logging into a file.
func configureLogging() {
	// Check if the log directory exists; if not, create it
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	// Define the log file name based on the current date
	logFileName := fmt.Sprintf("%s/Log%s.log", logDir, time.Now().Format("20060102"))
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Check the APP_ENV environment variable
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "PRODUCTION" {
		// If not production, write logs to both the terminal and the log file
		log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	} else {
		// If production, write logs only to the log file
		log.SetOutput(logFile)
	}

	// Set log flags for consistent log formatting
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("Logging initialized: %s", logFileName)
	log.Printf("Environment: %s", appEnv)
}

// recoverFromPanic recovers from panics and logs the stack trace.
func recoverFromPanic() {
	if r := recover(); r != nil {
		log.Printf("Application panic recovered: %v", r)
		log.Printf("Stack trace: %s", debug.Stack())
	}
}

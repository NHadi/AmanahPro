package main

import (
	_ "AmanahPro/services/user-management/docs" // Swagger docs
	"AmanahPro/services/user-management/internal/application/services"
	domainServices "AmanahPro/services/user-management/internal/domain/services"
	"AmanahPro/services/user-management/internal/handlers"
	"AmanahPro/services/user-management/internal/infrastructure/persistence"
	"AmanahPro/services/user-management/internal/infrastructure/repositories"
	"log"
	"net/http"
	"os"

	"github.com/NHadi/AmanahPro-common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/gin-swagger"
)

const defaultPort = "8081"

// @title User Management API
// @version 1.0
// @description This is the User Management API documentation with role and permission management.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token with "Bearer " prefix, e.g., "Bearer <token>"
// @host localhost:8081
// @BasePath /
func main() {

	// Determine the runtime environment
	envFilePath := determineEnvFilePath("../../.env.local")

	// Load environment variables
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Main, Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	// Initialize DB
	db, err := persistence.InitializeDB(os.Getenv("DATABASE_AUTH_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	menuRepo := repositories.NewMenuRepository(db)
	roleMenuRepo := repositories.NewRoleMenuRepository(db)
	userRoleRepo := repositories.NewUserRoleRepository(db)

	// Initialize domain services
	roleAssignmentService := domainServices.NewRoleAssignmentService(userRoleRepo)

	// Initialize application services
	userService := services.NewUserService(userRepo, roleAssignmentService)
	roleService := services.NewRoleService(roleRepo)
	menuService := services.NewMenuService(menuRepo, roleMenuRepo)
	permissionService := services.NewPermissionService(roleMenuRepo)

	// Initialize handlers
	loginHandler := handlers.NewLoginHandler(userService, jwtSecret)
	userHandler := handlers.NewUserHandler(userService)
	roleHandler := handlers.NewRoleHandler(roleService)
	menuHandler := handlers.NewMenuHandler(menuService)
	permissionHandler := handlers.NewPermissionHandler(permissionService)

	// Initialize Gin router
	r := gin.Default()

	// Initialize common logger
	logger, err := middleware.InitializeLogger("user-mangement", os.Getenv("ELASTICSEARCH_URL"), "user-management-logs")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	// Attach common logging middleware
	r.Use(middleware.GinLoggingMiddleware(logger))

	// Middleware to log requests
	r.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
	// Public route for login
	r.POST("/login", func(c *gin.Context) {
		loginHandler.Login(c.Writer, c.Request)
	})

	// Add Health Check Endpoint
	r.GET("/health", func(c *gin.Context) {
		// Perform health checks for dependencies
		healthChecks := map[string]string{}

		healthChecks["database"] = "healthy"

		// Respond with health status
		statusCode := http.StatusOK

		c.JSON(statusCode, gin.H{
			"status":  "healthy",
			"details": healthChecks,
		})
	})
	// Group for protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware(jwtSecret))

	// Menu Routes - Only accessible with a valid JWT token
	api.GET("/menus/:roleID", menuHandler.GetAccessibleMenus)
	api.POST("/menus", menuHandler.CreateMenu)
	// User Routes
	api.POST("/users", userHandler.CreateUser)                           // Create a new user
	api.GET("/users/organizations", userHandler.LoadUsersByOrganization) // Get users by organization
	api.PUT("/users/:user_id", userHandler.UpdateUser)                   // Update an existing user

	// Role Routes
	api.POST("/roles", roleHandler.CreateRole)
	api.GET("/roles", roleHandler.GetRoles)

	// Permission Routes
	api.POST("/permissions/assign", permissionHandler.AssignPermission)

	// Swagger documentation route
	r.GET("/swagger/*any", httpSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from environment or default to 8081
	port := os.Getenv("USER_MANAGEMENT_PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(r.Run(":" + port))
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

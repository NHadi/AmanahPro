package main

import (
	"log"
	"net/http"
	"os"

	_ "AmanahPro/services/user-management/docs" // Swagger docs
	"AmanahPro/services/user-management/internal/application/services"
	domainServices "AmanahPro/services/user-management/internal/domain/services"
	"AmanahPro/services/user-management/internal/handlers"
	"AmanahPro/services/user-management/internal/infrastructure/persistence"
	"AmanahPro/services/user-management/internal/infrastructure/repositories"
	"AmanahPro/services/user-management/internal/middleware"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const defaultPort = "8080"

// @title User Management API
// @version 1.0
// @description This is the User Management API documentation with role and permission management.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token with "Bearer " prefix, e.g., "Bearer <token>"
// @host localhost:8080
// @BasePath /
func main() {

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	// Initialize DB
	db, err := persistence.InitializeDB()
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
	menuService := services.NewMenuService(menuRepo)
	permissionService := services.NewPermissionService(roleMenuRepo)

	// Initialize handlers
	loginHandler := handlers.NewLoginHandler(userService, jwtSecret)
	userHandler := handlers.NewUserHandler(userService)
	roleHandler := handlers.NewRoleHandler(roleService)
	menuHandler := handlers.NewMenuHandler(menuService)
	permissionHandler := handlers.NewPermissionHandler(permissionService)

	// Initialize router
	r := mux.NewRouter()

	// Public route for login
	r.HandleFunc("/login", loginHandler.Login).Methods("POST")
	// Define User Routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// Define Role Routes
	r.HandleFunc("/roles", roleHandler.CreateRole).Methods("POST")
	r.HandleFunc("/roles", roleHandler.GetRoles).Methods("GET")

	// Define Permission Routes
	r.HandleFunc("/permissions/assign", permissionHandler.AssignPermission).Methods("POST")
	// Group protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware(jwtSecret))
	api.Use(middleware.LoggingMiddleware) // Logging middleware

	// Menu Routes - Only accessible with a valid JWT token
	api.HandleFunc("/menus/{roleID:[0-9]+}", menuHandler.GetAccessibleMenus).Methods("GET")
	api.HandleFunc("/menus", menuHandler.CreateMenu).Methods("POST")
	// Swagger documentation route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

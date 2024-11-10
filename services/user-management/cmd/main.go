package main

import (
	"log"
	"net/http"

	_ "AmanahPro/services/user-management/docs" // Swagger docs
	"AmanahPro/services/user-management/internal/application/services"
	domainServices "AmanahPro/services/user-management/internal/domain/services"
	"AmanahPro/services/user-management/internal/handlers"
	"AmanahPro/services/user-management/internal/infrastructure/persistence"
	"AmanahPro/services/user-management/internal/infrastructure/repositories"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User Management API
// @version 1.0
// @description This is the User Management API documentation with role and permission management.
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize DB
	db, err := persistence.InitializeDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories and services
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	menuRepo := repositories.NewMenuRepository(db)
	roleMenuRepo := repositories.NewRoleMenuRepository(db)
	userRoleRepo := repositories.NewUserRoleRepository(db)
	roleAssignmentService := domainServices.NewRoleAssignmentService(userRoleRepo)
	userService := services.NewUserService(userRepo, roleAssignmentService)
	roleService := services.NewRoleService(roleRepo)
	menuService := services.NewMenuService(menuRepo)
	permissionService := services.NewPermissionService(roleMenuRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	roleHandler := handlers.NewRoleHandler(roleService)
	menuHandler := handlers.NewMenuHandler(menuService)
	permissionHandler := handlers.NewPermissionHandler(permissionService) // Permission handler

	// Initialize router
	r := mux.NewRouter()

	// User Routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// Role Routes
	r.HandleFunc("/roles", roleHandler.CreateRole).Methods("POST")
	r.HandleFunc("/roles", roleHandler.GetRoles).Methods("GET")

	// Menu Routes
	r.HandleFunc("/menus/{roleID}", menuHandler.GetAccessibleMenus).Methods("GET")
	r.HandleFunc("/menus", menuHandler.CreateMenu).Methods("POST")

	// Permission Routes
	r.HandleFunc("/permissions/assign", permissionHandler.AssignPermission).Methods("POST")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

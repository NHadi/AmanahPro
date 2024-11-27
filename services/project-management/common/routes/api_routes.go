package routes

import (
	"AmanahPro/services/project-management/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes registers all API routes with their respective handlers
func RegisterAPIRoutes(api *gin.RouterGroup, handlers *handlers.Handlers) {
	// Project routes
	projects := api.Group("/projects")
	{
		projects.POST("", handlers.Project.CreateProject)                      // Create a project
		projects.PUT("", handlers.Project.UpdateProject)                       // Update a project
		projects.DELETE("/:id", handlers.Project.DeleteProject)                // Delete a project by ID
		projects.GET("/search", handlers.Project.SearchProjectsByOrganization) // Search projects
	}
}

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
		projects.POST("", handlers.Project.CreateProject)                      // Create an PROJECT
		projects.PUT("/:project_id", handlers.Project.UpdateProject)           // Update an PROJECT
		projects.DELETE("/:project_id", handlers.Project.DeleteProject)        // Delete an PROJECT
		projects.GET("/search", handlers.Project.SearchProjectsByOrganization) // Search projects
	}
}

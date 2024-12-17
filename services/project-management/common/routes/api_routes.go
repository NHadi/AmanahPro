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
		projects.POST("", handlers.Project.CreateProject)                                // Create an PROJECT
		projects.PUT("/:project_id", handlers.Project.UpdateProject)                     // Update an PROJECT
		projects.DELETE("/:project_id", handlers.Project.DeleteProject)                  // Delete an PROJECT
		projects.GET("/search", handlers.Project.SearchProjectsByOrganization)           // Search projects
		projects.GET("/:project_id/users", handlers.Project.SearchProjectUsersByProject) // Search projects

		projects.POST("/:project_id/assign-user", handlers.Project.AssignUser)
		projects.DELETE("/:project_id/unassign-user", handlers.Project.UnAssignUser)
		projects.PUT("/:project_id/change-user/:user_id", handlers.Project.ChangeUser)
	}

	// Financial routes
	financial := api.Group("/project-financials")
	{
		financial.POST("", handlers.ProjectFinancial.CreateProjectFinancial)                        // Create a financial record
		financial.PUT("/:financial_id", handlers.ProjectFinancial.UpdateProjectFinancial)           // Update a financial record
		financial.DELETE("/:financial_id", handlers.ProjectFinancial.DeleteProjectFinancial)        // Delete a financial record
		financial.GET("/:financial_id", handlers.ProjectFinancial.GetProjectFinancialByID)          // Get a financial record by ID
		financial.GET("/project/:project_id", handlers.ProjectFinancial.GetAllFinancialByProjectID) // Get all financial records for a project
	}
}

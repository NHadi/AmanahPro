package routes

import (
	"AmanahPro/services/ba-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes registers all API routes with their respective handlers
func RegisterAPIRoutes(api *gin.RouterGroup, handlers *handlers.Handlers) {
	// BA routes
	bas := api.Group("/ba")
	{
		bas.GET("/filter", handlers.BA.FilterBAs)   // Filter BAs
		bas.POST("", handlers.BA.CreateBA)          // Create a BA
		bas.PUT("/:ba_id", handlers.BA.UpdateBA)    // Update a BA
		bas.DELETE("/:ba_id", handlers.BA.DeleteBA) // Delete a BA

		// BA Sections
		sections := bas.Group("/:ba_id/sections")
		{
			sections.POST("", handlers.BA.CreateBASection)               // Create a section
			sections.PUT("/:section_id", handlers.BA.UpdateBASection)    // Update a section
			sections.DELETE("/:section_id", handlers.BA.DeleteBASection) // Delete a section

			// BA Details
			details := sections.Group("/:section_id/details")
			{
				details.POST("", handlers.BA.CreateBADetail)              // Create a detail
				details.PUT("/:detail_id", handlers.BA.UpdateBADetail)    // Update a detail
				details.DELETE("/:detail_id", handlers.BA.DeleteBADetail) // Delete a detail

				// BA Progress
				progress := details.Group("/:detail_id/progress")
				{
					progress.POST("", handlers.BA.CreateBAProgress)                // Create progress
					progress.PUT("/:progress_id", handlers.BA.UpdateBAProgress)    // Update progress
					progress.DELETE("/:progress_id", handlers.BA.DeleteBAProgress) // Delete progress
				}
			}
		}
	}
}

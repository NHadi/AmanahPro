package routes

import (
	"AmanahPro/services/ba-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes registers all API routes with their respective handlers
func RegisterAPIRoutes(api *gin.RouterGroup, handlers *handlers.Handlers) {
	// SPK routes
	bas := api.Group("/ba")
	{
		bas.GET("/filter", handlers.Spk.FilterSpks)   // Filter SPKs
		bas.POST("", handlers.Spk.CreateSpk)          // Create an SPK
		bas.PUT("/:ba_id", handlers.Spk.UpdateSpk)    // Update an SPK
		bas.DELETE("/:ba_id", handlers.Spk.DeleteSpk) // Delete an SPK

		// SPK Sections
		sections := bas.Group("/:ba_id/sections")
		{
			sections.POST("", handlers.Spk.CreateSpkSection)               // Create a section
			sections.PUT("/:section_id", handlers.Spk.UpdateSpkSection)    // Update a section
			sections.DELETE("/:section_id", handlers.Spk.DeleteSpkSection) // Delete a section

			// SPK Details
			details := sections.Group("/:section_id/details")
			{
				details.POST("", handlers.Spk.CreateSpkDetail)              // Create a detail
				details.PUT("/:detail_id", handlers.Spk.UpdateSpkDetail)    // Update a detail
				details.DELETE("/:detail_id", handlers.Spk.DeleteSpkDetail) // Delete a detail
			}
		}
	}
}

package routes

import (
	"AmanahPro/services/sph-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes registers all API routes with their respective handlers
func RegisterAPIRoutes(api *gin.RouterGroup, handlers *handlers.Handlers) {
	// SPH routes
	sphs := api.Group("/sph")
	{
		sphs.GET("/filter", handlers.Sph.FilterSphs) // Filter SPHs
		sphs.POST("", handlers.Sph.CreateSph)        // Create an SPH
		sphs.POST("/import", handlers.Sph.ImportSphFromExcel)
		sphs.PUT("/:sph_id", handlers.Sph.UpdateSph)    // Update an SPH
		sphs.DELETE("/:sph_id", handlers.Sph.DeleteSph) // Delete an SPH

		// SPH Sections
		sections := sphs.Group("/:sph_id/sections")
		{
			sections.POST("", handlers.Sph.CreateSphSection)               // Create a section
			sections.PUT("/:section_id", handlers.Sph.UpdateSphSection)    // Update a section
			sections.DELETE("/:section_id", handlers.Sph.DeleteSphSection) // Delete a section

			// SPH Details
			details := sections.Group("/:section_id/details")
			{
				details.POST("", handlers.Sph.CreateSphDetail)              // Create a detail
				details.PUT("/:detail_id", handlers.Sph.UpdateSphDetail)    // Update a detail
				details.DELETE("/:detail_id", handlers.Sph.DeleteSphDetail) // Delete a detail
			}
		}
	}
}

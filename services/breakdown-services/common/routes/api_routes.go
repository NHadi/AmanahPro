package routes

import (
	"AmanahPro/services/breakdown-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes registers all API routes with their respective handlers
func RegisterAPIRoutes(api *gin.RouterGroup, handlers *handlers.Handlers) {
	// Breakdown routes
	breakdowns := api.Group("/breakdowns")
	{
		breakdowns.GET("/filter", handlers.Breakdown.FilterBreakdowns)          // Filter breakdowns
		breakdowns.POST("", handlers.Breakdown.CreateBreakdown)                 // Create a breakdown
		breakdowns.PUT("/:breakdown_id", handlers.Breakdown.UpdateBreakdown)    // Update a breakdown
		breakdowns.DELETE("/:breakdown_id", handlers.Breakdown.DeleteBreakdown) // Delete a breakdown

		// Breakdown Sections
		sections := breakdowns.Group("/:breakdown_id/sections")
		{
			sections.POST("", handlers.Breakdown.CreateBreakdownSection)               // Create a section
			sections.PUT("/:section_id", handlers.Breakdown.UpdateBreakdownSection)    // Update a section
			sections.DELETE("/:section_id", handlers.Breakdown.DeleteBreakdownSection) // Delete a section

			// Breakdown Items
			items := sections.Group("/:section_id/items")
			{
				items.POST("", handlers.Breakdown.CreateBreakdownItem)            // Create an item
				items.PUT("/:item_id", handlers.Breakdown.UpdateBreakdownItem)    // Update an item
				items.DELETE("/:item_id", handlers.Breakdown.DeleteBreakdownItem) // Delete an item
			}
		}
	}
}

package routes

import (
	"AmanahPro/services/bank-services/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes registers all API routes with their respective handlers
func RegisterAPIRoutes(api *gin.RouterGroup, handlers *handlers.Handlers) {
	// Upload Routes
	api.POST("/upload", handlers.Upload.UploadBatch)

	// Transaction Routes
	api.GET("/transactions", handlers.Transaction.GetTransactionsByBankAndPeriod)

	// Reconciliation Routes
	api.POST("/reconcile", handlers.Reconciliation.ReconcileHandler)
}

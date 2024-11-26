package handlers

import (
	"AmanahPro/services/bank-services/internal/application/services"

	"github.com/gin-gonic/gin"
)

// ReconciliationHandler handles reconciliation-related requests
type ReconciliationHandler struct {
	ReconciliationService *services.ReconciliationService
}

// NewReconciliationHandler creates a new instance of ReconciliationHandler
func NewReconciliationHandler(reconciliationService *services.ReconciliationService) *ReconciliationHandler {
	return &ReconciliationHandler{
		ReconciliationService: reconciliationService,
	}
}

// ReconcileHandler triggers the reconciliation process
// @Summary Trigger Reconciliation
// @Description Manually trigger the reconciliation process
// @Tags Reconciliation
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/reconcile [post]
func (h *ReconciliationHandler) ReconcileHandler(c *gin.Context) {
	err := h.ReconciliationService.PerformReconciliation()
	if err != nil {
		c.JSON(500, gin.H{"message": "Reconciliation failed", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Reconciliation completed successfully"})
}

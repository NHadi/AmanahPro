package handlers

import (
	"AmanahPro/services/bank-services/internal/application/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

// GetTransactionsByBankAndPeriod handles the request to fetch transactions
// @Summary Get Transactions by Bank ID and Period
// @Description Fetch transactions by bank ID and date range
// @Tags Permissions
// @Security BearerAuth
// @Param bank_id query int true "Bank Account ID"
// @Param periode_start query string true "Start date of the period (YYYY-MM-DD)"
// @Param periode_end query string true "End date of the period (YYYY-MM-DD)"
// @Produce json
// @Success 200 {object} []map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/transactions [get]
func (h *TransactionHandler) GetTransactionsByBankAndPeriod(c *gin.Context) {
	// Parse query parameters
	bankIDStr := c.Query("bank_id")
	bankID, err := strconv.ParseUint(bankIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bank_id"})
		return
	}

	periodeStartStr := c.Query("periode_start")
	periodeStart, err := time.Parse("2006-01-02", periodeStartStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid periode_start format"})
		return
	}

	periodeEndStr := c.Query("periode_end")
	periodeEnd, err := time.Parse("2006-01-02", periodeEndStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid periode_end format"})
		return
	}

	// Call the service method
	transactions, err := h.transactionService.GetTransactionsByBankAndPeriod(uint(bankID), periodeStart, periodeEnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

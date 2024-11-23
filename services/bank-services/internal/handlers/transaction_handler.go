package handlers

import (
	"AmanahPro/services/bank-services/internal/application/services"
	"net/http"
	"strconv"

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
// @Summary Get Transactions by Bank ID and optional Year
// @Description Fetch transactions by bank ID and optional year
// @Tags Transactions
// @Security BearerAuth
// @Param bank_id query int true "Bank Account ID"
// @Param year query int false "Year (optional)"
// @Produce json
// @Success 200 {object} []dto.BankAccountTransactionDTO
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string
// @Router /api/transactions [get]
func (h *TransactionHandler) GetTransactionsByBankAndPeriod(c *gin.Context) {
	// Parse bank_id from query parameters
	bankIDStr := c.Query("bank_id")
	bankID, err := strconv.ParseUint(bankIDStr, 10, 64)
	if err != nil || bankID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing bank_id"})
		return
	}

	// Parse year from query parameters (optional)
	yearStr := c.Query("year")
	var year *int
	if yearStr != "" {
		yearValue, err := strconv.Atoi(yearStr)
		if err != nil || yearValue < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year format"})
			return
		}
		year = &yearValue
	}

	// Call the service method
	transactions, err := h.transactionService.GetTransactionsByBankAndPeriod(uint(bankID), year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

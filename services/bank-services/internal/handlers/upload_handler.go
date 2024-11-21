package handlers

import (
	"AmanahPro/services/bank-services/internal/application/services"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadService   *services.UploadService
	transactionRepo repositories.BankAccountTransactionRepository
	batchRepo       repositories.BatchRepository
}

func NewUploadHandler(
	uploadService *services.UploadService,
	transactionRepo repositories.BankAccountTransactionRepository,
	batchRepo repositories.BatchRepository,
) *UploadHandler {
	return &UploadHandler{
		uploadService:   uploadService,
		transactionRepo: transactionRepo,
		batchRepo:       batchRepo,
	}
}

// UploadBatch handles the CSV upload and processes transactions
// @Summary Upload CSV File for Transactions
// @Description Upload a CSV file for a specific bank account and period
// @Tags Permissions
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param account_id formData int true "Bank Account ID"
// @Param periode_start formData string true "Start date of the period (YYYY-MM-DD)"
// @Param periode_end formData string true "End date of the period (YYYY-MM-DD)"
// @Param uploaded_by formData string true "Uploader's name"
// @Param file formData file true "CSV file"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]interface{}
// @Router /api/upload [post]
func (h *UploadHandler) UploadBatch(c *gin.Context) {
	// Retrieve form data
	accountIDStr := c.PostForm("account_id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account_id"})
		return
	}

	periodeStartStr := c.PostForm("periode_start")
	periodeEndStr := c.PostForm("periode_end")
	uploadedBy := c.PostForm("uploaded_by")
	if uploadedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uploaded_by is required"})
		return
	}

	periodeStart, err := time.Parse("2006-01-02", periodeStartStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid periode_start format"})
		return
	}

	periodeEnd, err := time.Parse("2006-01-02", periodeEndStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid periode_end format"})
		return
	}

	// Check for duplicate batch
	exists, err := h.batchRepo.BatchExists(uint(accountID), periodeStart, periodeEnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate batch existence"})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A batch for this account and period already exists"})
		return
	}

	// Retrieve the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Save the file temporarily
	tempFilePath := "./temp.csv"
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer os.Remove(tempFilePath) // Clean up the file after processing

	// Process the file using the UploadService
	transactions, err := h.uploadService.ParseAndSave(tempFilePath, uint(accountID), periodeStart, periodeEnd, uploadedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message":            "Batch uploaded successfully",
		"transactions_count": len(transactions),
		"uploaded_by":        uploadedBy,
		"account_id":         accountID,
		"periode_start":      periodeStart.Format("2006-01-02"),
		"periode_end":        periodeEnd.Format("2006-01-02"),
	})
}

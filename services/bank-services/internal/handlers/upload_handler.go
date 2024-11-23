package handlers

import (
	"AmanahPro/services/bank-services/internal/application/services"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"net/http"
	"os"
	"strconv"

	jwtModels "github.com/NHadi/AmanahPro-common/models"

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
// @Param year formData int true "Year"
// @Param month formData int true "Month"
// @Param uploaded_by formData string true "Uploader's name"
// @Param file formData file true "CSV file"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]interface{}
// @Router /api/upload [post]
func (h *UploadHandler) UploadBatch(c *gin.Context) {
	// Extract claims from the JWT token
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims, ok := userClaims.(*jwtModels.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Use username or email from claims as uploaded_by
	uploadedBy := claims.Username // or claims.Email

	// Retrieve form data
	accountIDStr := c.PostForm("account_id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account_id"})
		return
	}

	yearStr := c.PostForm("year")
	year, err := strconv.ParseUint(yearStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	monthStr := c.PostForm("month")
	month, err := strconv.ParseUint(monthStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month"})
		return
	}

	// Declare exists explicitly to avoid redeclaration
	var batchExists bool
	batchExists, err = h.batchRepo.BatchExists(uint(accountID), uint(year), uint(month))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate batch existence"})
		return
	}

	if batchExists {
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
	transactions, err := h.uploadService.ParseAndSave(tempFilePath, uint(accountID), uint(year), uint(month), uploadedBy)
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
		"year":               year,
		"month":              month,
	})
}

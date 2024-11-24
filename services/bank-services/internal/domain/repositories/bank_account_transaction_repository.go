package repositories

import (
	"AmanahPro/services/bank-services/internal/application/dto"
	"AmanahPro/services/bank-services/internal/domain/models"
)

type BankAccountTransactionRepository interface {
	InsertWithRollback(batch *models.UploadBatch, transactions []models.BankAccountTransactions) error
	GetTransactionsByBankAndPeriod(organizationID, bankID uint, year *int) ([]dto.BankAccountTransactionDTO, error)
	GetByBatchID(batchID uint) ([]models.BankAccountTransactions, error)
}

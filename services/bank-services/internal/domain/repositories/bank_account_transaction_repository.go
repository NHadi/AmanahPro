package repositories

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"time"
)

type BankAccountTransactionRepository interface {
	InsertWithRollback(batch *models.UploadBatch, transactions []models.BankAccountTransactions) error
	GetTransactionsByBankAndPeriod(bankID uint, periodeStart, periodeEnd time.Time) ([]models.BankAccountTransactions, error)
}

package services

import (
	"AmanahPro/services/bank-services/internal/domain/models"
	"AmanahPro/services/bank-services/internal/domain/repositories"
	"time"
)

type TransactionService struct {
	bankAccountTransactionRepository repositories.BankAccountTransactionRepository
}

func NewTransactionService(elasticsearchRepo repositories.BankAccountTransactionRepository) *TransactionService {
	return &TransactionService{
		bankAccountTransactionRepository: elasticsearchRepo,
	}
}

func (s *TransactionService) GetTransactionsByBankAndPeriod(bankID uint, periodeStart, periodeEnd time.Time) ([]models.BankAccountTransactions, error) {
	// Fetch transactions from Elasticsearch
	return s.bankAccountTransactionRepository.GetTransactionsByBankAndPeriod(bankID, periodeStart, periodeEnd)
}

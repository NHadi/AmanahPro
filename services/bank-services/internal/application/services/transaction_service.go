package services

import (
	"AmanahPro/services/bank-services/internal/application/dto"
	"AmanahPro/services/bank-services/internal/domain/repositories"
)

type TransactionService struct {
	bankAccountTransactionRepository repositories.BankAccountTransactionRepository
}

func NewTransactionService(elasticsearchRepo repositories.BankAccountTransactionRepository) *TransactionService {
	return &TransactionService{
		bankAccountTransactionRepository: elasticsearchRepo,
	}
}

func (s *TransactionService) GetTransactionsByBankAndPeriod(bankID uint, year *int) ([]dto.BankAccountTransactionDTO, error) {
	// Fetch transactions from Elasticsearch
	return s.bankAccountTransactionRepository.GetTransactionsByBankAndPeriod(bankID, year)
}

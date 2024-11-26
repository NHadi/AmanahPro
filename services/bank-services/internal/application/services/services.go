package services

type Services struct {
	UploadService         *UploadService
	TransactionService    *TransactionService
	ReconciliationService *ReconciliationService
	ConsumerService       *ConsumerService
}

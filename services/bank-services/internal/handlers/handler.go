package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	Upload         *UploadHandler
	Transaction    *TransactionHandler
	Reconciliation *ReconciliationHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	uploadHandler *UploadHandler,
	transactionHandler *TransactionHandler,
	reconciliationHandler *ReconciliationHandler,
) *Handlers {
	return &Handlers{
		Upload:         uploadHandler,
		Transaction:    transactionHandler,
		Reconciliation: reconciliationHandler,
	}
}

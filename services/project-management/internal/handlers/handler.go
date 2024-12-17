package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	Project          *ProjectHandler
	ProjectFinancial *ProjectFinancialHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	projectHandler *ProjectHandler,
	projectFinancialHandler *ProjectFinancialHandler,
) *Handlers {
	return &Handlers{
		Project:          projectHandler,
		ProjectFinancial: projectFinancialHandler,
	}
}

package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	Project *ProjectHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	projectHandler *ProjectHandler,
) *Handlers {
	return &Handlers{
		Project: projectHandler,
	}
}

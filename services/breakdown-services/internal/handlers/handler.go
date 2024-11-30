package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	Breakdown *BreakdownHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	breakdownHandler *BreakdownHandler,
) *Handlers {
	return &Handlers{
		Breakdown: breakdownHandler,
	}
}

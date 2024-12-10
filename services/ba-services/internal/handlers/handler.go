package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	BA *BAHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	baHandler *BAHandler,
) *Handlers {
	return &Handlers{
		BA: baHandler,
	}
}

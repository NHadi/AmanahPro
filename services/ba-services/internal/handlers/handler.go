package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	Spk *SpkHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	sphHandler *SpkHandler,
) *Handlers {
	return &Handlers{
		Spk: sphHandler,
	}
}

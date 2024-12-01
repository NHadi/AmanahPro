package handlers

// Handlers aggregates all individual handlers
type Handlers struct {
	Sph *SphHandler
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(
	sphHandler *SphHandler,
) *Handlers {
	return &Handlers{
		Sph: sphHandler,
	}
}

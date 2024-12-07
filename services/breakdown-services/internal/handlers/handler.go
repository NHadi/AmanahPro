package handlers

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
	"time"
)

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

// Helper functions for merging non-zero or non-empty fields
func firstNonZeroInt(newVal, oldVal int) int {
	if newVal != 0 {
		return newVal
	}
	return oldVal
}

func firstNonEmptyString(newVal, oldVal string) string {
	if newVal != "" {
		return newVal
	}
	return oldVal
}

func firstNonEmptyStringPointer(newVal, oldVal *string) *string {
	if newVal != nil && *newVal != "" {
		return newVal
	}
	return oldVal
}

func firstNonZeroTime(newVal, oldVal time.Time) time.Time {
	if !newVal.IsZero() {
		return newVal
	}
	return oldVal
}

func firstNonZeroCustomDate(newVal, oldVal *models.CustomDate) *models.CustomDate {
	if newVal != nil && !newVal.IsZero() {
		return newVal
	}
	return oldVal
}

package dto

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

// BreakdownDTO represents the structure for Breakdown API requests/responses.
type BreakdownDTO struct {
	BreakdownId int                `json:"BreakdownId,omitempty"` // Optional for Create
	ProjectId   int                `json:"ProjectId"`
	ProjectName string             `json:"ProjectName"`
	Subject     string             `json:"Subject"`
	Location    *string            `json:"Location,omitempty"`
	Date        *models.CustomDate `json:"Date,omitempty"`
}

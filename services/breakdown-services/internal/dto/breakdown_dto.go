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

// BreakdownSectionDTO represents a breakdown section.
type BreakdownSectionDTO struct {
	BreakdownSectionId int                `json:"section_id,omitempty"` // Optional for Create
	BreakdownId        int                `json:"breakdown_id"`
	SectionTitle       string             `json:"section_title"`
	Items              []BreakdownItemDTO `json:"items,omitempty"` // Nested items
}

// BreakdownItemDTO represents a breakdown item.
type BreakdownItemDTO struct {
	BreakdownItemId int     `json:"item_id,omitempty"` // Optional for Create
	SectionId       int     `json:"section_id"`
	Description     string  `json:"description"`
	UnitPrice       float64 `json:"unit_price"`
}

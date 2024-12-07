package dto

import "AmanahPro/services/breakdown-services/internal/domain/models"

// BreakdownItemDTO represents a breakdown item.
type BreakdownItemDTO struct {
	BreakdownItemId int     `json:"BreakdownItemId,omitempty"` // Optional for Create
	Description     string  `json:"Description"`
	UnitPrice       float64 `json:"UnitPrice"`
}

// ToModel maps the DTO to the BreakdownSection model
func (dto *BreakdownItemDTO) ToModel(SectionId int, createdBy int) *models.BreakdownItem {
	return &models.BreakdownItem{
		Description: dto.Description,
		UnitPrice:   dto.UnitPrice,
		SectionId:   SectionId,
		CreatedBy:   &createdBy,
	}
}

// ToModelForUpdate maps the DTO to the BreakdownSection model for updates
func (dto *BreakdownItemDTO) ToModelForUpdate(existing *models.BreakdownItem, updatedBy int) *models.BreakdownItem {
	if dto.Description != "" {
		existing.Description = dto.Description
	}

	if dto.UnitPrice != 0 {
		existing.UnitPrice = dto.UnitPrice
	}

	existing.UpdatedBy = &updatedBy
	return existing
}

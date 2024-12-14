package dto

import (
	"AmanahPro/services/breakdown-services/internal/domain/models"
)

// BreakdownItemDTO represents a breakdown item.
type MstBreakdownItemDTO struct {
	BreakdownItemId int     `json:"BreakdownItemId,omitempty"` // Optional for Create
	Description     string  `json:"Description"`
	UnitPrice       float64 `json:"UnitPrice"`
	Sort            *int    `json:"Sort"`
}

// ToModel maps the DTO to the BreakdownSection model
func (dto *MstBreakdownItemDTO) ToModel(MstBreakdownSectionId, createdBy int) *models.MstBreakdownItem {
	return &models.MstBreakdownItem{
		Description:           dto.Description,
		Sort:                  *dto.Sort,
		UnitPrice:             dto.UnitPrice,
		MstBreakdownSectionId: MstBreakdownSectionId,
		CreatedBy:             &createdBy,
	}
}

// ToModelForUpdate maps the DTO to the BreakdownSection model for updates
func (dto *MstBreakdownItemDTO) ToModelForUpdate(existing *models.MstBreakdownItem, updatedBy int) *models.MstBreakdownItem {
	if dto.Description != "" {
		existing.Description = dto.Description
	}

	if dto.UnitPrice != 0 {
		existing.UnitPrice = dto.UnitPrice
	}

	if dto.Sort != nil {
		existing.Sort = *dto.Sort
	}

	existing.UpdatedBy = &updatedBy
	return existing
}

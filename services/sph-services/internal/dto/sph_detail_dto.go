package dto

import "AmanahPro/services/sph-services/internal/domain/models"

type SphDetailDTO struct {
	// SphDetailId     int      `json:"SphDetailId,omitempty"`
	ItemDescription *string  `json:"ItemDescription"`
	Quantity        *float64 `json:"Quantity"`
	Unit            *string  `json:"Unit"`
	UnitPrice       *float64 `json:"UnitPrice"`
	DiscountPrice   *float64 `json:"DiscountPrice,omitempty"`
	TotalPrice      *float64 `json:"TotalPrice,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *SphDetailDTO) ToModel(sectionID int, userID int) *models.SphDetail {
	return &models.SphDetail{
		// SphDetailId:     dto.SphDetailId,
		SectionId:       sectionID,
		ItemDescription: dto.ItemDescription,
		Quantity:        dto.Quantity,
		Unit:            dto.Unit,
		UnitPrice:       dto.UnitPrice,
		DiscountPrice:   dto.DiscountPrice,
		TotalPrice:      dto.TotalPrice,
		CreatedBy:       &userID,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SphDetailDTO) ToModelForUpdate(existing *models.SphDetail, userID int) *models.SphDetail {
	// Update fields only if the DTO contains non-nil or non-zero values
	if dto.ItemDescription != nil {
		existing.ItemDescription = dto.ItemDescription
	}
	if dto.Quantity != nil {
		existing.Quantity = dto.Quantity
	}
	if dto.Unit != nil {
		existing.Unit = dto.Unit
	}
	if dto.UnitPrice != nil {
		existing.UnitPrice = dto.UnitPrice
	}
	if dto.DiscountPrice != nil {
		existing.DiscountPrice = dto.DiscountPrice
	}
	if dto.TotalPrice != nil {
		existing.TotalPrice = dto.TotalPrice
	}

	// Always update the UpdatedBy field
	existing.UpdatedBy = &userID

	return existing
}

package dto

import "AmanahPro/services/sph-services/internal/domain/models"

type SphDetailDTO struct {
	SphDetailId     int      `json:"SphDetailId,omitempty"`
	ItemDescription *string  `json:"ItemDescription" binding:"required"`
	Quantity        *float64 `json:"Quantity" binding:"required"`
	Unit            *string  `json:"Unit" binding:"required"`
	UnitPrice       *float64 `json:"UnitPrice" binding:"required"`
	DiscountPrice   *float64 `json:"DiscountPrice,omitempty"`
	TotalPrice      *float64 `json:"TotalPrice,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *SphDetailDTO) ToModel(sectionID int, userID int) *models.SphDetail {
	return &models.SphDetail{
		SphDetailId:     dto.SphDetailId,
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
	existing.ItemDescription = dto.ItemDescription
	existing.Quantity = dto.Quantity
	existing.Unit = dto.Unit
	existing.UnitPrice = dto.UnitPrice
	existing.DiscountPrice = dto.DiscountPrice
	existing.TotalPrice = dto.TotalPrice
	existing.UpdatedBy = &userID
	return existing
}

package dto

import "AmanahPro/services/ba-services/internal/domain/models"

type BADetailDTO struct {
	DetailId      int      `json:"DetailId,omitempty"`
	SectionID     *int     `json:"SectionID"`
	ItemName      *string  `json:"ItemName"`
	Quantity      *float64 `json:"Quantity"`
	Unit          *string  `json:"Unit"`
	UnitPrice     *float64 `json:"UnitPrice"`
	DiscountPrice *float64 `json:"DiscountPrice,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *BADetailDTO) ToModel(userID int) *models.BADetail {
	return &models.BADetail{
		SectionID:     dto.SectionID,
		ItemName:      dto.ItemName,
		Quantity:      *dto.Quantity,
		Unit:          dto.Unit,
		UnitPrice:     dto.UnitPrice,
		DiscountPrice: dto.DiscountPrice,
		CreatedBy:     &userID,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *BADetailDTO) ToModelForUpdate(existing *models.BADetail, userID int) *models.BADetail {
	if dto.ItemName != nil {
		existing.ItemName = dto.ItemName
	}
	if dto.Quantity != nil {
		existing.Quantity = *dto.Quantity
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
	existing.UpdatedBy = &userID

	return existing
}

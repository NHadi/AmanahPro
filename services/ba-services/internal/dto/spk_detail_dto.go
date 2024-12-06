package dto

import "AmanahPro/services/ba-services/internal/domain/models"

type SpkDetailDTO struct {
	DetailId          int      `json:"DetailId,omitempty"`
	Description       *string  `json:"Description" binding:"required"`
	Quantity          *float64 `json:"Quantity" binding:"required"`
	Unit              *string  `json:"Unit" binding:"required"`
	UnitPriceJasa     *float64 `json:"UnitPriceJasa" binding:"required"`
	TotalJasa         *float64 `json:"TotalJasa" binding:"required"`
	UnitPriceMaterial float64  `json:"UnitPriceMaterial" binding:"required"`
	TotalMaterial     float64  `json:"TotalMaterial" binding:"required"`
}

// ToModel maps the DTO to the domain model
func (dto *SpkDetailDTO) ToModel(userID int) *models.SPKDetail {
	return &models.SPKDetail{
		DetailId:          dto.DetailId,
		Description:       dto.Description,
		Quantity:          *dto.Quantity,
		Unit:              dto.Unit,
		UnitPriceJasa:     *dto.UnitPriceJasa,
		TotalJasa:         *dto.TotalJasa,
		UnitPriceMaterial: dto.UnitPriceMaterial,
		TotalMaterial:     dto.TotalMaterial,
		CreatedBy:         &userID,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SpkDetailDTO) ToModelForUpdate(existing *models.SPKDetail, userID int) *models.SPKDetail {
	existing.Description = dto.Description
	existing.Quantity = *dto.Quantity
	existing.Unit = dto.Unit
	existing.UnitPriceJasa = *dto.UnitPriceJasa
	existing.TotalJasa = *dto.TotalJasa
	existing.UpdatedBy = &userID
	existing.UnitPriceMaterial = dto.UnitPriceMaterial
	existing.TotalMaterial = dto.TotalMaterial

	return existing
}

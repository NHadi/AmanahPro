package dto

import "AmanahPro/services/spk-services/internal/domain/models"

type SpkDetailDTO struct {
	DetailId          int      `json:"DetailId,omitempty"`
	Description       *string  `json:"Description"`
	Quantity          *float64 `json:"Quantity"`
	Unit              *string  `json:"Unit"`
	UnitPriceJasa     *float64 `json:"UnitPriceJasa"`
	TotalJasa         *float64 `json:"TotalJasa"`
	UnitPriceMaterial float64  `json:"UnitPriceMaterial"`
	TotalMaterial     float64  `json:"TotalMaterial"`
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
	if dto.Description != nil {
		existing.Description = dto.Description
	}
	if dto.Quantity != nil {
		existing.Quantity = *dto.Quantity
	}
	if dto.Unit != nil {
		existing.Unit = dto.Unit
	}
	if dto.UnitPriceJasa != nil {
		existing.UnitPriceJasa = *dto.UnitPriceJasa
	}
	if dto.TotalJasa != nil {
		existing.TotalJasa = *dto.TotalJasa
	}
	if dto.UnitPriceMaterial != 0 {
		existing.UnitPriceMaterial = dto.UnitPriceMaterial
	}
	if dto.TotalMaterial != 0 {
		existing.TotalMaterial = dto.TotalMaterial
	}

	existing.UpdatedBy = &userID

	return existing
}

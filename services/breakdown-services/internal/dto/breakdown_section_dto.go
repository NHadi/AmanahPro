package dto

import "AmanahPro/services/breakdown-services/internal/domain/models"

type BreakdownSectionDTO struct {
	SectionTitle string `json:"SectionTitle"`
}

// ToModel maps the DTO to the BreakdownSection model
func (dto *BreakdownSectionDTO) ToModel(breakdownID int, createdBy int) *models.BreakdownSection {
	return &models.BreakdownSection{
		BreakdownId:  breakdownID,
		SectionTitle: dto.SectionTitle,
		CreatedBy:    &createdBy,
	}
}

// ToModelForUpdate maps the DTO to the BreakdownSection model for updates
func (dto *BreakdownSectionDTO) ToModelForUpdate(existing *models.BreakdownSection, updatedBy int) *models.BreakdownSection {
	existing.SectionTitle = dto.SectionTitle
	existing.UpdatedBy = &updatedBy
	return existing
}

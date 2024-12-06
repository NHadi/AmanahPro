package dto

import "AmanahPro/services/ba-services/internal/domain/models"

type SpkSectionDTO struct {
	SectionId    int     `json:"SectionId,omitempty"`
	SectionTitle *string `json:"SectionTitle" binding:"required"`
}

// ToModel maps the DTO to the domain model
func (dto *SpkSectionDTO) ToModel(userID int) *models.SPKSection {
	section := &models.SPKSection{
		SectionId:    dto.SectionId,
		SectionTitle: dto.SectionTitle,
		CreatedBy:    &userID,
	}

	return section
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SpkSectionDTO) ToModelForUpdate(existing *models.SPKSection, userID int) *models.SPKSection {
	existing.SectionTitle = dto.SectionTitle
	existing.UpdatedBy = &userID

	return existing
}

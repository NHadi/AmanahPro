package dto

import "AmanahPro/services/sph-services/internal/domain/models"

type SphSectionDTO struct {
	SphSectionId int     `json:"SphSectionId,omitempty"`
	SphId        int     `json:"SphId" binding:"required"`
	SectionTitle *string `json:"SectionTitle" binding:"required"`
}

// ToModel maps the DTO to the domain model
func (dto *SphSectionDTO) ToModel(sphID int, userID int) *models.SphSection {
	return &models.SphSection{
		SphSectionId: dto.SphSectionId,
		SphId:        sphID,
		SectionTitle: dto.SectionTitle,
		CreatedBy:    &userID,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SphSectionDTO) ToModelForUpdate(existing *models.SphSection, userID int) *models.SphSection {
	existing.SectionTitle = dto.SectionTitle
	existing.UpdatedBy = &userID
	return existing
}

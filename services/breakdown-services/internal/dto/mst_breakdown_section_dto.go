package dto

import "AmanahPro/services/breakdown-services/internal/domain/models"

// MstBreakdownSectionDTO represents a master breakdown section.
type MstBreakdownSectionDTO struct {
	MstBreakdownSectionId int    `json:"MstBreakdownSectionId,omitempty"` // Optional for Create
	SectionTitle          string `json:"SectionTitle"`
	Sort                  *int   `json:"Sort"`
}

// ToModel maps the DTO to the MstBreakdownSection model for creation
func (dto *MstBreakdownSectionDTO) ToModel(createdBy int) *models.MstBreakdownSection {
	return &models.MstBreakdownSection{
		SectionTitle: dto.SectionTitle,
		Sort:         *dto.Sort,
		CreatedBy:    &createdBy,
	}
}

// ToModelForUpdate maps the DTO to the MstBreakdownSection model for updates
func (dto *MstBreakdownSectionDTO) ToModelForUpdate(existing *models.MstBreakdownSection, updatedBy int) *models.MstBreakdownSection {
	if dto.SectionTitle != "" {
		existing.SectionTitle = dto.SectionTitle
	}

	if dto.Sort != nil {
		existing.Sort = *dto.Sort
	}

	existing.UpdatedBy = &updatedBy
	return existing
}

package dto

import "AmanahPro/services/ba-services/internal/domain/models"

type BASectionDTO struct {
	BASectionId int     `json:"BASectionId,omitempty"`
	BAID        int     `json:"BAID"`
	SectionName *string `json:"SectionName"`
}

// ToModel maps the DTO to the domain model
func (dto *BASectionDTO) ToModel(userID int) *models.BASection {
	return &models.BASection{
		BAID:        &dto.BAID,
		SectionName: dto.SectionName,
		CreatedBy:   &userID,
	}
}

func (dto *BASectionDTO) ToModelForUpdate(existing *models.BASection, userID int) *models.BASection {
	if dto.SectionName != nil {
		existing.SectionName = dto.SectionName
	}

	existing.UpdatedBy = &userID
	return existing
}

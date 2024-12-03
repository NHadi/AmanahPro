package dto

import "AmanahPro/services/spk-services/internal/domain/models"

type SpkDTO struct {
	SpkId int `json:"SpkId,omitempty"`
	SphId int `json:"SphId,omitempty"`

	ProjectId   *int               `json:"ProjectId" binding:"required"`
	ProjectName *string            `json:"ProjectName" binding:"required"`
	Subject     *string            `json:"Subject" binding:"required"`
	Date        *models.CustomDate `json:"Date,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *SpkDTO) ToModel(userID int) *models.SPK {
	spk := &models.SPK{
		SpkId:       dto.SpkId,
		ProjectId:   dto.ProjectId,
		ProjectName: dto.ProjectName,
		Subject:     dto.Subject,
		Date:        dto.Date,
		CreatedBy:   &userID,
	}

	return spk
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SpkDTO) ToModelForUpdate(existing *models.SPK, userID int) *models.SPK {
	existing.ProjectId = dto.ProjectId
	existing.ProjectName = dto.ProjectName
	existing.Subject = dto.Subject
	existing.Date = dto.Date
	existing.UpdatedBy = &userID

	return existing
}

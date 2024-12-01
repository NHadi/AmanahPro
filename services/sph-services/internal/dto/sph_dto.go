package dto

import (
	"AmanahPro/services/sph-services/internal/domain/models"
)

type SphDTO struct {
	SphId         int                `json:"SphId,omitempty"`
	ProjectId     *int               `json:"ProjectId" binding:"required"`
	ProjectName   *string            `json:"ProjectName" binding:"required"`
	Subject       *string            `json:"Subject" binding:"required"`
	Location      *string            `json:"Location,omitempty"`
	Date          *models.CustomDate `json:"Date,omitempty"`
	RecepientName *string            `json:"RecepientName,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *SphDTO) ToModel(organizationID *int, userID int) *models.Sph {
	return &models.Sph{
		SphId:          dto.SphId,
		ProjectId:      dto.ProjectId,
		ProjectName:    dto.ProjectName,
		Subject:        dto.Subject,
		Location:       dto.Location,
		Date:           dto.Date,
		RecepientName:  dto.RecepientName,
		OrganizationId: organizationID,
		CreatedBy:      &userID,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SphDTO) ToModelForUpdate(existing *models.Sph, userID int) *models.Sph {
	existing.ProjectId = dto.ProjectId
	existing.ProjectName = dto.ProjectName
	existing.Subject = dto.Subject
	existing.Location = dto.Location
	existing.Date = dto.Date
	existing.RecepientName = dto.RecepientName
	existing.UpdatedBy = &userID
	return existing
}

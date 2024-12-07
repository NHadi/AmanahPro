package dto

import (
	"AmanahPro/services/sph-services/internal/domain/models"
)

type SphDTO struct {
	SphId         int                `json:"SphId,omitempty"`
	ProjectId     *int               `json:"ProjectId" binding:"required"`
	ProjectName   *string            `json:"ProjectName"`
	Subject       *string            `json:"Subject,omitempty"`
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
	// Update fields only if the DTO contains non-empty values
	if dto.ProjectId != nil {
		existing.ProjectId = dto.ProjectId
	}
	if dto.ProjectName != nil {
		existing.ProjectName = dto.ProjectName
	}
	if dto.Subject != nil {
		existing.Subject = dto.Subject
	}
	if dto.Location != nil {
		existing.Location = dto.Location
	}
	if dto.Date != nil { // Check for zero value in time.Time
		existing.Date = dto.Date
	}
	if dto.RecepientName != nil {
		existing.RecepientName = dto.RecepientName
	}

	// Always update the UpdatedBy field
	existing.UpdatedBy = &userID

	return existing
}

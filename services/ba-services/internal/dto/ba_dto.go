package dto

import "AmanahPro/services/ba-services/internal/domain/models"

type BADTO struct {
	ProjectId      *int               `json:"ProjectId"`
	ProjectName    *string            `json:"ProjectName"`
	BASubject      *string            `json:"BASubject"`
	BADate         *models.CustomDate `json:"BADate"`
	OrganizationId *int               `json:"OrganizationId,omitempty"`
	SphId          int                `json:"SphId,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *BADTO) ToModel(userID int) *models.BA {
	return &models.BA{
		ProjectId:   dto.ProjectId,
		ProjectName: dto.ProjectName,
		BASubject:   *dto.BASubject,
		BADate:      *dto.BADate,
		CreatedBy:   &userID,
		SphId:       &dto.SphId,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *BADTO) ToModelForUpdate(existing *models.BA, userID int) *models.BA {
	if dto.ProjectId != nil {
		existing.ProjectId = dto.ProjectId
	}
	if dto.ProjectName != nil {
		existing.ProjectName = dto.ProjectName
	}
	if dto.BASubject != nil {
		existing.BASubject = *dto.BASubject
	}
	if dto.BADate != nil {
		existing.BADate = *dto.BADate
	}

	existing.UpdatedBy = &userID
	return existing
}

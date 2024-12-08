package dto

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

// ProjectDTO represents the Data Transfer Object for the Project entity
type ProjectDTO struct {
	ProjectID   int                `json:"ProjectID,omitempty"` // Project ID
	ProjectName string             `json:"ProjectName"`         // Project Name (Required)
	Location    *string            `json:"Location,omitempty"`  // Project Location
	StartDate   *models.CustomDate `json:"StartDate,omitempty"`
	EndDate     *models.CustomDate `json:"EndDate,omitempty"`
	Description *string            `json:"Description,omitempty"` // Project Description
	Status      *string            `json:"Status,omitempty"`      // Project Status
}

// ToModel maps the DTO to the domain model for creation
func (dto *ProjectDTO) ToModel(userID int, organizationID int) *models.Project {
	project := &models.Project{
		ProjectID:      dto.ProjectID,
		ProjectName:    dto.ProjectName,
		Location:       dto.Location,
		StartDate:      dto.StartDate,
		EndDate:        dto.EndDate,
		Description:    dto.Description,
		Status:         dto.Status,
		CreatedBy:      &userID,
		OrganizationID: &organizationID,
	}

	return project
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *ProjectDTO) ToModelForUpdate(existing *models.Project, userID int) *models.Project {
	if dto.ProjectName != "" {
		existing.ProjectName = dto.ProjectName
	}
	if dto.Location != nil {
		existing.Location = dto.Location
	}
	if dto.StartDate != nil {
		existing.StartDate = dto.StartDate
	}
	if dto.EndDate != nil {
		existing.EndDate = dto.EndDate
	}
	if dto.Description != nil {
		existing.Description = dto.Description
	}
	if dto.Status != nil {
		existing.Status = dto.Status
	}

	existing.UpdatedBy = &userID

	return existing
}

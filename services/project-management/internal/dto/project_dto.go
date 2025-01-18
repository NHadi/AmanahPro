package dto

import (
	"AmanahPro/services/project-management/internal/domain/models"
)

// ProjectDTO represents the Data Transfer Object for the Project entity
type ProjectDTO struct {
	ProjectID                 int                `json:"ProjectID,omitempty"` // Project ID
	ProjectName               string             `json:"ProjectName"`         // Project Name (Required)
	Location                  *string            `json:"Location,omitempty"`  // Project Location
	StartDate                 *models.CustomDate `json:"StartDate,omitempty"`
	EndDate                   *models.CustomDate `json:"EndDate,omitempty"`
	Description               *string            `json:"Description,omitempty"`             // Project Description
	Status                    *string            `json:"Status,omitempty"`                  // Project Status
	PO_SPHDate                *models.CustomDate `gorm:"column:PO_SPHDate;type:date;null"`  // PO/SPH date
	SPH                       *float64           `gorm:"column:SPH;type:decimal(18,2;null"` // SPH
	Termin                    *float64           `gorm:"column:Termin;decimal(18,2;null"`   // Termin
	TotalSPK                  *float64           `gorm:"column:TotalSPK;type:decimal(18,2);null"`
	TotalBreakdown            *float64           `gorm:"column:TotalBreakdown;type:decimal(18,2);null"`
	ProgressCurrentM2         *float64           `gorm:"column:ProgressCurrentM2;type:decimal(10,2);null"`        // Current progress in M2
	ProgressCurrentPercentage *float64           `gorm:"column:ProgressCurrentPercentage;type:decimal(5,2);null"` // Current progress percentage
}

// ToModel maps the DTO to the domain model for creation
func (dto *ProjectDTO) ToModel(userID int, organizationID int) *models.Project {
	project := &models.Project{
		ProjectID:                 dto.ProjectID,
		ProjectName:               dto.ProjectName,
		Location:                  dto.Location,
		StartDate:                 dto.StartDate,
		EndDate:                   dto.EndDate,
		Description:               dto.Description,
		Status:                    dto.Status,
		CreatedBy:                 &userID,
		OrganizationID:            &organizationID,
		PO_SPHDate:                dto.PO_SPHDate,
		SPH:                       dto.SPH,
		Termin:                    dto.Termin,
		TotalSPK:                  dto.TotalSPK,
		TotalBreakdown:            dto.TotalBreakdown,
		ProgressCurrentM2:         dto.ProgressCurrentM2,
		ProgressCurrentPercentage: dto.ProgressCurrentPercentage,
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

	if dto.Termin != nil {
		existing.Termin = dto.Termin
	}

	if dto.PO_SPHDate != nil {
		existing.PO_SPHDate = dto.PO_SPHDate
	}

	if dto.TotalSPK != nil {
		if existing.TotalSPK == nil {
			existing.TotalSPK = new(float64)
		}
		*existing.TotalSPK += *dto.TotalSPK
	}
	if dto.TotalBreakdown != nil {
		if existing.TotalBreakdown == nil {
			existing.TotalBreakdown = new(float64)
		}
		*existing.TotalBreakdown += *dto.TotalBreakdown

	}
	if dto.SPH != nil {
		existing.SPH = dto.SPH
	}
	if dto.ProgressCurrentM2 != nil {
		existing.ProgressCurrentM2 = dto.ProgressCurrentM2
	}
	if dto.ProgressCurrentPercentage != nil {
		existing.ProgressCurrentPercentage = dto.ProgressCurrentPercentage
	}

	existing.UpdatedBy = &userID

	return existing
}

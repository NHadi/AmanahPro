package dto

import "AmanahPro/services/spk-services/internal/domain/models"

type SpkDTO struct {
	SpkId int `json:"SpkId,omitempty"`
	SphId int `json:"SphId,omitempty"`

	ProjectId     *int     `json:"ProjectId"`
	ProjectName   *string  `json:"ProjectName"`
	Mandor        *string  `json:"Mandor"`
	TotalJasa     *float64 `json:"TotalJasa"`
	TotalMaterial *float64 `json:"TotalMaterial"`

	Subject *string            `json:"Subject"`
	Date    *models.CustomDate `json:"Date,omitempty"`
}

// ToModel maps the DTO to the domain model
func (dto *SpkDTO) ToModel(userID int) *models.SPK {
	spk := &models.SPK{
		SpkId:         dto.SpkId,
		ProjectId:     dto.ProjectId,
		ProjectName:   dto.ProjectName,
		Subject:       dto.Subject,
		Date:          dto.Date,
		CreatedBy:     &userID,
		Mandor:        dto.Mandor,
		TotalJasa:     dto.TotalJasa,
		TotalMaterial: dto.TotalMaterial,
	}

	return spk
}

// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *SpkDTO) ToModelForUpdate(existing *models.SPK, userID int) *models.SPK {
	// Use values from dto if not nil; otherwise, keep the existing values
	if dto.ProjectId != nil {
		existing.ProjectId = dto.ProjectId
	}
	if dto.ProjectName != nil {
		existing.ProjectName = dto.ProjectName
	}
	if dto.Subject != nil {
		existing.Subject = dto.Subject
	}
	if dto.Date != nil {
		existing.Date = dto.Date
	}
	if dto.TotalJasa != nil {
		existing.TotalJasa = dto.TotalJasa
	}
	if dto.TotalMaterial != nil {
		existing.TotalMaterial = dto.TotalMaterial
	}

	existing.UpdatedBy = &userID

	return existing
}

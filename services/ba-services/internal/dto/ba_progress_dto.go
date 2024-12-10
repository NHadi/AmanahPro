package dto

import "AmanahPro/services/ba-services/internal/domain/models"

// BAProgressDTO represents the data transfer object for BA Progress
type BAProgressDTO struct {
	BAProgressId               int      `json:"BAProgressId,omitempty"`
	DetailId                   int      `json:"DetailId,omitempty"`
	ProgressPreviousM2         *float64 `json:"ProgressPreviousM2,omitempty"`
	ProgressPreviousPercentage *float64 `json:"ProgressPreviousPercentage,omitempty"`
	ProgressCurrentM2          *float64 `json:"ProgressCurrentM2,omitempty"`
	ProgressCurrentPercentage  *float64 `json:"ProgressCurrentPercentage,omitempty"`
	CreatedBy                  *int     `json:"CreatedBy,omitempty"`
	UpdatedBy                  *int     `json:"UpdatedBy,omitempty"`
	OrganizationId             *int     `json:"OrganizationId,omitempty"`
}

// ToModel maps the DTO to the domain model for creating a new record
func (dto *BAProgressDTO) ToModel(userID int) *models.BAProgress {
	return &models.BAProgress{
		DetailId:                   dto.DetailId,
		ProgressPreviousM2:         dto.ProgressPreviousM2,
		ProgressPreviousPercentage: dto.ProgressPreviousPercentage,
		ProgressCurrentM2:          dto.ProgressCurrentM2,
		ProgressCurrentPercentage:  dto.ProgressCurrentPercentage,
		CreatedBy:                  &userID,
		OrganizationId:             dto.OrganizationId,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updating an existing record
func (dto *BAProgressDTO) ToModelForUpdate(existing *models.BAProgress, userID int) *models.BAProgress {
	// Update fields if DTO values are not nil
	if dto.ProgressPreviousM2 != nil {
		existing.ProgressPreviousM2 = dto.ProgressPreviousM2
	}
	if dto.ProgressPreviousPercentage != nil {
		existing.ProgressPreviousPercentage = dto.ProgressPreviousPercentage
	}
	if dto.ProgressCurrentM2 != nil {
		existing.ProgressCurrentM2 = dto.ProgressCurrentM2
	}
	if dto.ProgressCurrentPercentage != nil {
		existing.ProgressCurrentPercentage = dto.ProgressCurrentPercentage
	}

	// Update metadata
	existing.UpdatedBy = &userID

	return existing
}

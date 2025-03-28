package dto

import (
	"AmanahPro/services/project-management/internal/domain/models"
	"time"
)

// ProjectFinancialDTO represents the Data Transfer Object for ProjectFinancial entity
type ProjectFinancialDTO struct {
	ProjectID       int                `json:"ProjectID"`               // Project ID (Required)
	ProjectUserID   *int               `json:"ProjectUserID,omitempty"` // User ID (Optional)
	TransactionDate *models.CustomDate `json:"TransactionDate"`         // Transaction Date (Required)
	Description     *string            `json:"Description"`             // Financial Record Description (Required)
	Amount          float64            `json:"Amount"`                  // Amount (Required)
	AmountDeviden   *float64           `json:"AmountDeviden,omitempty"` // Amount (Required)
	TransactionType string             `json:"TransactionType"`         // Transaction Type (In/Out)
	Category        *string            `json:"Category"`                // Category (BB, Operational, General)
	CreatedAt       *time.Time         `json:"CreatedAt,omitempty"`     // Created Date
	UpdatedAt       *time.Time         `json:"UpdatedAt,omitempty"`     // Updated Date
}

// ToModel maps the DTO to the domain model for creation
func (dto *ProjectFinancialDTO) ToModel(userID int, organizationID int) *models.ProjectFinancial {
	return &models.ProjectFinancial{
		ProjectID:       dto.ProjectID,
		ProjectUserID:   dto.ProjectUserID,
		TransactionDate: *dto.TransactionDate,
		Description:     *dto.Description,
		Amount:          dto.Amount,
		AmountDeviden:   dto.AmountDeviden,
		TransactionType: dto.TransactionType,
		Category: func() string {
			if dto.Category != nil {
				return *dto.Category
			}
			return "" // Default empty string if nil
		}(),
		CreatedBy:      &userID,
		OrganizationID: &organizationID,
	}
}

// ToModelForUpdate maps the DTO to the domain model for updates
// ToModelForUpdate maps the DTO to the domain model for updates
func (dto *ProjectFinancialDTO) ToModelForUpdate(existing *models.ProjectFinancial, userID int) *models.ProjectFinancial {
	if dto.ProjectID != 0 {
		existing.ProjectID = dto.ProjectID
	}
	if dto.ProjectUserID != nil {
		existing.ProjectUserID = dto.ProjectUserID
	}
	if dto.TransactionDate != nil {
		existing.TransactionDate = *dto.TransactionDate
	}
	if dto.Description != nil {
		existing.Description = *dto.Description
	}
	if dto.Amount != 0 {
		existing.Amount = dto.Amount
	}
	if dto.TransactionType != "" {
		existing.TransactionType = dto.TransactionType
	}
	if dto.Category != nil {
		existing.Category = *dto.Category
	}
	if dto.AmountDeviden != nil {
		existing.AmountDeviden = dto.AmountDeviden
	}

	existing.UpdatedBy = &userID

	return existing
}

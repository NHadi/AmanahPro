package models

import (
	"time"
)

type ProjectRecap struct {
	ID               int        `gorm:"primaryKey;autoIncrement"` // Primary key
	ProjectID        int        `gorm:"not null"`                 // Foreign key to Projects
	TotalOpname      float64    `gorm:"type:decimal(18,2);null"`  // Total opname
	TotalPengeluaran float64    `gorm:"type:decimal(18,2);null"`  // Total expenditure
	Margin           float64    `gorm:"type:decimal(18,2);null"`  // Margin value
	MarginPercentage float64    `gorm:"type:decimal(5,2);null"`   // Margin percentage
	CreatedBy        *int       `gorm:"null"`                     // Created by user ID
	CreatedAt        *time.Time `gorm:"autoCreateTime"`           // Creation timestamp
	UpdatedBy        *int       `gorm:"null"`                     // Updated by user ID
	UpdatedAt        *time.Time `gorm:"autoUpdateTime"`           // Update timestamp
	DeletedBy        *int       `gorm:"null"`                     // Deleted by user ID
	DeletedAt        *time.Time `gorm:"index;null"`               // Deletion timestamp
	OrganizationID   int        `gorm:"null"`                     // Organization ID
}

func (ProjectRecap) TableName() string {
	return "ProjectRekap"
}

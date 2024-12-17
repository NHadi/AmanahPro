package models

import (
	"time"
)

type ProjectUser struct {
	ID             int        `gorm:"column:ID;primaryKey;autoIncrement"`           // Primary key
	ProjectID      int        `gorm:"column:ProjectID;not null"`                    // Foreign key to Projects
	UserID         *int       `gorm:"column:UserID"`                                // User ID (nullable)
	UserName       string     `gorm:"column:UserrName;type:nvarchar(250);not null"` // User name
	Role           string     `gorm:"column:Role;type:nvarchar(50);not null"`       // Role
	CreatedAt      *time.Time `gorm:"column:CreatedAt;default:current_timestamp"`   // Created timestamp
	CreatedBy      *int       `gorm:"column:CreatedBy"`                             // Created by user ID
	UpdatedAt      *time.Time `gorm:"column:UpdatedAt"`                             // Updated timestamp
	UpdatedBy      *int       `gorm:"column:UpdatedBy"`                             // Updated by user ID
	DeletedAt      *time.Time `gorm:"column:DeletedAt;index"`                       // Soft delete timestamp
	DeletedBy      *int       `gorm:"column:DeletedBy"`                             // Deleted by user ID
	OrganizationID *int       `gorm:"column:OrganizationID"`                        // Organization ID

	// Relationships
	Project *Project `gorm:"foreignKey:ProjectID;references:ProjectID"` // Relationship to Projects table
}

// TableName overrides the table name in GORM
func (ProjectUser) TableName() string {
	return "ProjectUser"
}

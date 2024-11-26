package models

import (
	"time"
)

// Project represents the Projects table in the database
type Project struct {
	ProjectID      int        `gorm:"primaryKey;autoIncrement"`                    // Primary key
	ProjectName    string     `gorm:"type:varchar(255);not null"`                  // Name of the project
	Location       *string    `gorm:"type:varchar(255);null"`                      // Project location
	StartDate      *time.Time `gorm:"type:date;null"`                              // Start date of the project
	EndDate        *time.Time `gorm:"type:date;null"`                              // End date of the project
	Description    *string    `gorm:"type:text;null"`                              // Description of the project
	Status         *string    `gorm:"type:varchar(20);default:'in-progress';null"` // Status of the project
	CreatedBy      *int       `gorm:"null"`                                        // Created by user ID
	CreatedAt      *time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`    // Creation timestamp
	UpdatedBy      *int       `gorm:"null"`                                        // Updated by user ID
	UpdatedAt      *time.Time `gorm:"autoUpdateTime;null"`                         // Update timestamp
	DeletedBy      *int       `gorm:"null"`                                        // Deleted by user ID
	DeletedAt      *time.Time `gorm:"index;null"`                                  // Deletion timestamp
	OrganizationID int        `gorm:"null"`                                        // Organization ID
}

// TableName specifies the table name for Project
func (Project) TableName() string {
	return "Projects"
}

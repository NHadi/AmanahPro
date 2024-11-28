package models

import "time"

// Breakdown represents the Breakdowns table in the database
type Breakdown struct {
	BreakdownId    int        `gorm:"primaryKey;autoIncrement"`   // Primary key
	ProjectId      int        `gorm:"not null"`                   // Foreign key to Projects
	Subject        string     `gorm:"type:varchar(255);not null"` // Breakdown subject
	Location       *string    `gorm:"type:varchar(255);null"`     // Optional location
	Date           *time.Time `gorm:"type:date;null"`             // Breakdown date
	CreatedBy      *int       `gorm:"null"`                       // Created by user ID
	CreatedAt      *time.Time `gorm:"autoCreateTime"`             // Creation timestamp
	UpdatedBy      *int       `gorm:"null"`                       // Updated by user ID
	UpdatedAt      *time.Time `gorm:"autoUpdateTime"`             // Update timestamp
	DeletedBy      *int       `gorm:"null"`                       // Deleted by user ID
	DeletedAt      *time.Time `gorm:"index;null"`                 // Deletion timestamp
	OrganizationId *int       `gorm:"null"`                       // Organization ID
}

// TableName specifies the table name for Breakdown
func (Breakdown) TableName() string {
	return "Breakdowns"
}

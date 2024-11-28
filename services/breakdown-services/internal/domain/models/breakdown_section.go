package models

import "time"

// BreakdownSection represents the BreakdownSections table in the database
type BreakdownSection struct {
	BreakdownSectionId int        `gorm:"primaryKey;autoIncrement"`   // Primary key
	BreakdownId        int        `gorm:"not null"`                   // Foreign key to Breakdowns
	SectionTitle       string     `gorm:"type:varchar(255);not null"` // Section title
	CreatedBy          *int       `gorm:"null"`                       // Created by user ID
	CreatedAt          *time.Time `gorm:"autoCreateTime"`             // Creation timestamp
	UpdatedBy          *int       `gorm:"null"`                       // Updated by user ID
	UpdatedAt          *time.Time `gorm:"autoUpdateTime"`             // Update timestamp
	DeletedBy          *int       `gorm:"null"`                       // Deleted by user ID
	DeletedAt          *time.Time `gorm:"index;null"`                 // Deletion timestamp
	OrganizationId     *int       `gorm:"null"`                       // Organization ID
}

// TableName specifies the table name for BreakdownSection
func (BreakdownSection) TableName() string {
	return "BreakdownSections"
}

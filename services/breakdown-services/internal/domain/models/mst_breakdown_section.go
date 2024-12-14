package models

import "time"

// BreakdownSection represents the BreakdownSections table in the database
type MstBreakdownSection struct {
	MstBreakdownSectionId int        `gorm:"primaryKey;column:MstBreakdownSectionId;autoIncrement"` // Primary key
	SectionTitle          string     `gorm:"column:SectionTitle;type:varchar(255);not null"`        // Section title
	Sort                  int        `gorm:"column:Sort;not null"`                                  // Sort
	CreatedBy             *int       `gorm:"column:CreatedBy;null"`                                 // Created by user ID
	CreatedAt             *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                       // Creation timestamp
	UpdatedBy             *int       `gorm:"column:UpdatedBy;null"`                                 // Updated by user ID
	UpdatedAt             *time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`                       // Update timestamp
	DeletedBy             *int       `gorm:"column:DeletedBy;null"`                                 // Deleted by user ID
	DeletedAt             *time.Time `gorm:"column:DeletedAt;index;null"`                           // Deletion timestamp
	OrganizationId        *int       `gorm:"column:OrganizationId;null"`                            // Organization ID

	// Relations
	Items []MstBreakdownItem `gorm:"foreignKey:MstBreakdownSectionId"` // One-to-Many relationship with BreakdownItem
}

// TableName specifies the table name for BreakdownSection
func (MstBreakdownSection) TableName() string {
	return "MstBreakdownSections"
}

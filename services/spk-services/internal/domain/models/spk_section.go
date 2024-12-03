package models

import "time"

// SPKSection represents the SPK Sections table in the database
type SPKSection struct {
	SectionId      int        `gorm:"primaryKey;column:SectionId;autoIncrement"`      // Primary key
	SpkId          int        `gorm:"column:SpkId;not null"`                          // Foreign key to SPK
	SectionTitle   *string    `gorm:"column:SectionTitle;type:varchar(255);not null"` // Section title
	CreatedBy      *int       `gorm:"column:CreatedBy;null"`                          // Created by user ID
	CreatedAt      *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                // Creation timestamp
	UpdatedBy      *int       `gorm:"column:UpdatedBy;null"`                          // Updated by user ID
	UpdatedAt      *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`           // Update timestamp
	DeletedBy      *int       `gorm:"column:DeletedBy;null"`                          // Deleted by user ID
	DeletedAt      *time.Time `gorm:"column:DeletedAt;index;null"`                    // Deletion timestamp
	OrganizationId *int       `gorm:"column:OrganizationId;null"`                     // Organization ID

	// Relations
	Details []SPKDetail `gorm:"foreignKey:SectionId"` // One-to-Many relationship with SPK Details
}

// TableName specifies the table name for SPKSection
func (SPKSection) TableName() string {
	return "SPKSections"
}

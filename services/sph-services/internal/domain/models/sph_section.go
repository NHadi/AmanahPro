package models

import (
	"encoding/json"
	"time"
)

// SphSection represents the SPHSections table in the database
type SphSection struct {
	SphSectionId   int        `gorm:"primaryKey;column:SphSectionId;autoIncrement"` // Primary key
	SphId          int        `gorm:"column:SphId;not null"`                        // Foreign key to SPH
	SectionTitle   *string    `gorm:"column:SectionTitle;type:varchar(255);null"`   // Section title
	CreatedBy      *int       `gorm:"column:CreatedBy;null"`                        // Created by user ID
	CreatedAt      time.Time  `gorm:"column:CreatedAt;autoCreateTime"`              // Creation timestamp
	UpdatedBy      *int       `gorm:"column:UpdatedBy;null"`                        // Updated by user ID
	UpdatedAt      *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`         // Update timestamp
	DeletedBy      *int       `gorm:"column:DeletedBy;null"`                        // Deleted by user ID
	DeletedAt      *time.Time `gorm:"column:DeletedAt;index;null"`                  // Deletion timestamp
	OrganizationId *int       `gorm:"column:OrganizationId;null"`                   // Organization ID

	// Relations
	Details []SphDetail `gorm:"foreignKey:SectionId"` // One-to-Many relationship with SPH Details
}

// TableName specifies the table name for SphSection
func (SphSection) TableName() string {
	return "SphSections"
}

func (s SphSection) MarshalJSON() ([]byte, error) {
	type Alias SphSection
	return json.Marshal(&struct {
		Details []SphDetail `json:"Details"`
		Alias
	}{
		Details: s.Details, // `detail` already handles its own custom marshaling
		Alias:   (Alias)(s),
	})
}

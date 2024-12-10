package models

import (
	"encoding/json"
	"time"
)

// BASection represents the BASection table in the database
type BASection struct {
	BASectionId    int        `gorm:"primaryKey;column:BASectionId;autoIncrement"` // Primary key
	BAID           *int       `gorm:"column:BAID;null"`                            // Reference to BA ID
	SphSectionId   int        `gorm:"column:SphSectionId;null"`                    // Reference to SphSectionId
	SectionName    *string    `gorm:"column:SectionName;type:varchar(255);null"`   // Section Name
	CreatedBy      *int       `gorm:"column:CreatedBy;null"`                       // Created by user ID
	CreatedAt      *time.Time `gorm:"column:CreatedAt;autoCreateTime"`             // Creation timestamp
	UpdatedBy      *int       `gorm:"column:UpdatedBy;null"`                       // Updated by user ID
	UpdatedAt      *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`        // Update timestamp
	DeletedBy      *int       `gorm:"column:DeletedBy;null"`                       // Deleted by user ID
	DeletedAt      *time.Time `gorm:"column:DeletedAt;index;null"`                 // Deletion timestamp
	OrganizationId *int       `gorm:"column:OrganizationId;null"`                  // Organization ID
	Details        []BADetail `gorm:"foreignKey:SectionID;references:BASectionId"` // Relationship to BADetail
}

// TableName specifies the table name for BASection
func (BASection) TableName() string {
	return "BASection"
}

func (s BASection) MarshalJSON() ([]byte, error) {
	type Alias BASection
	return json.Marshal(&struct {
		Details []BADetail `json:"Details"`
		Alias
	}{
		Details: s.Details, // `BADetail` already handles its own custom marshaling
		Alias:   (Alias)(s),
	})
}

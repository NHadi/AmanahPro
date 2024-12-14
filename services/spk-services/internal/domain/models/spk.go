package models

import (
	"encoding/json"
	"time"
)

// SPK represents the SPK table in the database
type SPK struct {
	SpkId          int         `gorm:"primaryKey;column:SpkId;autoIncrement"`        // Primary key
	SphId          int         `gorm:"column:SphId;not null"`                        // Foreign key to SPH
	ProjectId      *int        `gorm:"column:ProjectId;null"`                        // Foreign key to Projects
	ProjectName    *string     `gorm:"column:ProjectName;type:varchar(255);null"`    // Project name
	Subject        *string     `gorm:"column:Subject;type:varchar(255);null"`        // SPK subject
	Date           *CustomDate `gorm:"column:Date;type:date;null"`                   // Date of SPK
	TotalJasa      *float64    `gorm:"column:TotalJasa;type:decimal(15,2);null"`     // Total Jasa Cost
	TotalMaterial  *float64    `gorm:"column:TotalMaterial;type:decimal(15,2);null"` // Total Material Cost
	CreatedBy      *int        `gorm:"column:CreatedBy;null"`                        // Created by user ID
	CreatedAt      *time.Time  `gorm:"column:CreatedAt;autoCreateTime"`              // Creation timestamp
	UpdatedBy      *int        `gorm:"column:UpdatedBy;null"`                        // Updated by user ID
	UpdatedAt      *time.Time  `gorm:"column:UpdatedAt;autoUpdateTime;null"`         // Update timestamp
	DeletedBy      *int        `gorm:"column:DeletedBy;null"`                        // Deleted by user ID
	DeletedAt      *time.Time  `gorm:"column:DeletedAt;index;null"`                  // Deletion timestamp
	OrganizationId *int        `gorm:"column:OrganizationId;null"`                   // Organization ID

	// Relations
	Sections []SPKSection `gorm:"foreignKey:SpkId"` // One-to-Many relationship with SPK Sections
}

// TableName specifies the table name for SPK
func (SPK) TableName() string {
	return "SPK"
}

func (b SPK) MarshalJSON() ([]byte, error) {
	type Alias SPK
	return json.Marshal(&struct {
		Sections []SPKSection `json:"Sections"`
		Alias
	}{
		Sections: b.Sections, // `BASection` handles its own custom marshaling
		Alias:    (Alias)(b),
	})
}

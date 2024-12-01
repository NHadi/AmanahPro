package models

import "time"

// Sph represents the SPH table in the database
type Sph struct {
	SphId          int         `gorm:"primaryKey;column:SphId;autoIncrement"`       // Primary key
	ProjectId      *int        `gorm:"column:ProjectId;null"`                       // Foreign key to Projects
	ProjectName    *string     `gorm:"column:ProjectName;type:varchar(255);null"`   // Project name
	Subject        *string     `gorm:"column:Subject;type:varchar(255);null"`       // SPH subject
	Location       *string     `gorm:"column:Location;type:varchar(255);null"`      // Location of SPH
	Date           *CustomDate `gorm:"column:Date;type:date;null"`                  // Date of SPH
	RecepientName  *string     `gorm:"column:RecepientName;type:varchar(255);null"` // Name of the recipient
	CreatedBy      *int        `gorm:"column:CreatedBy;null"`                       // Created by user ID
	CreatedAt      *time.Time  `gorm:"column:CreatedAt;autoCreateTime"`             // Creation timestamp
	UpdatedBy      *int        `gorm:"column:UpdatedBy;null"`                       // Updated by user ID
	UpdatedAt      *time.Time  `gorm:"column:UpdatedAt;autoUpdateTime;null"`        // Update timestamp
	DeletedBy      *int        `gorm:"column:DeletedBy;null"`                       // Deleted by user ID
	DeletedAt      *time.Time  `gorm:"column:DeletedAt;index;null"`                 // Deletion timestamp
	OrganizationId *int        `gorm:"column:OrganizationId;null"`                  // Organization ID

	// Relations
	Sections []SphSection `gorm:"foreignKey:SphId"` // One-to-Many relationship with SPH Sections
}

// TableName specifies the table name for Sph
func (Sph) TableName() string {
	return "Sph"
}

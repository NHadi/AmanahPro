package models

import "time"

// Breakdown represents the Breakdowns table in the database
type Breakdown struct {
	BreakdownId    int         `gorm:"primaryKey;column:BreakdownId;autoIncrement"`   // Primary key
	ProjectId      int         `gorm:"column:ProjectId;not null"`                     // Foreign key to Projects
	ProjectName    string      `gorm:"column:ProjectName;type:varchar(255);not null"` // Project name
	Subject        string      `gorm:"column:Subject;type:varchar(255);not null"`     // Breakdown subject
	Location       *string     `gorm:"column:Location;type:varchar(255);null"`        // Optional location
	Date           *CustomDate `gorm:"column:Date;type:date;null"`                    // Breakdown date
	Total          *float64    `gorm:"column:Total;type:decimal(18,2);null"`          // Total price of the item
	CreatedBy      *int        `gorm:"column:CreatedBy;null"`                         // Created by user ID
	CreatedAt      *time.Time  `gorm:"column:CreatedAt;autoCreateTime"`               // Creation timestamp
	UpdatedBy      *int        `gorm:"column:UpdatedBy;null"`                         // Updated by user ID
	UpdatedAt      *time.Time  `gorm:"column:UpdatedAt;autoUpdateTime"`               // Update timestamp
	DeletedBy      *int        `gorm:"column:DeletedBy;null"`                         // Deleted by user ID
	DeletedAt      *time.Time  `gorm:"column:DeletedAt;index;null"`                   // Deletion timestamp
	OrganizationId *int        `gorm:"column:OrganizationId;null"`                    // Organization ID

	// Relations
	Sections []BreakdownSection `gorm:"foreignKey:BreakdownId"` // One-to-Many
}

// TableName specifies the table name for Breakdown
func (Breakdown) TableName() string {
	return "Breakdowns"
}

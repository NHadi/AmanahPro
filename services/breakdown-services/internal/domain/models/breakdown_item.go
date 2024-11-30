package models

import "time"

// BreakdownItem represents the BreakdownItems table in the database
type BreakdownItem struct {
	BreakdownItemId int        `gorm:"primaryKey;column:BreakdownItemId;autoIncrement"` // Primary key
	SectionId       int        `gorm:"column:SectionId;not null"`                       // Foreign key to BreakdownSections
	Description     string     `gorm:"column:Description;type:varchar(255);not null"`   // Description of the item
	UnitPrice       float64    `gorm:"column:UnitPrice;type:decimal(15,2);not null"`    // Unit price
	CreatedBy       *int       `gorm:"column:CreatedBy;null"`                           // Created by user ID
	CreatedAt       *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                 // Creation timestamp
	UpdatedBy       *int       `gorm:"column:UpdatedBy;null"`                           // Updated by user ID
	UpdatedAt       *time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`                 // Update timestamp
	DeletedBy       *int       `gorm:"column:DeletedBy;null"`                           // Deleted by user ID
	DeletedAt       *time.Time `gorm:"column:DeletedAt;index;null"`                     // Deletion timestamp
	OrganizationId  *int       `gorm:"column:OrganizationId;null"`                      // Organization ID
}

// TableName specifies the table name for BreakdownItem
func (BreakdownItem) TableName() string {
	return "BreakdownItems"
}

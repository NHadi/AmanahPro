package models

import "time"

// BreakdownItem represents the BreakdownItems table in the database
type MstBreakdownItem struct {
	MstBreakdownItemId    int        `gorm:"primaryKey;column:MstBreakdownItemId;autoIncrement"` // Primary key
	MstBreakdownSectionId int        `gorm:"column:MstBreakdownSectionId;not null"`              // Foreign key to MstBreakdownSectionId
	Description           string     `gorm:"column:Description;type:varchar(255);not null"`      // Description of the item
	UnitPrice             float64    `gorm:"column:UnitPrice;type:decimal(15,2);not null"`       // Unit price
	Sort                  int        `gorm:"column:Sort;not null"`                               // Sort
	CreatedBy             *int       `gorm:"column:CreatedBy;null"`                              // Created by user ID
	CreatedAt             *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                    // Creation timestamp
	UpdatedBy             *int       `gorm:"column:UpdatedBy;null"`                              // Updated by user ID
	UpdatedAt             *time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`                    // Update timestamp
	DeletedBy             *int       `gorm:"column:DeletedBy;null"`                              // Deleted by user ID
	DeletedAt             *time.Time `gorm:"column:DeletedAt;index;null"`                        // Deletion timestamp
	OrganizationId        *int       `gorm:"column:OrganizationId;null"`                         // Organization ID
}

// TableName specifies the table name for MstBreakdownItem
func (MstBreakdownItem) TableName() string {
	return "MstBreakdownItems"
}

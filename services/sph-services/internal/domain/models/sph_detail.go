package models

import "time"

// SphDetail represents the SPHDetails table in the database
type SphDetail struct {
	SphDetailId     int        `gorm:"primaryKey;column:SphDetailId;autoIncrement"`   // Primary key
	SectionId       int        `gorm:"column:SectionId;not null"`                     // Foreign key to SPH Section
	ItemDescription *string    `gorm:"column:ItemDescription;type:varchar(255);null"` // Item description
	Quantity        *float64   `gorm:"column:Quantity;type:decimal(10,2);null"`       // Quantity of the item
	Unit            *string    `gorm:"column:Unit;type:varchar(10);null"`             // Unit of the item
	UnitPrice       *float64   `gorm:"column:UnitPrice;type:decimal(15,2);null"`      // Unit price of the item
	DiscountPrice   *float64   `gorm:"column:DiscountPrice;type:decimal(15,2);null"`  // Discounted price of the item
	TotalPrice      *float64   `gorm:"column:TotalPrice;type:decimal(15,2);null"`     // Total price of the item
	CreatedBy       *int       `gorm:"column:CreatedBy;null"`                         // Created by user ID
	CreatedAt       time.Time  `gorm:"column:CreatedAt;autoCreateTime"`               // Creation timestamp
	UpdatedBy       *int       `gorm:"column:UpdatedBy;null"`                         // Updated by user ID
	UpdatedAt       *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`          // Update timestamp
	DeletedBy       *int       `gorm:"column:DeletedBy;null"`                         // Deleted by user ID
	DeletedAt       *time.Time `gorm:"column:DeletedAt;index;null"`                   // Deletion timestamp
	OrganizationId  *int       `gorm:"column:OrganizationId;null"`                    // Organization ID
}

// TableName specifies the table name for SphDetail
func (SphDetail) TableName() string {
	return "SphDetails"
}

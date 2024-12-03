package models

import "time"

// SPKDetail represents the SPK Details table in the database
type SPKDetail struct {
	DetailId          int        `gorm:"primaryKey;column:DetailId;autoIncrement"`             // Primary key
	SectionId         int        `gorm:"column:SectionId;not null"`                            // Foreign key to SPK Section
	SphItemId         *int       `gorm:"column:SphItemId;null"`                                // Reference to SPH Item (optional)
	Description       *string    `gorm:"column:Description;type:varchar(255);not null"`        // Item description
	Quantity          float64    `gorm:"column:Quantity;type:decimal(10,2);not null"`          // Item quantity
	Unit              *string    `gorm:"column:Unit;type:varchar(10);not null"`                // Unit of measurement
	UnitPriceJasa     float64    `gorm:"column:UnitPriceJasa;type:decimal(15,2);not null"`     // Unit price for Jasa
	TotalJasa         float64    `gorm:"column:TotalJasa;type:decimal(15,2);not null"`         // Total Jasa cost
	UnitPriceMaterial float64    `gorm:"column:UnitPriceMaterial;type:decimal(15,2);not null"` // Unit price for Material
	TotalMaterial     float64    `gorm:"column:TotalMaterial;type:decimal(15,2);not null"`     // Total Material cost
	CreatedBy         *int       `gorm:"column:CreatedBy;null"`                                // Created by user ID
	CreatedAt         *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                      // Creation timestamp
	UpdatedBy         *int       `gorm:"column:UpdatedBy;null"`                                // Updated by user ID
	UpdatedAt         *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`                 // Update timestamp
	DeletedBy         *int       `gorm:"column:DeletedBy;null"`                                // Deleted by user ID
	DeletedAt         *time.Time `gorm:"column:DeletedAt;index;null"`                          // Deletion timestamp
	OrganizationId    *int       `gorm:"column:OrganizationId;null"`                           // Organization ID
}

// TableName specifies the table name for SPKDetail
func (SPKDetail) TableName() string {
	return "SPKDetails"
}

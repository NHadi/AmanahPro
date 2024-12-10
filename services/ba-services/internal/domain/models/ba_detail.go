package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// BADetail represents the BADetails table in the database
type BADetail struct {
	DetailId         int        `gorm:"primaryKey;column:DetailId;autoIncrement"`       // Primary key
	SectionID        *int       `gorm:"column:SectionID;null"`                          // Reference to Section ID
	SphItemId        *int       `gorm:"column:SphItemId;null"`                          // Reference to SPH Item
	ItemName         *string    `gorm:"column:ItemName;type:varchar(255);null"`         // Item Name
	Quantity         float64    `gorm:"column:Quantity;type:decimal(10,2);not null"`    // Quantity
	Unit             *string    `gorm:"column:Unit;type:varchar(10);null"`              // Unit
	UnitPrice        *float64   `gorm:"column:UnitPrice;type:decimal(15,2);null"`       // Unit price of the item
	DiscountPrice    *float64   `gorm:"column:DiscountPrice;type:decimal(15,2);null"`   // Discounted price of the item
	TotalPrice       *float64   `gorm:"column:TotalPrice;type:decimal(15,2);null"`      // Discounted price of the item
	WeightPercentage *float64   `gorm:"column:WeightPercentage;type:decimal(5,2);null"` // Weight Percentage
	CreatedBy        *int       `gorm:"column:CreatedBy;null"`                          // Created by user ID
	CreatedAt        *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                // Creation timestamp
	UpdatedBy        *int       `gorm:"column:UpdatedBy;null"`                          // Updated by user ID
	UpdatedAt        *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`           // Update timestamp
	DeletedBy        *int       `gorm:"column:DeletedBy;null"`                          // Deleted by user ID
	DeletedAt        *time.Time `gorm:"column:DeletedAt;index;null"`                    // Deletion timestamp
	OrganizationId   *int       `gorm:"column:OrganizationId;null"`                     // Organization ID

	Progress []BAProgress `gorm:"foreignKey:DetailId;references:DetailId"` // Relationship to BAProgress
}

// TableName specifies the table name for BADetail
func (BADetail) TableName() string {
	return "BADetails"
}

func (d BADetail) MarshalJSON() ([]byte, error) {
	type Alias BADetail
	return json.Marshal(&struct {
		UnitPrice     string `json:"UnitPrice"`
		DiscountPrice string `json:"DiscountPrice"`
		TotalPrice    string `json:"TotalPrice"`
		Alias
	}{
		UnitPrice:     formatFloat(d.UnitPrice),
		DiscountPrice: formatFloat(d.DiscountPrice),
		TotalPrice:    formatFloat(d.TotalPrice),
		Alias:         (Alias)(d),
	})
}

// Helper function to format float64 values with two decimal places
func formatFloat(value *float64) string {
	if value == nil {
		return "0.00"
	}
	return fmt.Sprintf("%.2f", *value)
}

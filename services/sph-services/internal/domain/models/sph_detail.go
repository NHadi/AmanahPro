package models

import (
	"encoding/json"
	"fmt"
	"time"
)

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

func (d SphDetail) MarshalJSON() ([]byte, error) {
	type Alias SphDetail
	return json.Marshal(&struct {
		Quantity      string `json:"Quantity"`
		UnitPrice     string `json:"UnitPrice"`
		DiscountPrice string `json:"DiscountPrice"`
		TotalPrice    string `json:"TotalPrice"`
		Alias
	}{
		Quantity:      formatFloat(d.Quantity),
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

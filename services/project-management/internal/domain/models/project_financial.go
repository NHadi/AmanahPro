package models

import (
	"time"
)

type ProjectFinancial struct {
	ID              int        `gorm:"column:ID;primaryKey;autoIncrement"`                // Primary key
	ProjectID       int        `gorm:"column:ProjectID;not null"`                         // Foreign key to Projects table
	ProjectUserID   *int       `gorm:"column:ProjectUserId"`                              // Nullable foreign key to ProjectUser
	TransactionDate CustomDate `gorm:"column:TransactionDate;type:date;not null"`         // Transaction date
	Description     string     `gorm:"column:Descrtiption;type:nvarchar(255);not null"`   // Transaction description
	Amount          float64    `gorm:"column:Amount;type:decimal(18,2);not null"`         // Amount (income/expense)
	AmountDeviden   *float64   `gorm:"column:AmountDeviden;type:decimal(18,2);null"`      // Amount (income/expense)
	TransactionType string     `gorm:"column:TransactionType;type:nvarchar(50);not null"` // Income or Expense
	Category        string     `gorm:"column:Category;type:nvarchar(255);not null"`       // Category
	CreatedAt       *time.Time `gorm:"column:CreatedAt"`                                  // Created timestamp
	CreatedBy       *int       `gorm:"column:CreatedBy"`                                  // Created by user ID
	UpdatedAt       *time.Time `gorm:"column:UpdatedAt"`                                  // Updated timestamp
	UpdatedBy       *int       `gorm:"column:UpdatedBy"`                                  // Updated by user ID
	DeletedAt       *time.Time `gorm:"column:DeletedAt;index"`                            // Soft delete timestamp
	DeletedBy       *int       `gorm:"column:DeletedBy"`                                  // Deleted by user ID
	OrganizationID  *int       `gorm:"column:OrganizationID"`                             // Organization ID

	// Relationships
	Project     *Project     `gorm:"foreignKey:ProjectID;references:ProjectID"` // Relationship to Projects table
	ProjectUser *ProjectUser `gorm:"foreignKey:ProjectUserID;references:ID"`    // Relationship to ProjectUser table
}

// TableName overrides the table name in GORM
func (ProjectFinancial) TableName() string {
	return "ProjectFinancial"
}

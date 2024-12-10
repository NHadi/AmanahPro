package models

import "time"

// BAProgress represents the BAProgress table in the database
type BAProgress struct {
	BAProgressId               int        `gorm:"primaryKey;column:BAProgressId;autoIncrement"`             // Primary key
	DetailId                   int        `gorm:"column:DetailId;not null"`                                 // Reference to Detail ID
	ProgressPreviousM2         *float64   `gorm:"column:ProgressPreviousM2;type:decimal(10,2);null"`        // Previous progress in M2
	ProgressPreviousPercentage *float64   `gorm:"column:ProgressPreviousPercentage;type:decimal(5,2);null"` // Previous progress percentage
	ProgressCurrentM2          *float64   `gorm:"column:ProgressCurrentM2;type:decimal(10,2);null"`         // Current progress in M2
	ProgressCurrentPercentage  *float64   `gorm:"column:ProgressCurrentPercentage;type:decimal(5,2);null"`  // Current progress percentage
	CreatedBy                  *int       `gorm:"column:CreatedBy;null"`                                    // Created by user ID
	CreatedAt                  *time.Time `gorm:"column:CreatedAt;autoCreateTime"`                          // Creation timestamp
	UpdatedBy                  *int       `gorm:"column:UpdatedBy;null"`                                    // Updated by user ID
	UpdatedAt                  *time.Time `gorm:"column:UpdatedAt;autoUpdateTime;null"`                     // Update timestamp
	DeletedBy                  *int       `gorm:"column:DeletedBy;null"`                                    // Deleted by user ID
	DeletedAt                  *time.Time `gorm:"column:DeletedAt;index;null"`                              // Deletion timestamp
	OrganizationId             *int       `gorm:"column:OrganizationId;null"`                               // Organization ID
}

// TableName specifies the table name for BAProgress
func (BAProgress) TableName() string {
	return "BAProgress"
}

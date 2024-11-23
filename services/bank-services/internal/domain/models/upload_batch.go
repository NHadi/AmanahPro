package models

import (
	"time"
)

type UploadBatch struct {
	BatchID    uint       `gorm:"column:BatchID;primaryKey;autoIncrement"`      // Maps to [BatchID]
	AccountID  uint       `gorm:"column:AccountID;not null"`                    // Maps to [AccountID]
	FileName   string     `gorm:"column:FileName;size:255;not null"`            // Maps to [FileName]
	Month      uint       `gorm:"column:Month;not null"`                        // Maps to [Month]
	Year       uint       `gorm:"column:Year;not null"`                         // Maps to [Year]
	UploadedBy string     `gorm:"column:UploadedBy;size:255;not null"`          // Maps to [UploadedBy]
	UploadDate time.Time  `gorm:"column:UploadDate;autoCreateTime"`             // Maps to [UploadDate]
	CreatedAt  time.Time  `gorm:"column:CreatedAt;autoCreateTime"`              // Maps to [CreatedAt]
	UpdatedAt  time.Time  `gorm:"column:UpdatedAt;autoUpdateTime"`              // Maps to [UpdatedAt]
	DeletedAt  *time.Time `gorm:"column:DeletedAt" json:"deleted_at,omitempty"` // Replace gorm.DeletedAt

	// Relationships
	Account BankAccount `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}

func (UploadBatch) TableName() string {
	return "UploadBatch" // Maps to [AmanahDb].[dbo].[UploadBatch]
}

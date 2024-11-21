package models

import (
	"time"

	"gorm.io/gorm"
)

type UploadBatch struct {
	BatchID      uint           `gorm:"column:BatchID;primaryKey;autoIncrement"` // Maps to [BatchID]
	AccountID    uint           `gorm:"column:AccountID;not null"`               // Maps to [AccountID]
	FileName     string         `gorm:"column:FileName;size:255;not null"`       // Maps to [FileName]
	PeriodeStart time.Time      `gorm:"column:PeriodeStart;not null"`            // Maps to [PeriodeStart]
	PeriodeEnd   time.Time      `gorm:"column:PeriodeEnd;not null"`              // Maps to [PeriodeEnd]
	UploadedBy   string         `gorm:"column:UploadedBy;size:255;not null"`     // Maps to [UploadedBy]
	UploadDate   time.Time      `gorm:"column:UploadDate;autoCreateTime"`        // Maps to [UploadDate]
	CreatedAt    time.Time      `gorm:"column:CreatedAt;autoCreateTime"`         // Maps to [CreatedAt]
	UpdatedAt    time.Time      `gorm:"column:UpdatedAt;autoUpdateTime"`         // Maps to [UpdatedAt]
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`                 // Maps to [deleted_at]

	// Relationships
	Account BankAccount `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}

func (UploadBatch) TableName() string {
	return "UploadBatch" // Maps to [AmanahDb].[dbo].[UploadBatch]
}

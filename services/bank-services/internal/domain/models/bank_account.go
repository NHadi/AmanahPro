package models

import (
	"time"

	"gorm.io/gorm"
)

type BankAccount struct {
	AccountID    uint           `gorm:"primaryKey;autoIncrement"`
	NoRekening   string         `gorm:"size:50;not null;unique"`
	Nama         string         `gorm:"size:255;not null"`
	KodeMataUang string         `gorm:"size:10;not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete
}

func (BankAccount) TableName() string {
	return "BankAccounts" // Match the exact name from the database
}

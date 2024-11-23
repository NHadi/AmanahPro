package models

import (
	"time"
)

type BankAccount struct {
	AccountID    uint       `gorm:"primaryKey;autoIncrement"`
	NoRekening   string     `gorm:"size:50;not null;unique"`
	Nama         string     `gorm:"size:255;not null"`
	KodeMataUang string     `gorm:"size:10;not null"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `gorm:"column:DeletedAt" json:"deleted_at,omitempty"` // Replace gorm.DeletedAt
}

func (BankAccount) TableName() string {
	return "BankAccounts" // Match the exact name from the database
}

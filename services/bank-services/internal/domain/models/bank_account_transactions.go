package models

import (
	"time"
)

type BankAccountTransactions struct {
	ID         uint      `gorm:"column:ID;primaryKey;autoIncrement"`            // Maps to [ID]
	AccountID  uint      `gorm:"column:AccountID;not null"`                     // Maps to [AccountID]
	BatchID    uint      `gorm:"column:BatchID;not null"`                       // Maps to [BatchID]
	Tanggal    time.Time `gorm:"column:Tanggal;not null"`                       // Maps to [Tanggal]
	Keterangan string    `gorm:"column:Keterangan;type:text;not null"`          // Maps to [Keterangan]
	Cabang     string    `gorm:"column:Cabang;size:50;not null"`                // Maps to [Cabang]
	Credit     float64   `gorm:"column:Credit;type:decimal(18,2);default:0.00"` // Maps to [Credit]
	Debit      float64   `gorm:"column:Debit;type:decimal(18,2);default:0.00"`  // Maps to [Debit]
	Saldo      float64   `gorm:"column:Saldo;type:decimal(18,2);default:0.00"`  // Maps to [Saldo]
	CreatedAt  time.Time `gorm:"column:CreatedAt;autoCreateTime"`               // Maps to [CreatedAt]
	UpdatedAt  time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`               // Maps to [UpdatedAt]

	// Relationships
	Account BankAccount `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
	Batch   UploadBatch `gorm:"foreignKey:BatchID;constraint:OnDelete:CASCADE"`
}

func (BankAccountTransactions) TableName() string {
	return "BankAccountTransactions" // Maps to [AmanahDb].[dbo].[BankAccountTransactions]
}

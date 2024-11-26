package dto

import "time"

type BankAccountTransactionDTO struct {
	ID             uint    `json:"ID"`
	AccountID      uint    `json:"AccountID"`
	BatchID        uint    `json:"BatchID"`
	Tanggal        string  `json:"Tanggal"` // Use string for ISO date format
	Keterangan     string  `json:"Keterangan"`
	Cabang         string  `json:"Cabang"`
	Credit         float64 `json:"Credit"`
	Debit          float64 `json:"Debit"`
	Saldo          float64 `json:"Saldo"`
	OrganizationId uint    `json:"OrganizationId"`
	UpdatedAt      time.Time
}

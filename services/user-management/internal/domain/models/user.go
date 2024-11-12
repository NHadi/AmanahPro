package models

import (
	"time"
)

type User struct {
	UserID    string    `gorm:"type:uniqueidentifier;primaryKey;default:NEWID()"`
	Username  string    `gorm:"type:nvarchar(50);unique;not null"`
	Email     string    `gorm:"type:nvarchar(100);unique;not null"`
	Password  string    `gorm:"type:nvarchar(255);not null"`
	Status    bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"default:GETDATE()"`
	UpdatedAt time.Time `gorm:"default:GETDATE()"`
}

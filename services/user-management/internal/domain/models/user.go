package models

import (
	"time"
)

type User struct {
	UserID         int       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username       string    `gorm:"type:nvarchar(100);unique;not null"`
	Email          string    `gorm:"type:nvarchar(100);unique;not null"`
	Password       string    `gorm:"type:nvarchar(100);not null"`
	Status         bool      `gorm:"default:true"`
	CreatedAt      time.Time `gorm:"default:GETDATE()"`
	UpdatedAt      time.Time `gorm:"default:GETDATE()"`
	OrganizationId *int      `gorm:"column:organization_id;"`
}

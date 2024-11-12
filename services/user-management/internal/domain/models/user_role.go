package models

import (
	"time"
)

type UserRole struct {
	UserID     string    `gorm:"type:uniqueidentifier;primaryKey"`
	RoleID     string    `gorm:"type:uniqueidentifier;primaryKey"`
	AssignedAt time.Time `gorm:"default:GETDATE()"`
}

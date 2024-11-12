package models

import (
	"time"
)

type RoleMenu struct {
	RoleID     string    `gorm:"type:uniqueidentifier;primaryKey"`
	MenuID     string    `gorm:"type:uniqueidentifier;primaryKey"`
	Permission string    `gorm:"type:nvarchar(10);not null;"`
	AssignedAt time.Time `gorm:"default:GETDATE()"`
}

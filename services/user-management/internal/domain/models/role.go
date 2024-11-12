package models

import (
	"time"
)

type Role struct {
	RoleID      string    `gorm:"column:role_id;type:uniqueidentifier;primaryKey"`
	RoleName    string    `gorm:"column:role_name;type:nvarchar(50);unique;not null"`
	Description string    `gorm:"column:description;type:text"`
	CreatedAt   time.Time `gorm:"column:created_at;default:GETDATE()"`
}

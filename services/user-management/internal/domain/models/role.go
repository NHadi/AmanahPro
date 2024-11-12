package models

import (
	"time"
)

type Role struct {
	RoleID      int       `gorm:"column:role_id;primaryKey;autoIncrement"`
	RoleName    string    `gorm:"column:name;type:varchar(100);unique;not null"`
	Description string    `gorm:"column:description;type:varchar(max)"`
	CreatedAt   time.Time `gorm:"column:created_at;default:GETDATE()"`
}

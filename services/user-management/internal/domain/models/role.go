package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	RoleID      uuid.UUID `gorm:"type:uniqueidentifier;primaryKey"`
	RoleName    string    `gorm:"type:nvarchar(50);unique;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"default:GETDATE()"`
}

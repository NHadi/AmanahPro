package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	UserID     uuid.UUID `gorm:"type:uniqueidentifier;primaryKey"`
	RoleID     uuid.UUID `gorm:"type:uniqueidentifier;primaryKey"`
	AssignedAt time.Time `gorm:"default:GETDATE()"`
}

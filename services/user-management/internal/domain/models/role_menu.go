package models

import (
	"time"

	"github.com/google/uuid"
)

type RoleMenu struct {
	RoleID     uuid.UUID `gorm:"type:uniqueidentifier;primaryKey"`
	MenuID     uuid.UUID `gorm:"type:uniqueidentifier;primaryKey"`
	Permission string    `gorm:"type:nvarchar(10);not null;"`
	AssignedAt time.Time `gorm:"default:GETDATE()"`
}

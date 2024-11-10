package models

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	MenuID    uuid.UUID  `gorm:"type:uniqueidentifier;primaryKey;default:NEWID()"`
	ParentID  *uuid.UUID `gorm:"type:uniqueidentifier"`
	MenuName  string     `gorm:"type:nvarchar(50);not null"`
	Path      string     `gorm:"type:nvarchar(100);not null"`
	Icon      string     `gorm:"type:nvarchar(50)"`
	Order     int        `gorm:"type:int"`
	CreatedAt time.Time  `gorm:"default:GETDATE()"`
}

package models

// SysOrganization model
type SysOrganization struct {
	OrganizationID int    `gorm:"column:OrganizationId;primaryKey;autoIncrement"`
	Name           string `gorm:"type:nvarchar(250);not null"`
	Description    string `gorm:"type:nvarchar(250)"`
	Address        string `gorm:"type:text"`
	LogoURL        string `gorm:"type:nvarchar(50)"`
}

// TableName overrides the table name in GORM
func (SysOrganization) TableName() string {
	return "SysOrganization"
}

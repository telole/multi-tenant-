package models

import (
	"time"
)

type Project struct {
	id uint `gorm:"primarykey"`
	TenantID string `gorm:"type:uuid;not null"`
	Name string
	Description string
	CreatedAt time.Time

	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID"`
}
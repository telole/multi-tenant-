package models

import "time"

type Student struct {
	ID uint `gorm:"primarykey"`
	TenantID string `gorm:"type:uuid;not null"`
	CompanyID uint `gorm:"not null"`
	Name string
	Class string
	startDate time.Time
	endDate time.Time

	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID"`
	Company Company `gorm:"foreignKey:CompanyID;references:ID"`
}
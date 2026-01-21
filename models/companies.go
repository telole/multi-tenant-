package models

import "time"

type Companie struct { 
	ID uint `gorm:"primarykey"`
	TenantID string `gorm:"type:uuid;not null"`
	Name string 
	Address string
	CreatedAt time.Time

	Tennant Tenant `gorm:"foreignKey:TenantID;references:ID"`
}
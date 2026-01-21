package models

import "time"

type Ticket struct { 
	id uint `gorm:"primarykey"`
	TenantID string `gorm:"type:uuid;not null"`
	UserID uint `gorm:"not null"`
	title string
	Description string
	Status string
	CreatedAt time.Time

	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID"`
	User User `gorm:"foreignKey:UserID;references:ID"`
}
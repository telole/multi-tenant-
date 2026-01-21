package models


type Permission struct {
	ID uint `gorm:"primarykey"`
	Code string `gorm:"size:100;not null;unique"`
	Description string

	Roles []Role `gorm:"many2many:role_permissions;"`
}
package models

type UserRole struct{
	UserID uint `gorm:"primarykey"`
	RoleID uint `gorm:"primarykey"`
}
package models

type RolePermission struct{
	role_id uint `gorm:"primarykey"`
	Permisson_id uint `gorm:"primarykey"`
}
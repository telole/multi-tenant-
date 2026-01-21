package models

type Role struct{

	ID uint `gorm:"primarykey"`
	TenantID string `gorm:"type:uuid;not null"`
	Name string

	//relasinya adalah hahahahhaahahahhahaaha
	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID"`
	User []User `gorm:"many2many:user_roles;"`

}
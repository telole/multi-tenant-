package models

import "time"

type Tenant struct { 
	ID string `gorm:"type:uuid;primarykey"`
	Name string
	Slug string
	IsActive bool
	CreatedAt time.Time
}
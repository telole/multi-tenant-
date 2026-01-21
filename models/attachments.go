package models

import "time"

type Attachment struct {
	ID        uint      `gorm:"primaryKey"`
	TicketID  uint      `gorm:"not null"`
	FileName  string    `gorm:"not null"`
	FileUrl   string    `gorm:"not null"`
	CreatedAt time.Time

	Ticket Ticket `gorm:"foreignKey:TicketID;references:ID"`
}
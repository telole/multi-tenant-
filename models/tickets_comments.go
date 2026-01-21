package models

import "time"

type TicketComment struct {
	ID        uint      `gorm:"primaryKey"`
	TicketID  uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Comment   string
	CreatedAt time.Time

	Ticket Ticket `gorm:"foreignKey:TicketID;references:ID"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
}

package request

type TicketComment struct {
	Comment string `json:"comment" validate:"required"`
}
package request

type CreateTicketRequest struct { 
	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateTicketRequest struct { 
	// Title string `json:"title" validate:"required"`
	// Description string `json:"description" validate:"required"`
	Status string `json:"status" validate:"required,oneof=open in_progress closed"`
}
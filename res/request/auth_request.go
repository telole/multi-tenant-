package request

type RegisterRequest struct {
	TenantID  string `json:"tenant_id" validate:"required,uuid"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct { 
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
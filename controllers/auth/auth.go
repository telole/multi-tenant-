package auth

import (
	"backend/models"
	"backend/res/request"
	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
type AuthController struct { 
	DB *gorm.DB
}

func (a *AuthController) Register(c echo.Context) error {
	req := new (request.RegisterRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, echo.Map{"error" : "invalid request"})
	}

	var tenant models.Tenant
	if err := a.DB.Where("id = ?", req.TenantID).First(&tenant).Error; err != nil {
		return c.JSON(400, echo.Map{
			"error": "tenant not found",
		})
	}

	var existing models.User
	a.DB.Where("email = ? AND tenant_id = ?", req.Email, req.TenantID).First(&existing)
	if existing.ID != 0 { 
		return c.JSON(400, echo.Map{"error" : "Email already exists"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password),10)

	user := models.User{ 
		TenantID: tenant.ID,
		Name: req.Name,
		Email: req.Email,
		Password: string(hashed),
		IsActive: true,
	}

	if err := a.DB.Create(&user).Error; err != nil { 
		return c.JSON(500, echo.Map{"error" : "failed to create user"})
	}

	return c.JSON(201, echo.Map{"message" : "user created successfully", "user" : user})
}
	
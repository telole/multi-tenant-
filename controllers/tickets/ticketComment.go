package tickets

import (
	"backend/models"
	"backend/res/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TicketCommentController struct {
	DB *gorm.DB
}

func (t *TicketCommentController) isUserAdmin(c echo.Context) bool {
	roles, ok := c.Get("roles").([]string)
	if !ok {
		return false
	}

	for _, role := range roles {
		if role == "admin" || role == "Super Admin" {
			return true
		}
	}

	return false
}


func (t *TicketCommentController) CreateComment(c echo.Context) error {
	req := new(request.TicketComment)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, echo.Map{"error": "invalid request"})
	}

	if req.Comment == "" {
		return c.JSON(400, echo.Map{"error" : "comment cannot be empty"})
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(400, echo.Map{"error": "user_id not found"})
	}

	tenantID, ok := c.Get("tenant_id").(string)
	if !ok {
		return c.JSON(400, echo.Map{"error": "tenant_id not found"})
	}

	TicketID := c.Param("ticket_id")

	var ticket models.Ticket

	query := t.DB.Where("id = ? AND tenant_id = ?", TicketID, tenantID)
	if !t.isUserAdmin(c) { 
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&ticket).Error; err != nil {
		return c.JSON(404, echo.Map{"error": "ticket not found"})
	}

	comment := models.TicketComment{ 
		TicketID:  ticket.ID,
		UserID:    userID,
		Comment:   req.Comment,
	}

	if err := t.DB.Create(&comment).Error; err != nil {
		return c.JSON(500, echo.Map{"error": "failed to create comment"})
	}

	t.DB.Preload("User").First(&comment, comment.ID)

	return c.JSON(201, echo.Map{
		"message": "comment created successfully",
		"comment": comment,
	})

}



package tickets

import (
	"backend/models"
	"backend/res/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TicketController struct {
	DB *gorm.DB
}

func (t*TicketController) CreateTicket(c echo.Context) error { 
	req := new(request.CreateTicketRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, echo.Map{"error" : "invalid request"})
	}

	UserID := c.Get("user_id").(uint)
	TenantID := c.Get("tenant_id").(string)

	ticket := models.Ticket{ 
		TenantID: TenantID,
		UserID: UserID,
		Title: req.Title,
		Description: req.Description,
		Status: "open",
	}

	if err := t.DB.Create(&ticket).Error; err != nil { 
		return c.JSON(500, echo.Map{"error" : "failed to create ticket"})
	}

	return c.JSON(201, echo.Map{"message" : "ticket created successfully", "status" : ticket})
}   

func (t*TicketController) GetTickets(c echo.Context) error { 
	tentantID := c.Get("tentant_id").(string)

	var tickets []models.Ticket

	if err := t.DB.Where("tenant_id = ?", tentantID).Order("created_at DESC").Find(&tickets).Error; err != nil {
		return c.JSON(500, echo.Map{"error" : "failed to fetch tickets"})
	}

	return c.JSON(200, echo.Map{"tickets" : tickets})
}

func (t*TicketController) GetTicketByID(c echo.Context) error { 
	TenantiD := c.Get("tenant_id").(string)
	ticketID := c.Param("id")

	var ticket models.Ticket

	if err := t.DB.Where("tenant_id = ? AND id = ?", TenantiD, ticketID).First(&ticket).Error; err != nil {
		return c.JSON(404, echo.Map{"error" : "ticket not found"})
	}

	return c.JSON(200, echo.Map{"ticket" : ticket})
}

func (t*TicketController) UpdateTicketStatus(c echo.Context) error { 
	req := new(request.UpdateTicketRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(401, echo.Map{"error" : "invalid request"})
	}

	tenantID := c.Get("tenant_id")
	id := c.Param("id")

	if err := t.DB.Model(&models.Ticket{}).Where("id = ? AND tenant_id = ?", id, tenantID).Update("status", req.Status).Error;
	err != nil {
		return c.JSON(401, echo.Map{"messsage" : "rekwes mu eror mas"})
	}
	return c.JSON(200, echo.Map{"message" : "ticket status updated", "status" : req.Status}) 
}
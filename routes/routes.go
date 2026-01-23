package routes

import (
	"backend/controllers/auth"
	"backend/controllers/profile"
	"backend/controllers/tickets"
	"backend/res/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	//definisi Controller
	authController := &auth.AuthController{DB: db}
	profileController := &profile.ProfileController{DB: db}
	ticketController := &tickets.TicketController{DB: db}

	//AuthController Routes
	auth := e.Group("/auth") 
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)

	authProtected := e.Group("/auth", middleware.AuthMiddleware)
	authProtected.POST("/logout", authController.Logout)


	api := e.Group("/api", middleware.AuthMiddleware)
	api.GET("/me", profileController.GetProfile)

	ticket := api.Group("/tickets")
	{
		ticket.POST("", ticketController.CreateTicket)
		ticket.GET("", ticketController.GetTickets)
		ticket.GET("/:id", ticketController.GetTicketByID)
		ticket.PUT("/:id/status", ticketController.UpdateTicketStatus)
		// ticket.DELETE("/:id", ticketController.DeleteTicket)
	}
}
	
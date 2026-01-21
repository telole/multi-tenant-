package main

import (
	"log"

	"backend/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=db_juan port=5050 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	e := echo.New()

	routes.InitRoutes(e, db)


	e.Logger.Fatal(e.Start(":3000"))
}

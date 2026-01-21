package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=db_juan port=5050 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	fmt.Println("Connected to the database successfully.")
}

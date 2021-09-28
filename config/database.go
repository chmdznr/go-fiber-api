package config

import (
	"fiber-api/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func DbConnect()  {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))))
	if err != nil {
		panic(fmt.Errorf("Fatal error connect DB: %w \n", err))
	}

	DB = db
	fmt.Println("Connection has been established successfully.")

	err = db.AutoMigrate(&models.Users{})
	if err != nil {
		panic(fmt.Errorf("Fatal error automigrate DB: %w \n", err))
	}

}

package database

import (
	"fmt"
	"log"
	"todos/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm.DB
var DB *gorm.DB
var err error

// Connect to database
func Connect() {
	// dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	// config.DbHost, config.DbPort, config.DbUsername, config.DbName, config.DbPassword)
	// dsn := "user:postgres@tcp(127.0.0.1:5432)/bk?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUsername,
		config.DbPassword,
		config.DbName,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	fmt.Println("Database connected")
}

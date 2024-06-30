package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"idotno.fr/echo/models"
)

func GetDbConnection() *gorm.DB {
	HOST := GetEnv("ECHO_DB_HOST", "localhost")
	USER := GetEnv("ECHO_DB_USER", "postgres")
	PASS := GetEnv("ECHO_DB_PASS", "postgres")
	NAME := GetEnv("ECHO_DB_NAME", "echo")
	PORT := GetEnv("ECHO_DB_PORT", "5432")

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, NAME, PORT)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	} else {
		log.Println("Database connection established")
	}

	db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	} else {
		log.Println("Database migration completed")
	}

	return db
}

package services

import (
	"golang.org/x/crypto/bcrypt"
	"idotno.fr/echo/models"
)

func CreateUser(username string, password string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Add user to DB
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	res := dbConn.Create(&user)
	return res.Error
}

func ListUsers() ([]models.User, error) {
	var users []models.User

	err := dbConn.Find(&users).Error

	return users, err
}

func GetUser(id uint) (models.User, error) {
	var user models.User

	err := dbConn.First(&user, id).Error

	return user, err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	err := dbConn.Where("username = ?", username).First(&user).Error

	return user, err
}

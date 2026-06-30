package repository

import (
	"mini-store-api/config"
	"mini-store-api/models"
)

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := config.DB.Where("email= ?", email).First(&user).Error
	return user, err
}

func CreateUser(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error
	return user, err
}
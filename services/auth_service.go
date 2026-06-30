package services

import (
	"errors"
	"mini-store-api/models"
	"mini-store-api/repository"
	"mini-store-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func SignupService(input models.SignupInput) (string, models.User, error) {
	//check if user alredy existes
	_, err := repository.FindUserByEmail(input.Email)
	if err == nil {
		return "", models.User{}, errors.New("email alredy registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", models.User{}, errors.New("failed to hash password")
	}

	user := models.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := repository.CreateUser(user)
	if err != nil {
		return "", models.User{}, errors.New("failed to create a user")
	}

	token, err := utils.GenerateJWT(createdUser.ID, createdUser.Email)
	if err != nil {
		return "", models.User{}, errors.New("failed generate a token")
	}

	return token, createdUser, nil
}

func LoginService(input models.LoginInput) (string, models.User, error) {
	user, err := repository.FindUserByEmail(input.Email)
	if err != nil {
		return "", models.User{}, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", models.User{}, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", models.User{}, errors.New("failed to generate token")
	}
	return token, user, nil
}
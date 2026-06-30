package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique; not null"`
    Password string `json:"-"`
}

type SignupInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInout struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
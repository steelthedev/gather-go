package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username     string `json:"username"`
	Email        string `json:"email"`
	Full_name    string `json:"full_name"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
}

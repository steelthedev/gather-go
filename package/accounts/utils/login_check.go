package utils

import (
	"gather-go/package/accounts/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyPassword(password, hashedpwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpwd), []byte(password))
}

func LoginCheck(db *gorm.DB, email string, password string) (string, error) {
	var err error

	user := models.User{}

	err = db.Where("Email=?", email).First(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {

		return "", err
	}

	token, err := GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

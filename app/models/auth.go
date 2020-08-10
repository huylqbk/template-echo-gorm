package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"red-coins/config"
)

type JwtClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

func AuthLogin(email string, password string) *User {
	user := new(User)
	res := config.DB.Debug().Where("email = ?", email).First(&user)
	
	if res.Error == nil {
		err := VerifyPassword(user.Password, password)

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return nil
		}

		return user
	}

	return nil
}

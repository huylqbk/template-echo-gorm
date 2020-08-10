package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"os"
	"red-coins/app/models"
	"time"
)

func AuthMakeToken(user *models.User) (string, error) {
	secret := os.Getenv("API_SECRET")
	claims := models.JwtClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

func AuthGetUser(c echo.Context) *models.User {
	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(*models.JwtClaims)

	user := models.UserShow(claims.ID)
	
	if user != nil {
		return user
	}
	
	return nil
}

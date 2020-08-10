package controllers

import (
	"github.com/labstack/echo/v4"
	"red-coins/app/helpers"
	"red-coins/app/models"
)

// Index redCoins
// @Summary Home Page
// @Description Display Home Page
// @Tags Home
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router / [get]
func Index(c echo.Context) error {
	return c.JSON(200, "Welcome to Echo")
}

// Login redCoins
// @Summary Login
// @Description Login User in API
// @Tags Auth
// @Produce  json
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /api/login [post]
func Login(c echo.Context) error {
	login := models.Login{}

	if err := c.Bind(&login); err != nil {
		return c.JSON(422, err)
	}

	err := helpers.Validate(&login)

	if err != nil {
		return c.JSON(422, err)
	}

	user := models.AuthLogin(login.Email, login.Password)

	if user != nil {
		token, err := helpers.AuthMakeToken(user)
		if err != nil {
			return c.JSON(500, "Server Error")
		}
		return c.JSON(200, map[string]string{"token": token})
	}

	return c.JSON(404, "Not Found")
}

// Register redCoins
// @Summary Register
// @Description User Register
// @Tags Auth
// @Produce  json
// @Success 200 {string} string
// @Failure 404 {string} string
// @Router /api/register [post]
func Register(c echo.Context) error {
	user := helpers.AuthGetUser(c)
	if user != nil {
		return c.JSON(200, user)
	}

	return c.JSON(401, "Unauthorized")
}

// Logout redCoins
// @Summary Logout
// @Description User Logout
// @Tags Auth
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 401 {string} string
// @Router /api/logout [get]
func Logout(c echo.Context) error {
	user := helpers.AuthGetUser(c)
	if user != nil {
		return c.JSON(200, user)
	}

	return c.JSON(401, "Unauthorized")
}

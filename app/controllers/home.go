package controllers

import (
	"template-echo-gorm/app/helpers"
	"template-echo-gorm/app/models"

	"github.com/labstack/echo/v4"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "Welcome to Echo")
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
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
}

func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := helpers.AuthGetUser(c)
		if user != nil {
			return c.JSON(200, user)
		}

		return c.JSON(401, "Unauthorized")
	}
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := helpers.AuthGetUser(c)
		if user != nil {
			return c.JSON(200, user)
		}

		return c.JSON(401, "Unauthorized")
	}
}

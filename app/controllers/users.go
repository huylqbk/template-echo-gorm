package controllers

import (
	"strconv"
	"template-echo-gorm/app/helpers"
	"template-echo-gorm/app/models"

	"github.com/labstack/echo/v4"
)

func UserStore() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}

		if err := c.Bind(&user); err != nil {
			return c.JSON(422, err)
		}

		err := helpers.Validate(&user)

		if err != nil {
			return c.JSON(422, err)
		}

		res := models.UserStore(&user)
		if res {
			return c.JSON(201, user)
		}
		return c.JSON(400, "Bad Request")
	}
}

func UserUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err == nil {
			user := models.UserShow(id)
			if user != nil {
				if err := c.Bind(&user); err != nil {
					return c.JSON(422, err)
				}

				err := helpers.Validate(user)
				if err != nil {
					return c.JSON(422, err)
				}

				res := models.UserUpdate(user)
				if res {
					return c.JSON(200, user)
				}
			} else {
				return c.JSON(404, "Not Found")
			}
		}

		return c.JSON(400, "Bad Request")
	}
}

func UserDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(400, "Bad Request")
		}

		res := models.UserDelete(id)

		if res {
			return c.JSON(200, "Deleted")
		}

		return c.JSON(404, "Not Found")
	}
}

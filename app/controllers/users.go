package controllers

import (
	"github.com/labstack/echo/v4"
	"red-coins/app/helpers"
	"red-coins/app/models"
	"strconv"
)

// UserStore redCoins
// @Summary UserStore
// @Description Create User
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "Username"
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 201 {object} models.User
// @Failure 422 {string} string
// @Failure 400 {string} string
// @Router /api/users [post]
func UserStore(c echo.Context) error {
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

// UserUpdate redCoins
// @Summary UserUpdate
// @Description Update User
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Id"
// @Param username query string true "Username"
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 200 {object} models.User
// @Failure 422 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/users/{id} [put]
func UserUpdate(c echo.Context) error {
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

// UserDelete redCoins
// @Summary UserDelete
// @Description Delete USer
// @Tags User
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Id"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/users/{id} [delete]
func UserDelete(c echo.Context) error {
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

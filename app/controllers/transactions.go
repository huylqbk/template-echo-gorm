package controllers

import (
	"strconv"
	"template-echo-gorm/app/helpers"
	"template-echo-gorm/app/models"
	"template-echo-gorm/config"

	"github.com/labstack/echo/v4"
)

func TransactionStore() echo.HandlerFunc {
	return func(c echo.Context) error {
		transaction := models.Transaction{}

		if err := c.Bind(&transaction); err != nil {
			return c.JSON(422, err)
		}

		value, ok := config.RC.Get("price").Result()

		if ok != nil {
			return c.JSON(422, "error")
		}

		transaction.Cotation, ok = strconv.ParseFloat(value, 64)

		err := helpers.Validate(&transaction)

		if err != nil {
			return c.JSON(422, err)
		}

		res := models.TransactionStore(&transaction)

		if res {
			return c.JSON(201, transaction)
		}

		return c.JSON(400, "Bad Request")
	}
}

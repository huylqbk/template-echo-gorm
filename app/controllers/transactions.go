package controllers

import (
	"github.com/labstack/echo/v4"
	"red-coins/app/helpers"
	"red-coins/app/models"
	"red-coins/config"
	"strconv"
)

// TransactionStore redCoins
// @Summary TransactionStore
// @Description Save Transaction Data
// @Tags Transaction
// @Produce json
// @Security ApiKeyAuth
// @Success 201 {object} models.Transaction
// @Failure 422 {string} string
// @Failure 400 {string} string
// @Router /api/transactions [post]
func TransactionStore(c echo.Context) error {
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
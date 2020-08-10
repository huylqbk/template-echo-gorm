package main

import (
	"red-coins/app/router"
	"red-coins/config"
	"red-coins/migrations"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

// @title redCoins
// @version 1.0
// @description This is a sample server

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app := echo.New()

	db := config.Database()

	migrations.Migrate(db)

	defer db.Close()

	// config.Redis()
	// console.Schedule()
	router.Init(app)

	app.Server.Addr = ":3333"
	graceful.ListenAndServe(app.Server, 5*time.Second)
}

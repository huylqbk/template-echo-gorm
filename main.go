package main

import (
	"template-echo-gorm/app/router"
	"template-echo-gorm/config"
	"template-echo-gorm/migrations"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

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

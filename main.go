package main

import (
	"fmt"
	"template-echo-gorm/app/router"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	app := echo.New()

	// db := config.Database()

	// migrations.Migrate(db)

	// defer db.Close()

	// config.Redis()
	// console.Schedule()
	router.Init(app)

	app.Server.Addr = ":3333"
	fmt.Println("Server started at ", app.Server.Addr)
	graceful.ListenAndServe(app.Server, 5*time.Second)
}

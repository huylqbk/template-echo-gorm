package router

import (
	"red-coins/app/controllers"
	"red-coins/app/middlewares"
	_ "red-coins/docs"

	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(app *echo.Echo) {
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())

	app.GET("/", controllers.Index)
	app.GET("/docs/*", echoSwagger.WrapHandler)

	api := app.Group("/api", middlewares.Jwt())
	// api := app.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.GET("/logout", controllers.Logout)

		users := api.Group("/users")
		{
			users.POST("", controllers.UserStore)
			users.PUT("/:id", controllers.UserUpdate)
			users.DELETE("/:id", controllers.UserDelete)
		}

		transactions := api.Group("/transactions")
		{
			transactions.POST("", controllers.TransactionStore)
		}
	}

	log.Printf("Server started...")
}

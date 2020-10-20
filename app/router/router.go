package router

import (
	"template-echo-gorm/app/controllers"
	"template-echo-gorm/app/errors"
	"template-echo-gorm/app/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	app.HTTPErrorHandler = errors.HttpErrorHandler

	app.GET("/", controllers.Index())

	api := app.Group("/api", middlewares.Jwt())
	{
		api.POST("/login", controllers.Login())
		api.GET("/logout", controllers.Logout())

		users := api.Group("/users")
		{
			users.POST("", controllers.UserStore())
			users.PUT("/:id", controllers.UserUpdate())
			users.DELETE("/:id", controllers.UserDelete())
		}

		transactions := api.Group("/transactions")
		{
			transactions.POST("", controllers.TransactionStore())
		}
	}
}

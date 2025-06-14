package routes

import (
	"skeleton/bootstrap"
	"skeleton/pkg/response"
	"skeleton/src/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "skeleton/docs"
)

func Handle(app *bootstrap.Application) {
	// Register global routes
	app.App.GET("/", func(c echo.Context) error {
		return response.OK(c)
	})
	app.App.GET("/ping", func(c echo.Context) error {
		return response.OK(c, "PONG")
	})

	// show swagger docs
	app.App.GET("/swagger/*", echoSwagger.WrapHandler)

	// handler routes
	apiGroup := app.App.Group("/api/v1")
	handlers := handler.NewHandler(app)

	apiGroup.GET("/example", handlers.Example)
}

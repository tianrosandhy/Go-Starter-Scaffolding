package routes

import (
	"skeleton/bootstrap"
	"skeleton/pkg/response"
	"skeleton/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handle(app *bootstrap.Application) {
	app.App.Use(middleware.CORS())

	// Register global routes
	app.App.GET("/", func(c echo.Context) error {
		return response.OK(c)
	})
	app.App.GET("/ping", func(c echo.Context) error {
		return response.OK(c, "PONG")
	})

	// handler routes
	apiGroup := app.App.Group("/api/v1")
	handlers := handler.NewHandler(app)

	apiGroup.GET("/example", handlers.Example)
}

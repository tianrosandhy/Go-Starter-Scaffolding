package routes

import (
	"skeleton/bootstrap"
	"skeleton/pkg/response"
	"skeleton/src/modules/example"
	"skeleton/src/routes/middleware"

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
	mw := middleware.NewMiddleware(app.Log, app.Config)
	app.App.GET("/swagger/*", echoSwagger.WrapHandler, mw.SwaggerMiddleware)

	// Register module handlers
	example.NewExampleModuleRegistration(app)
}

package docs

import (
	"skeleton/bootstrap"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitSwaggerHost(app *bootstrap.Application) {
	if app.Config.GetString("ENVIRONMENT") != "production" {
		app.App.GET("/swagger/*", echoSwagger.WrapHandler)
		SwaggerInfo.Host = app.Config.GetString("APP_HOST")
	}
}

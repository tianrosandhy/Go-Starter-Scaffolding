package bootstrap

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewEcho(viperConfig *viper.Viper, log *logrus.Logger) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodHead},
	}))

	return e
}

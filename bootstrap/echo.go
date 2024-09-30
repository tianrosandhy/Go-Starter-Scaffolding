package bootstrap

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tianrosandhy/goconfigloader"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewEcho(cfg *goconfigloader.Config, log *logrus.Logger) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodHead},
	}))

	return e
}

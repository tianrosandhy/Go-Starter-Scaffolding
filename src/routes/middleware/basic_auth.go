package middleware

import (
	"encoding/base64"
	"skeleton/pkg/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) BasicAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		correctKey := m.Config.GetString("BASIC_AUTH")
		if len(correctKey) == 0 {
			return next(c)
		}

		// Get the basic auth credentials
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) == 0 {
			return response.ErrorMessage(c, "Unauthorized access", 403)
		}

		splitAuth := strings.Split(authHeader, " ")
		if len(splitAuth) != 2 {
			return response.ErrorMessage(c, "Invalid authorization header", 403)
		}
		authMode := strings.TrimSpace(strings.ToLower(splitAuth[0]))
		if authMode == "bearer" {
			// bearer auth will be checked in another middleware
			return next(c)
		} else if authMode != "basic" {
			return response.ErrorMessage(c, "Invalid authorization type requirement", 403)
		}

		keyData := splitAuth[1]
		decodedByte, err := base64.StdEncoding.DecodeString(keyData)
		if err != nil || correctKey != string(decodedByte) {
			return response.ErrorMessage(c, "Invalid authentication credential", 403)
		}

		// basic key correct or doesnt need to be checked
		return next(c)
	}
}

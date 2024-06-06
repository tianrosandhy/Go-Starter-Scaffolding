package middleware

import (
	"skeleton/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) BearerAuthentication(c *fiber.Ctx) error {
	correctBasicKey := m.Config.GetString("BASIC_AUTH")

	authHeader := c.Get("Authorization")
	if len(authHeader) == 0 && len(correctBasicKey) == 0 {
		// no need to do authentication check if not defined & not needed by config
		return c.Next()
	}

	if len(authHeader) == 0 {
		return response.ErrorMessage(c, "Unauthorized access", 403)
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) != 2 {
		return response.ErrorMessage(c, "Invalid authorization header", 403)
	}
	authMode := strings.TrimSpace(strings.ToLower(splitAuth[0]))
	if authMode == "basic" {
		// basic auth is already checked in another middleware
		return c.Next()
	} else if authMode != "bearer" {
		return response.ErrorMessage(c, "Invalid authorization type requirement", 403)
	}

	bearerToken := splitAuth[1]
	// TODO : validate bearer token
	if len(bearerToken) > 0 {
		return c.Next()
	}

	return response.ErrorMessage(c, "Invalid bearer token", 403)
}

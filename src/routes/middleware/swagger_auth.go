package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// SwaggerMiddleware authenticating before calling next request
func (m *Middleware) SwaggerMiddleware(c *fiber.Ctx) error {
	okstring := m.Config.GetString("SWAGGER_AUTH")
	if len(okstring) == 0 {
		return c.Next()
	}

	authHeader := c.Get("Authorization")
	if len(authHeader) == 0 {
		return sendUnauthorized(c)
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) != 2 {
		return sendUnauthorized(c)
	}

	tipe := splitAuth[0]
	hash := splitAuth[1]

	if strings.ToLower(tipe) != "basic" {
		return sendUnauthorized(c)
	}

	decodedByte, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return sendUnauthorized(c)
	}

	if string(decodedByte) != okstring {
		return sendUnauthorized(c)
	}

	return c.Next()
}

func sendUnauthorized(c *fiber.Ctx) error {
	c.Set("WWW-Authenticate", `Basic realm="mydomain"`)
	return c.SendStatus(401)
}

package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewFiber(viperConfig *viper.Viper, log *logrus.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: viperConfig.GetInt("MAX_BODY_LIMIT") * 1024 * 1024,
		Prefork:   viperConfig.GetBool("PREFORK"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Errorln(err)
			if res, ok := err.(*fiber.Error); ok {
				return c.Status(res.Code).JSON(fiber.Map{
					"type":   "error",
					"errors": res.Message,
				})
			}
			return c.Status(500).JSON(fiber.Map{
				"type":   "error",
				"errors": "Server error",
			})
		},
	})

	return app
}

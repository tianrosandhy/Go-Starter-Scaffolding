package bootstrap

import (
	"skeleton/config"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Application struct {
	Config    *viper.Viper
	Log       *logrus.Logger
	Validator *validator.Validate
	DB        *gorm.DB
	App       *echo.Echo
	Redis     *redis.Client
}

func NewApplication() *Application {
	config := NewViperConfig(config.Environment)
	logger := NewLogger(config)
	validator := NewValidator(config)
	db := NewDatabase(config, logger)
	app := NewEcho(config, logger)
	redis := NewRedis(config)

	application := Application{
		Config:    config,
		Log:       logger,
		Validator: validator,
		DB:        db,
		Redis:     redis,
		App:       app,
	}

	return &application
}

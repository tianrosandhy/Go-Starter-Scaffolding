package handler

import (
	"skeleton/bootstrap"
	"skeleton/src/modules/example/repository"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ExampleHandler struct {
	Validator         *validator.Validate
	Config            *viper.Viper
	Log               *logrus.Logger
	Redis             *redis.Client
	ExampleRepository *repository.ExampleRepository
}

func NewExampleHandler(app *bootstrap.Application, exampleRepository *repository.ExampleRepository) *ExampleHandler {
	return &ExampleHandler{
		Validator:         app.Validator,
		Config:            app.Config,
		Log:               app.Log,
		Redis:             app.Redis,
		ExampleRepository: exampleRepository,
	}
}

package handler

import (
	"skeleton/bootstrap"
	"skeleton/src/modules/product/repository"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ProductHandler struct {
	Validator         *validator.Validate
	Config            *viper.Viper
	Log               *logrus.Logger
	Redis             *redis.Client
	ProductRepository *repository.ProductRepository
}

func NewProductHandler(app *bootstrap.Application, productRepository *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		Validator:         app.Validator,
		Config:            app.Config,
		Log:               app.Log,
		Redis:             app.Redis,
		ProductRepository: productRepository,
	}
}

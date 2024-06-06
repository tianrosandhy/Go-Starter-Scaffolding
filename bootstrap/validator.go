package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func NewValidator(viperConfig *viper.Viper) *validator.Validate {
	return validator.New()
}

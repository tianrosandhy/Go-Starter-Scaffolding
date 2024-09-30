package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/tianrosandhy/goconfigloader"
)

func NewValidator(cfg *goconfigloader.Config) *validator.Validate {
	return validator.New()
}

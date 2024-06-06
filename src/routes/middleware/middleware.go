package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Middleware struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewMiddleware(logger *logrus.Logger, viperConfig *viper.Viper) *Middleware {
	return &Middleware{
		Log:    logger,
		Config: viperConfig,
	}
}

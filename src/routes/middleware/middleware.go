package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/tianrosandhy/goconfigloader"
)

type Middleware struct {
	Log    *logrus.Logger
	Config *goconfigloader.Config
}

func NewMiddleware(logger *logrus.Logger, cfg *goconfigloader.Config) *Middleware {
	return &Middleware{
		Log:    logger,
		Config: cfg,
	}
}

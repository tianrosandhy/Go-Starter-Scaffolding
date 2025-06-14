package handler

import (
	"skeleton/bootstrap"
)

type Handler struct {
	App *bootstrap.Application
}

func NewHandler(app *bootstrap.Application) *Handler {
	return &Handler{
		App: app,
	}
}

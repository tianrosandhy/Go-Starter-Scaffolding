package handler

import (
	"skeleton/bootstrap"
	"skeleton/src/modules/example/repository"
)

type ExampleHandler struct {
	App               *bootstrap.Application
	ExampleRepository *repository.ExampleRepository
}

func NewExampleHandler(app *bootstrap.Application, exampleRepository *repository.ExampleRepository) *ExampleHandler {
	return &ExampleHandler{
		App:               app,
		ExampleRepository: exampleRepository,
	}
}

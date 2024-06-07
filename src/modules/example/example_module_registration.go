package example

import (
	"skeleton/bootstrap"
	"skeleton/src/modules/example/handler"
	"skeleton/src/modules/example/repository"

	"github.com/gofiber/fiber/v2"
)

func NewExampleModuleRegistration(app *bootstrap.Application, route fiber.Router) {
	// register repository & handler
	exampleRepository := repository.NewExampleRepository(app)
	exampleHandler := handler.NewExampleHandler(app, exampleRepository)

	// register routes to handlers
	route.Get("/example", exampleHandler.GetExampleHandler)
	route.Get("/example/:id", exampleHandler.GetSingleExampleHandler)
	route.Post("/example", exampleHandler.PostCreateExample)
	route.Patch("/example/:id", exampleHandler.PatchUpdateExample)
	route.Delete("/example/:id", exampleHandler.DeleteSingleExample)
}

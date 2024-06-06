package product

import (
	"skeleton/bootstrap"
	"skeleton/src/modules/product/handler"
	"skeleton/src/modules/product/repository"

	"github.com/gofiber/fiber/v2"
)

func NewProductModuleRegistration(app *bootstrap.Application, route fiber.Router) {
	// register repository & handler
	productRepository := repository.NewProductRepository(app)
	productHandler := handler.NewProductHandler(app, productRepository)

	// register routes to handlers
	route.Get("/product", productHandler.GetProductHandler)
	route.Get("/product/:id", productHandler.GetSingleProductHandler)
	route.Post("/product", productHandler.PostCreateProduct)
	route.Patch("/product/:id", productHandler.PatchUpdateProduct)
	route.Delete("/product/:id", productHandler.DeleteSingleProduct)
}

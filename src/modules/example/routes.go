package example

import (
	"skeleton/bootstrap"
	"skeleton/src/modules/example/handler"
	"skeleton/src/modules/example/repository"
	"skeleton/src/routes/middleware"
)

func NewExampleModuleRegistration(app *bootstrap.Application) {
	// register repository & handler
	exampleRepository := repository.NewExampleRepository(app)
	exampleHandler := handler.NewExampleHandler(app, exampleRepository)

	mid := middleware.NewMiddleware(app.Log, app.Config)
	group := app.App.Group("/api/example", mid.BasicAuthentication)

	// register routes to handlers
	group.GET("/v1/lists", exampleHandler.ListsExample)
	group.GET("/v1/detail/:id", exampleHandler.DetailExample)
	group.POST("/v1/store", exampleHandler.StoreExample)
	group.POST("/v1/update/:id", exampleHandler.UpdateExample)
	group.POST("/v1/delete/:id", exampleHandler.DeleteExample)
}

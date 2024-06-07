package routes

import (
	"skeleton/bootstrap"
	"skeleton/pkg/recovery"
	"skeleton/pkg/response"
	"skeleton/pkg/telelogger"
	"skeleton/src/modules/example"
	"skeleton/src/routes/middleware"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "skeleton/docs"
)

func Handle(app *bootstrap.Application) {
	api := RouteBootstrap(app)

	// Register global routes
	api.Get("/", func(c *fiber.Ctx) error {
		return response.OK(c)
	})
	api.Get("/ping", func(c *fiber.Ctx) error {
		return response.OK(c, "PONG")
	})

	// Register module handlers
	example.NewExampleModuleRegistration(app, api)
}

func RouteBootstrap(app *bootstrap.Application) fiber.Router {
	// grouping routes
	api := app.App.Group(app.Config.GetString("ENDPOINT"))

	// app middleware
	middleware := middleware.NewMiddleware(app.Log, app.Config)
	tlog := telelogger.NewTeleLogger(app.Log, app.Config)

	api.Use(cors.New())
	api.Use(limiter.New(limiter.Config{
		Max:        app.Config.GetInt("LIMITER_MAX_HIT"),
		Expiration: time.Duration(app.Config.GetInt("LIMITER_DURATION")) * time.Second,
	}))
	api.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			err := recovery.StackTrace(e)
			tlog.PushError(err)
		},
	}))

	// basic & bearer auth
	api.Use(middleware.BasicAuthentication)
	api.Use(middleware.BearerAuthentication)

	// swagger
	swagApi := api.Group("/swag", middleware.SwaggerMiddleware)
	swagApi.Get("/docs/*", swagger.HandlerDefault)

	return api
}

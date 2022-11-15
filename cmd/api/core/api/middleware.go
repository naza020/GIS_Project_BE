package api

import (
	"gis-project-backend/pkg/core"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	reqId "github.com/gofiber/fiber/v2/middleware/requestid"
)

func RegisterMiddleware(app *fiber.App, coreRegistry *core.CoreRegistry) {

	app.Use(reqId.New())

	app.Use(NewLogger())
	app.Use(NewRecover())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

}

func NewLogger() fiber.Handler {
	return logger.New()
}

func NewLogRequest(coreRegistry *core.CoreRegistry) fiber.Handler {

	return func(c *fiber.Ctx) (err error) {
		chainErr := c.Next()

		return chainErr
	}
}

func NewRecover() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
	})
}

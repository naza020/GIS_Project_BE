package api

import (
	"github.com/gofiber/fiber/v2"
)

func ServerConfig() fiber.Config {
	return fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.ErrInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e
			}
			return ResponseErrorWithCode(c, code, err)
		},
	}
}

package adapter

import (
	"github.com/gofiber/fiber/v2"
)

type AuthorAdapter interface {
	GetHandler() HttpHandler
}

type HttpHandler interface {
	Store(c *fiber.Ctx) error
}

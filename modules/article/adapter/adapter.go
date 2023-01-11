package adapter

import "github.com/gofiber/fiber/v2"

type ArticleAdapter interface {
	InitializeRoute()
	GetHttpHandler() HttpHandler
}

type HttpHandler interface {
	FindAll(c *fiber.Ctx) error
}
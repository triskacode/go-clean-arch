package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/delivery"
)

type authorModule struct {
	handler adapter.HttpHandler
}

func NewModule(app *fiber.App) (mod *authorModule) {
	mod = new(authorModule)

	mod.handler = delivery.NewHttpHandler(app)
	return
}

func (m authorModule) GetHandler() adapter.HttpHandler {
	return m.handler
}

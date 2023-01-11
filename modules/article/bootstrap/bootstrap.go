package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
	"github.com/triskacode/go-clean-arch/modules/article/delivery"
)

type articleAdapter struct {
	app         *fiber.App
	httpHandler adapter.HttpHandler
}

type ModuleDeps struct {
	App *fiber.App
}

func NewModule(deps ModuleDeps) (m *articleAdapter) {
	m = new(articleAdapter)
	m.app = deps.App
	m.httpHandler = delivery.NewHttpHandler()

	return
}

func (m *articleAdapter) InitializeRoute() {
	m.app.Get("/article", m.GetHttpHandler().FindAll)
}

func (m *articleAdapter) GetHttpHandler() adapter.HttpHandler {
	return m.httpHandler
}

package bootstrap

import "github.com/gofiber/fiber/v2"

type articleAdapter struct {
	app *fiber.App
}

type ModuleDeps struct {
	App *fiber.App
}

func NewModule(deps ModuleDeps) (m *articleAdapter) {
	m = new(articleAdapter)
	m.app = deps.App

	return
}

func (m *articleAdapter) InitializeRoute() {

}

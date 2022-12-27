package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/delivery"
)

type AuthorModule struct {
	handler HttpHandler
}

func NewModule(app *fiber.App) (mod *AuthorModule) {
	mod = new(AuthorModule)

	mod.handler = delivery.NewHttpHandler(app)
	return
}

func (m AuthorModule) GetHandler() HttpHandler {
	return m.handler
}

package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/delivery"
	"github.com/triskacode/go-clean-arch/modules/author/validation"
)

type authorModule struct {
	handler   adapter.HttpHandler
	validator adapter.AuthorValidator
}

func NewModule(app *fiber.App) (mod *authorModule) {
	mod = new(authorModule)

	mod.validator = validation.NewAuthorValidator()
	mod.handler = delivery.NewHttpHandler(app, mod.GetValidator())
	return
}

func (m authorModule) GetValidator() adapter.AuthorValidator {
	return m.validator
}

func (m authorModule) GetHandler() adapter.HttpHandler {
	return m.handler
}

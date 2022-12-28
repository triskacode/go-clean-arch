package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type httpHandler struct {
	validator adapter.AuthorValidator
}

func NewHttpHandler(app *fiber.App, validator adapter.AuthorValidator) (h *httpHandler) {
	h = new(httpHandler)
	h.validator = validator

	app.Post("/author", h.Store)
	return
}

func (h httpHandler) Store(c *fiber.Ctx) error {
	dto := new(dto.CreateAuthorDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := h.validator.ValidateCreateAuthorDto(*dto); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.SendString("pong")
}

package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
	"github.com/triskacode/go-clean-arch/modules/author/validation"
)

type HttpHandler struct{}

func NewHttpHandler(app *fiber.App) (h *HttpHandler) {
	h = new(HttpHandler)

	app.Post("/author", h.Store)
	return
}

func (h *HttpHandler) Store(c *fiber.Ctx) error {
	dto := new(dto.CreateAuthorDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := validation.ValidateCreateAuthorDto(*dto); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.SendString("pong")
}

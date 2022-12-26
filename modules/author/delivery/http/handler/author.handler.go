package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/delivery/http/validation"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorHandler struct{}

func New(app *fiber.App) {
	h := new(AuthorHandler)

	app.Post("/author", h.Store)
}

func (h *AuthorHandler) Store(c *fiber.Ctx) error {
	dto := new(dto.CreateAuthorDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validation.ValidateCreateAuthorDto(*dto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.SendString("pong")
}

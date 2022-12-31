package delivery

import (
	"github.com/gofiber/fiber/v2"
	httpAdapter "github.com/triskacode/go-clean-arch/adapter/http"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
	"github.com/triskacode/go-clean-arch/modules/author/validation"
)

type httpHandler struct {
	validator validation.AuthorValidator
}

func NewHttpHandler(app *fiber.App) (h *httpHandler) {
	h = new(httpHandler)
	h.validator = validation.NewAuthorValidator()

	app.Post("/author", h.Store)
	return
}

func (h httpHandler) Store(c *fiber.Ctx) error {
	dto := new(dto.CreateAuthorDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(httpAdapter.ErrorRespModel{
			Code:    fiber.StatusBadRequest,
			Message: "BAD_REQUEST",
			Errors:  err.Error(),
		})
	}

	if err := h.validator.ValidateCreateAuthorDto(*dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(httpAdapter.ErrorRespModel{
			Code:    fiber.StatusBadRequest,
			Message: "BAD_REQUEST",
			Errors:  err,
		})
	}

	return c.JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
	})
}

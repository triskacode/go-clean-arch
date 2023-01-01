package delivery

import (
	"github.com/gofiber/fiber/v2"
	httpAdapter "github.com/triskacode/go-clean-arch/adapter/http"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
	"github.com/triskacode/go-clean-arch/modules/author/validation"
)

type httpHandler struct {
	validator     validation.AuthorValidator
	authorUsecase adapter.AuthorUsecase
}

func NewHttpHandler(authorUsecase adapter.AuthorUsecase) (h *httpHandler) {
	h = new(httpHandler)

	h.validator = validation.NewAuthorValidator()
	h.authorUsecase = authorUsecase
	return
}

func (h httpHandler) Create(c *fiber.Ctx) error {
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

	author := h.authorUsecase.Create(*dto)

	return c.JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    author,
	})
}

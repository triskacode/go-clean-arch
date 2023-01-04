package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/exception"
	httpAdapter "github.com/triskacode/go-clean-arch/infrastructure/http"
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
		return exception.NewBadRequestException(err.Error())
	}

	if err := h.validator.ValidateCreateAuthorDto(*dto); err != nil {
		return exception.NewBadRequestException(err)
	}

	author, err := h.authorUsecase.Create(*dto)
	if err != nil {
		return err
	}

	return c.JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    author,
	})
}

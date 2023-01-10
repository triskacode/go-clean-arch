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

func (h *httpHandler) FindAll(c *fiber.Ctx) error {
	authors, err := h.authorUsecase.FindAll()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    authors,
	})
}

func (h *httpHandler) Create(c *fiber.Ctx) error {
	f := new(dto.CreateAuthorDto)
	if err := c.BodyParser(f); err != nil {
		return exception.NewBadRequestException(nil)
	}

	if err := h.validator.ValidateCreateAuthorDto(*f); err != nil {
		return exception.NewBadRequestException(err)
	}

	author, err := h.authorUsecase.Create(*f)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusCreated,
		Message: "CREATED",
		Data:    author,
	})
}

func (h *httpHandler) FindById(c *fiber.Ctx) error {
	p := new(dto.ParamIdDto)
	if err := c.ParamsParser(p); err != nil {
		return exception.NewNotFoundException(nil)
	}

	author, err := h.authorUsecase.FindById(*p)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    author,
	})
}

func (h *httpHandler) Update(c *fiber.Ctx) error {
	p := new(dto.ParamIdDto)
	if err := c.ParamsParser(p); err != nil {
		return exception.NewNotFoundException(nil)
	}

	f := new(dto.UpdateAuthorDto)
	if err := c.BodyParser(f); err != nil {
		return exception.NewBadRequestException(nil)
	}

	if err := h.validator.ValidateUpdateAuthorDto(*f); err != nil {
		return exception.NewBadRequestException(err)
	}

	author, err := h.authorUsecase.Update(*p, *f)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    author,
	})
}

func (h *httpHandler) Delete(c *fiber.Ctx) error {
	p := new(dto.ParamIdDto)
	if err := c.ParamsParser(p); err != nil {
		return exception.NewNotFoundException(nil)
	}

	author, err := h.authorUsecase.Delete(*p)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(httpAdapter.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    author,
	})
}

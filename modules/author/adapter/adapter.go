package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorAdapter interface {
	GetHandler() HttpHandler
	GetValidator() AuthorValidator
}

type HttpHandler interface {
	Store(c *fiber.Ctx) error
}

type AuthorValidator interface {
	ValidateCreateAuthorDto(dto dto.CreateAuthorDto) (model []exception.ErrorValidationModel)
}

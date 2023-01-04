package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/author/dto"
)

type AuthorAdapter interface {
	InitializeRoute()
	GetHttpHandler() HttpHandler
	GetAuthorUsecase() AuthorUsecase
	GetAuthorRepository() AuthorRepository
}

type HttpHandler interface {
	Create(c *fiber.Ctx) error
}

type AuthorUsecase interface {
	Create(f dto.CreateAuthorDto) (*dto.AuthorResponseDto, *exception.HttpException)
}

type AuthorRepository interface {
	Create(author *domain.Author) error
}

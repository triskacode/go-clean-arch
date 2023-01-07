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
	FindAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}

type AuthorUsecase interface {
	FindAll() (r *[]dto.AuthorResponseDto, e *exception.HttpException)
	Create(f dto.CreateAuthorDto) (r *dto.AuthorResponseDto, e *exception.HttpException)
	FindById(p dto.ParamIdDto) (r *dto.AuthorResponseDto, e *exception.HttpException)
}

type AuthorRepository interface {
	FindAll(authors *[]domain.Author) error
	Create(author *domain.Author) error
	FindById(author *domain.Author) error
}

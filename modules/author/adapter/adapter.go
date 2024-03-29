package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
	"github.com/triskacode/go-clean-arch/exception"
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
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type AuthorUsecase interface {
	FindAll() (r *[]dto.AuthorResponseDto, e *exception.HttpException)
	Create(f dto.CreateAuthorDto) (r *dto.AuthorResponseDto, e *exception.HttpException)
	FindById(p dto.ParamIdDto) (r *dto.AuthorResponseDto, e *exception.HttpException)
	Update(p dto.ParamIdDto, f dto.UpdateAuthorDto) (r *dto.AuthorResponseDto, e *exception.HttpException)
	Delete(p dto.ParamIdDto) *exception.HttpException
}

type AuthorRepository interface {
	FindAll(authors *[]*entity.Author) error
	Create(author *entity.Author) error
	FindOne(author *entity.Author) error
	Update(author *entity.Author, f dto.UpdateAuthorDto) error
	Delete(author *entity.Author) error
}

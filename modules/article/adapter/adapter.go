package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/domain/entity"
	"github.com/triskacode/go-clean-arch/exception"
)

type ArticleAdapter interface {
	InitializeRoute()
	GetHttpHandler() HttpHandler
	GetArticleUsecase() ArticleUsecase
	GetArticleRepository() ArticleRepository
}

type HttpHandler interface {
	FindAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}

type ArticleUsecase interface {
	FindAll() (r *[]dto.ArticleResponseDto, e *exception.HttpException)
	Create(f dto.CreateArticleDto) (r *dto.ArticleResponseDto, e *exception.HttpException)
	FindById(p dto.ParamIdDto) (r *dto.ArticleResponseDto, e *exception.HttpException)
}

type ArticleRepository interface {
	FindAll(articles *[]*entity.Article) error
	Create(article *entity.Article) error
	FindOne(article *entity.Article) error
}

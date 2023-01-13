package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/modules/article/dto"
)

type ArticleAdapter interface {
	InitializeRoute()
	GetHttpHandler() HttpHandler
	GetArticleUsecase() ArticleUsecase
	GetArticleRepository() ArticleRepository
}

type HttpHandler interface {
	FindAll(c *fiber.Ctx) error
}

type ArticleUsecase interface {
	FindAll() (r *[]dto.ArticleResponseDto, e *exception.HttpException)
}

type ArticleRepository interface {
	FindAll(articles *[]domain.Article) error
}

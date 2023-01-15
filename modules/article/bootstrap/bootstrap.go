package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
	"github.com/triskacode/go-clean-arch/modules/article/delivery"
	"github.com/triskacode/go-clean-arch/modules/article/repository"
	"github.com/triskacode/go-clean-arch/modules/article/usecase"
	author "github.com/triskacode/go-clean-arch/modules/author/adapter"
	"gorm.io/gorm"
)

type articleAdapter struct {
	app               *fiber.App
	httpHandler       adapter.HttpHandler
	articleUsecase    adapter.ArticleUsecase
	articleRepository adapter.ArticleRepository
}

type ModuleDeps struct {
	App              *fiber.App
	DB               *gorm.DB
	AuthorRepository author.AuthorRepository
}

func NewModule(deps ModuleDeps) (m *articleAdapter) {
	m = new(articleAdapter)
	m.app = deps.App
	m.articleRepository = repository.NewArticleRepository(deps.DB)
	m.articleUsecase = usecase.NewArticleUsecase(m.GetArticleRepository(), deps.AuthorRepository)
	m.httpHandler = delivery.NewHttpHandler(m.GetArticleUsecase())

	return
}

func (m *articleAdapter) InitializeRoute() {
	m.app.Get("/article", m.GetHttpHandler().FindAll)
	m.app.Post("/article", m.GetHttpHandler().Create)
	m.app.Get("/article/:id", m.GetHttpHandler().FindById)
}

func (m *articleAdapter) GetHttpHandler() adapter.HttpHandler {
	return m.httpHandler
}

func (m *articleAdapter) GetArticleUsecase() adapter.ArticleUsecase {
	return m.articleUsecase
}

func (m *articleAdapter) GetArticleRepository() adapter.ArticleRepository {
	return m.articleRepository
}

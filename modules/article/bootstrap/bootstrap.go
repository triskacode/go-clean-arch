package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
	"github.com/triskacode/go-clean-arch/modules/article/delivery"
	"github.com/triskacode/go-clean-arch/modules/article/repository"
	"github.com/triskacode/go-clean-arch/modules/article/usecase"
	"gorm.io/gorm"
)

type articleAdapter struct {
	app               *fiber.App
	httpHandler       adapter.HttpHandler
	articleUsecase    adapter.ArticleUsecase
	articleRepository adapter.ArticleRepository
}

type ModuleDeps struct {
	App *fiber.App
	DB  *gorm.DB
}

func NewModule(deps ModuleDeps) (m *articleAdapter) {
	m = new(articleAdapter)
	m.app = deps.App
	m.articleRepository = repository.NewArticleRepository(deps.DB)
	m.articleUsecase = usecase.NewArticleUsecase(m.GetArticleRepository())
	m.httpHandler = delivery.NewHttpHandler(m.GetArticleUsecase())

	return
}

func (m *articleAdapter) InitializeRoute() {
	m.app.Get("/article", m.GetHttpHandler().FindAll)
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

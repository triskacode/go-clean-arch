package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/modules/author/adapter"
	"github.com/triskacode/go-clean-arch/modules/author/delivery"
	"github.com/triskacode/go-clean-arch/modules/author/repository"
	"github.com/triskacode/go-clean-arch/modules/author/usecase"
	"gorm.io/gorm"
)

type authorModule struct {
	app              *fiber.App
	httpHandler      adapter.HttpHandler
	authorUsecase    adapter.AuthorUsecase
	authorRepository adapter.AuthorRepository
}

type ModuleDeps struct {
	App *fiber.App
	DB  *gorm.DB
}

func NewModule(deps ModuleDeps) (m *authorModule) {
	m = new(authorModule)
	m.app = deps.App
	m.authorRepository = repository.NewAuthorRepository(deps.DB)
	m.authorUsecase = usecase.NewAuthorUsecase(m.GetAuthorRepository())
	m.httpHandler = delivery.NewHttpHandler(m.GetAuthorUsecase())

	return
}

func (m *authorModule) InitializeRoute() {
	m.app.Get("/author", m.GetHttpHandler().FindAll)
	m.app.Post("/author", m.GetHttpHandler().Create)
	m.app.Get("/author/:id", m.GetHttpHandler().FindById)
	m.app.Patch("/author/:id", m.GetHttpHandler().Update)
	m.app.Delete("/author/:id", m.GetHttpHandler().Delete)
}

func (m *authorModule) GetHttpHandler() adapter.HttpHandler {
	return m.httpHandler
}

func (m *authorModule) GetAuthorUsecase() adapter.AuthorUsecase {
	return m.authorUsecase
}

func (m *authorModule) GetAuthorRepository() adapter.AuthorRepository {
	return m.authorRepository
}

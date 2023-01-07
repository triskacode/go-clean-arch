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

func NewModule(app *fiber.App, conn *gorm.DB) (m *authorModule) {
	m = new(authorModule)
	m.app = app
	m.authorRepository = repository.NewAuthorRepository(conn)
	m.authorUsecase = usecase.NewAuthorUsecase(m.GetAuthorRepository())
	m.httpHandler = delivery.NewHttpHandler(m.GetAuthorUsecase())

	return
}

func (m *authorModule) InitializeRoute() {
	m.app.Get("/author", m.GetHttpHandler().FindAll)
	m.app.Post("/author", m.GetHttpHandler().Create)
	m.app.Get("/author/:id", m.GetHttpHandler().FindById)
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

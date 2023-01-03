package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/triskacode/go-clean-arch/config"
)

type httpService struct {
	app    *fiber.App
	config *config.Config
}

func NewHttpService(cfg *config.Config) (httpSvc *httpService) {
	httpSvc = new(httpService)
	httpSvc.config = cfg
	httpSvc.app = fiber.New(fiber.Config{
		ErrorHandler: ExceptionHandler,
	})

	httpSvc.initializeMiddleware()
	httpSvc.app.Get("/ping", HealthCheckHandler)

	return
}

func (httpSvc httpService) initializeMiddleware() {
	httpSvc.app.Use(cors.New())
	httpSvc.app.Use(recover.New())
	httpSvc.app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
}

func (httpSvc httpService) GetApp() *fiber.App {
	return httpSvc.app
}

func (httpSvc httpService) Run() {
	port := fmt.Sprintf(":%d", httpSvc.config.App.Port)
	log.Fatal(httpSvc.app.Listen(port))
}

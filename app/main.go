package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	dbAdapter "github.com/triskacode/go-clean-arch/adapter/database"
	"github.com/triskacode/go-clean-arch/config"
	"github.com/triskacode/go-clean-arch/domain"
	author "github.com/triskacode/go-clean-arch/modules/author/bootstrap"
)

func main() {
	cfg := config.New()
	dbs := dbAdapter.NewConnection(cfg)
	defer dbs.CloseConnection()

	dbs.Migrate(&domain.Article{}, &domain.Author{})

	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	author.NewModule(app)

	port := fmt.Sprintf(":%d", cfg.App.Port)
	log.Fatal(app.Listen(port))
}

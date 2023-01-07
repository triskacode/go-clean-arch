package http

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/exception"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func ExceptionHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("\033[31m[ERROR]\033[0m ")
	log.Println(err)

	var e *exception.HttpException
	if errors.As(err, &e) {
		return c.Status(int(e.Code)).JSON(ErrorRespModel{
			Code:    e.Code,
			Message: e.Message,
			Errors:  e.Errors,
		})

	}

	return c.Status(fiber.StatusInternalServerError).JSON(ErrorRespModel{
		Code:    fiber.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
	})
}

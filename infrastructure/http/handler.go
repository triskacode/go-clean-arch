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

func ExceptionHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(err)

	var e *exception.HttpException
	if errors.As(err, &e) {
		return ctx.Status(int(e.Code)).JSON(ErrorRespModel{
			Code:    e.Code,
			Message: e.Message,
			Errors:  e.Errors,
		})

	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorRespModel{
		Code:    fiber.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
	})
}

package http

import (
	"errors"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/triskacode/go-clean-arch/exception"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func ExceptionHandler(c *fiber.Ctx, err error) error {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("\033[31m[ERROR]\033[0m ")
	log.Println(err)

	resp := new(ErrorRespModel)
	if e := new(exception.HttpException); errors.As(err, &e) {
		resp.Code = e.Code
		resp.Message = e.Message
		resp.Errors = e.Detail
	} else if e := new(fiber.Error); errors.As(err, &e) {
		message := utils.StatusMessage(e.Code)
		message = strings.ToUpper(message)
		message = strings.ReplaceAll(message, " ", "_")

		resp.Code = e.Code
		resp.Message = message
		if e.Message != utils.StatusMessage(e.Code) {
			resp.Errors = e.Message
		}
	} else {
		resp.Code = fiber.StatusInternalServerError
		resp.Message = "INTERNAL_SERVER_ERROR"
	}

	return c.Status(resp.Code).JSON(resp)
}

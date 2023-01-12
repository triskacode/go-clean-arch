package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/infrastructure/http"
)

type httpHandler struct {
}

func NewHttpHandler() (h *httpHandler) {
	h = new(httpHandler)

	return
}

func (h *httpHandler) FindAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(http.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
	})
}

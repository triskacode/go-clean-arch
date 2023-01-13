package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/infrastructure/http"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
)

type httpHandler struct {
	articleUsecase adapter.ArticleUsecase
}

func NewHttpHandler(articleUsecase adapter.ArticleUsecase) (h *httpHandler) {
	h = new(httpHandler)
	h.articleUsecase = articleUsecase

	return
}

func (h *httpHandler) FindAll(c *fiber.Ctx) error {
	articles, err := h.articleUsecase.FindAll()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(http.SuccessRespModel{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    articles,
	})
}

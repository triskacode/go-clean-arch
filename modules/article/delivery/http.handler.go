package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triskacode/go-clean-arch/domain/dto"
	"github.com/triskacode/go-clean-arch/exception"
	"github.com/triskacode/go-clean-arch/helper/validation"
	"github.com/triskacode/go-clean-arch/infrastructure/http"
	"github.com/triskacode/go-clean-arch/modules/article/adapter"
)

type httpHandler struct {
	validator      validation.ArticleValidator
	articleUsecase adapter.ArticleUsecase
}

func NewHttpHandler(articleUsecase adapter.ArticleUsecase) (h *httpHandler) {
	h = new(httpHandler)
	h.validator = validation.NewArticleValidator()
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

func (h *httpHandler) Create(c *fiber.Ctx) error {
	f := new(dto.CreateArticleDto)
	if err := c.BodyParser(f); err != nil {
		return exception.NewBadRequestException(nil)
	}

	if err := h.validator.ValidateCreateArticleDto(*f); err != nil {
		return exception.NewBadRequestException(err)
	}

	article, err := h.articleUsecase.Create(*f)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(http.SuccessRespModel{
		Code:    fiber.StatusCreated,
		Message: "CREATED",
		Data:    article,
	})
}

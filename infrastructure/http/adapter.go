package http

import "github.com/gofiber/fiber/v2"

type HttpAdapter interface {
	initializeMiddleware()
	GetApp() *fiber.App
	Run()
}

type SuccessRespModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorRespModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

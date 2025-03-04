package handlers

import "github.com/gofiber/fiber/v2"

type APIHandler struct{}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

func (h *APIHandler) SetupRoutes(router fiber.Router) {
	router.Get("/", h.HelloWorld)
}

func (h *APIHandler) HelloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

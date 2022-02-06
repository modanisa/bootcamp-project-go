package hello

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/hello", h.Hello)
}

func (h *Handler) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

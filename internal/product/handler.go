package product

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetProducts(ctx context.Context) ([]Product, error)
}

type Handler struct {
	svc Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		svc: service,
	}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/products", h.GetProducts)
}

func (h *Handler) GetProducts(ctx *fiber.Ctx) error {
	products, err := h.svc.GetProducts(ctx.Context())
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(products)
}

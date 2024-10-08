package internal

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) NotAllowed(ctx fiber.Ctx) error {
	return ctx.SendString("you don't belog here")
}

func (h *Handler) GetAll(ctx fiber.Ctx) error {
	return ctx.SendString("get all")
}

func (h *Handler) GetByID(ctx fiber.Ctx) error {
	return ctx.SendString(ctx.Params("event_id"))
}

func (h *Handler) GetByFilters(ctx fiber.Ctx) error {
	params := ctx.Queries()

	jj, _ := json.MarshalIndent(params, "", "  ")
	return ctx.SendString(string(jj))
}

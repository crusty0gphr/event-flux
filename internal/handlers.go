package internal

import (
	"encoding/json"
	"fmt"

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
	events, err := h.service.GetAll(ctx.Context())
	if err != nil {
		return ctx.SendString(
			fmt.Sprintf("handler failed: %s", err.Error()),
		)
	}

	bytes, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return ctx.SendString(
			fmt.Sprintf("handler failed: %s", err.Error()),
		)
	}
	return ctx.SendString(string(bytes))
}

func (h *Handler) GetByID(ctx fiber.Ctx) error {
	eventID := ctx.Params("event_id")
	event, err := h.service.GetByID(ctx.Context(), eventID)
	if err != nil {
		return ctx.SendString(
			fmt.Sprintf("handler failed: %s", err.Error()),
		)
	}

	bytes, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return ctx.SendString(
			fmt.Sprintf("handler failed: %s", err.Error()),
		)
	}
	return ctx.SendString(string(bytes))
}

func (h *Handler) GetByFilters(ctx fiber.Ctx) error {
	params := ctx.Queries()
	events, err := h.service.GetByFilter(ctx.Context(), params)
	if err != nil {
		return ctx.SendString(
			fmt.Sprintf("handler failed: %s", err.Error()),
		)
	}

	bytes, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return ctx.SendString(
			fmt.Sprintf("handler failed: %s", err.Error()),
		)
	}
	return ctx.SendString(string(bytes))
}

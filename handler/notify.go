package handler

import (
	"net/http"

	"github.com/famesensor/playground-go-telegram-bot/service/notify"
	"github.com/gofiber/fiber/v2"
)

type NotifyHandler struct {
	service notify.Port
}

func NewNotifyHandler(service notify.Port) *NotifyHandler {
	return &NotifyHandler{
		service: service,
	}
}

type notifyRequest struct {
	Message string `json:"message"`
}

func (h *NotifyHandler) Notify(c *fiber.Ctx) error {
	body := new(notifyRequest)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := h.service.Send(body.Message); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "success"})
}

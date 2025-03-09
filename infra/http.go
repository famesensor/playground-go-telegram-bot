package infra

import (
	"github.com/famesensor/playground-go-telegram-bot/config"
	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	App *fiber.App
}

func NewHTTPServer(cfg *config.Config) *HTTPServer {
	app := fiber.New()

	return &HTTPServer{
		App: app,
	}
}

func (h *HTTPServer) Start() {
	if err := h.App.Listen(":3000"); err != nil {
		panic(err)
	}
}

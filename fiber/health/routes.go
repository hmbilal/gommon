package health

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(router fiber.Router) {
	ping := router.Group("/health")
	ping.Get("", h.Ping)
	ping.Get("/details", h.Details)
}

package http

import (
	"github.com/sweetie-pie/line-recommendation/internal/port/mysql"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository *mysql.MySQL
}

func (h *Handler) Data(c *fiber.Ctx) error {
	return c.JSON(&Response{})
}

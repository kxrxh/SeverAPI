package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func monitorRoute(c *fiber.Ctx) error {
	key := c.Params("key", "")
	if key == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request. Key is required",
		})
	}
	if key != apiCore.devkey {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return monitor.New(monitor.Config{Title: "Sever Hack Metrics"})(c)
}

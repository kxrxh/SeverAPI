package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func aunthenticate(c *fiber.Ctx) error {
	user := c.FormValue("username", "null")
	password := c.FormValue("password", "null")

	if user == "null" || password == "null" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request. Username and password are required",
		})
	}

	return nil

}

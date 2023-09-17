package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// addConductorRoutes adds the conductor routes to the router.
//
// router: The fiber router to add the routes to.
//
// No return value.
func addConductorRoutes(router fiber.Router) {
	router.Get("/client", getClientInfoHandler)
}

type clientInfoBody struct {
	ClientCardUID string `json:"card_uid"`
}

// getClientInfoHandler handles the request to get client information.
//
// It expects a pointer to a fiber.Ctx object as a parameter.
// It returns an error.
func getClientInfoHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	log.Println(claims["exp"].(time.Time))
	body := clientInfoBody{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"client": claims["conductor_id"].(string),
	})
}

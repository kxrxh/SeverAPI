package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"sever.hack/api/database/controllers"
)

// addConductorRoutes adds the conductor routes to the router.
//
// router: The fiber router to add the routes to.
//
// No return value.
func addConductorRoutes(router fiber.Router) {
	router.Get("/client", getClientInfoHandler)
	router.Get("/dump", getUserDatabaseDump)
}

type clientInfoBody struct {
	ClientCardUID string `json:"card_uid"`
}

// !THIS IS JWT VERSION OF BELOW FUNCTION!
// getClientInfoHandler handles the request to get client information.
//
// It expects a pointer to a fiber.Ctx object as a parameter.
// It returns an error.
// func getClientInfoHandler(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	log.Println(claims["exp"].(time.Time))
// 	body := clientInfoBody{}
// 	if err := c.BodyParser(&body); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"client": claims["conductor_id"].(string),
// 	})
// }

func getClientInfoHandler(c *fiber.Ctx) error {
	var body clientInfoBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user, ok := controllers.GetClient(body.ClientCardUID)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to get client info",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Ok",
		"client":  user,
	})
}

func getUserDatabaseDump(c *fiber.Ctx) error {
	dumps, ok := controllers.GetUserDump()
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unable to get user dump",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Ok",
		"dumps":   dumps,
	})
}

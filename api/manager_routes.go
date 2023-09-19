package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"sever.hack/api/database/controllers"
	"sever.hack/api/database/models"
)

func addManagerRoutes(r fiber.Router) {
	// r.Post("/client/new", addClientRoute)
	r.Get("/client/db", getAllClientRoute)
	r.Patch("/client/", updateUserRoute)
}

type AddClientBody struct {
	Client clientData `json:"client"`
	Card   clientCard `json:"card"`
}
type clientData struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	Birth      string `json:"birth"`
	Sex        string `json:"sex"`
	Email      string `json:"email"`
	CityID     uint   `json:"city_id"`
	StateID    uint   `json:"state_id"`
	CardID     uint   `json:"card_id"`
	Snils      string `json:"snils"`
}

type clientCard struct {
	UID       string    `json:"uid"`
	TimeOpen  time.Time `json:"time_open"`
	TimeClose time.Time `json:"time_close"`
	StateID   uint      `json:"state_id"`
	ClientID  uint      `json:"client_id"`
}

// func addClientRoute(c *fiber.Ctx) error {
// 	var body addClientBody
// 	if err := c.BodyParser(&body); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	err := controllers.AddClient(addClientBody)
// 	return
// }

func getAllClientRoute(c *fiber.Ctx) error {
	res := controllers.GetAllDatabase()
	if len(res) == 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unable to get clients",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Ok",
		"clients": res,
	})
}

func updateUserRoute(c *fiber.Ctx) error {
	var body models.Client
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err := controllers.UpdateClient(&body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Ok",
	})
}

package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"sever.hack/api/database/controllers"
)

/*

Request body:
{
	"conductor_id": "...",
	"device_uid": "...",
	"conductor_pass": "..."
}

*/

type conductorLoginBody struct {
	ConductorID   string `json:"conductor_id"`
	DeviceUID     string `json:"device_uid"`
	ConductorPass string `json:"conductor_pass"`
}

// conductorLogin handles the login request for a conductor.
//
// It parses the request body, checks if the conductor exists, creates a JWT token,
// and returns the token if the login is successful.
//
// Parameters:
// - c: a pointer to the fiber.Ctx object representing the context of the HTTP request.
//
// Returns:
// - error: an error object if there is an error processing the request.
func conductorLogin(c *fiber.Ctx) error {
	var req conductorLoginBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	conductor := controllers.GetConductor(req.ConductorID, req.ConductorPass)
	if conductor == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	claims := jwt.MapClaims{
		"id":           conductor.ID,
		"conductor_id": req.ConductorID,
		"device_uid":   req.DeviceUID,
		"exp":          time.Now().Add(time.Hour * 20).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(apiCore.privateKey)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create JWT token"})
	}

	return c.JSON(fiber.Map{"message": "Success", "token": tokenString})
}

type managerLoginBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// managerLogin handles the login request for a manager.
//
// It parses the request body and checks if the manager's login and password are valid.
// If the login and password are valid, it generates a JWT token and returns it in the response.
//
// Parameters:
// - c: A pointer to the fiber.Ctx object representing the HTTP request context.
//
// Returns:
// - An error indicating any failures during the login process, or nil if the login is successful.
func managerLogin(c *fiber.Ctx) error {
	var req managerLoginBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	manager := controllers.GetManager(req.Login, req.Password)
	if manager == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := jwt.MapClaims{
		"id":            manager.ID,
		"manager_login": manager.Login,
		"exp":           time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return c.JSON(fiber.Map{
		"message": "Success",
		"token":   token,
	})
}

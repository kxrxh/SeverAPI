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

type ConductorLoginRequest struct {
	ConductorID   string `json:"conductor_id"`
	DeviceUID     string `json:"device_uid"`
	ConductorPass string `json:"conductor_pass"`
}

func conductorLogin(c *fiber.Ctx) error {
	req := ConductorLoginRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check if the conductor exists
	conductor := controllers.GetConductor(req.ConductorID, req.ConductorPass)
	if conductor == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Create JWT
	claims := jwt.MapClaims{
		"conductor_id": req.ConductorID,
		"device_uid":   req.DeviceUID,
		"exp":          time.Now().Add(time.Hour * 20).Unix(),
	}

	// Sign and get the complete encoded token as a string
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(apiCore.privateKey)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create JWT token",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"token":   tokenString,
	})

}

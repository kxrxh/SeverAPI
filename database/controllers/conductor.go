package controllers

import (
	"sever.hack/api/database"
	"sever.hack/api/database/models"
)

// GetConductor retrieves the conductor with the given ID and password from the database.
//
// Parameters:
// - conductorID: The ID of the conductor.
// - password: The password of the conductor.
//
// Returns:
// - *models.Conductor: The conductor object if found, otherwise nil.
func GetConductor(conductorID string, password string) *models.Conductor {
	var conductor models.Conductor
	database.PsqlDB.Where("conductor_id = ? AND conductor_pass = ?", conductorID, password).First(&conductor)
	// Check if conductor exists
	if conductor.ID == 0 {
		return nil
	}
	return &conductor
}

package controllers

import (
	"sever.hack/api/database"
	"sever.hack/api/database/models"
)

func GetConductor(conductorID string, password string) *models.Conductor {
	var conductor models.Conductor
	database.PsqlDB.Where("conductor_id = ? AND conductor_pass = ?", conductorID, password).First(&conductor)
	// Check if conductor exists
	if conductor.ID == 0 {
		return nil
	}
	return &conductor
}

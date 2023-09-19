package controllers

import (
	"time"

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

type ClientResponseStruct struct {
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Birth     time.Time        `json:"birth"`
	Benefits  []models.Benefit `json:"benefits"`
}

// GetClient retrieves a client and their benefits based on the provided client card UID.
//
// Parameters:
//   - clientCardUID: The UID of the client card.
//
// Returns:
//   - ClientResponseStruct: The response struct containing client information and benefits.
//   - bool: A boolean value indicating if the client and benefits were successfully retrieved.
func GetClient(clientCardUID string) (ClientResponseStruct, bool) {
	var client models.Client

	// Get client
	err := database.PsqlDB.
		Preload("Card").
		Where("client_cards.uid = ?", clientCardUID).
		First(&client).
		Error

	if err != nil {
		return ClientResponseStruct{}, false
	}

	return ClientResponseStruct{
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Birth:     client.Birth,
		Benefits:  getClientBenefits(client.ID),
	}, true
}

type UserDatabaseDumpStruct struct {
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Birth     string           `json:"birth"`
	Benefits  []models.Benefit `json:"benefits"`
	CardUID   string           `json:"card_uid"`
	State     string           `json:"state"`
}

func GetUserDump() ([]UserDatabaseDumpStruct, bool) {
	var dumps []UserDatabaseDumpStruct
	var clients []models.Client
	err := database.PsqlDB.Find(&clients).Error
	if err != nil {
		return []UserDatabaseDumpStruct{}, false
	}
	for _, client := range clients {
		dumps = append(dumps, UserDatabaseDumpStruct{
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Birth:     client.Birth.String(),
			Benefits:  getClientBenefits(client.ID),
			CardUID:   client.Card.UID,
			State:     client.Card.State.Name,
		})
	}
	return dumps, true
}

func getClientBenefits(clientID uint) []models.Benefit {
	var clientBenefits []models.ClientBenefit
	err := database.PsqlDB.
		Where("client_id = ?", clientID).
		Find(&clientBenefits).
		Error
	if err != nil {
		return []models.Benefit{}
	}

	var benefits []models.Benefit
	for _, clientBenefit := range clientBenefits {
		var benefit models.Benefit
		err = database.PsqlDB.Where("id = ?", clientBenefit.BenefitID).First(&benefit).Error
		if err != nil {
			continue
		}
		benefits = append(benefits, benefit)
	}

	return benefits
}

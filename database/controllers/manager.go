package controllers

import (
	"time"

	"sever.hack/api/database"
	"sever.hack/api/database/models"
)

// GetManager retrieves a manager from the database based on the provided login and password.
//
// Parameters:
//   - login: the login of the manager.
//   - password: the password of the manager.
//
// Returns:
//   - *models.Manager: a pointer to the retrieved manager.
func GetManager(login string, password string) *models.Manager {
	var manager models.Manager
	database.PsqlDB.Where("login = ? AND password = ?", login, password).First(&manager)
	// Check if manager exists
	if manager.ID == 0 {
		return nil
	}
	return &manager
}

func AddClient(user *models.Client, card *models.ClientCard) error {
	err := database.PsqlDB.Create(card).Error
	if err != nil {
		return err
	}
	return database.PsqlDB.Create(user).Error
}

func UpdateClient(user *models.Client) error {
	return database.PsqlDB.Save(user).Error
}

func AddConductor(user *models.Conductor) error {
	return database.PsqlDB.Create(user).Error
}

type DatabaseClient struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	MiddleName    string    `json:"middle_name"`
	Sex           string    `json:"sex"`
	Email         string    `json:"email"`
	City          string    `json:"city"`
	Phone         string    `json:"phone"`
	Birth         time.Time `json:"birth"`
	CardID        string    `json:"card_id"`
	Snils         string    `json:"snils"`
	AccountStates string    `json:"account_states"`
}

func GetAllDatabase() []DatabaseClient {
	var clients []models.Client
	err := database.PsqlDB.Find(&clients).Error
	if err != nil {
		return []DatabaseClient{}
	}
	var result []DatabaseClient
	for _, client := range clients {
		dbClient := DatabaseClient{
			FirstName:     client.FirstName,
			LastName:      client.LastName,
			MiddleName:    client.MiddleName,
			Sex:           client.Sex,
			Email:         client.Email,
			City:          client.City.Name,
			Phone:         client.Phone,
			Birth:         client.Birth,
			CardID:        client.Card.UID,
			Snils:         client.Snils,
			AccountStates: client.State.Name,
		}
		result = append(result, dbClient)
	}
	return result
}

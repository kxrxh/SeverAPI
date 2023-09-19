package controllers

import (
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

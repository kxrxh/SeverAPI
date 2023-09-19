package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sever.hack/api/database/models"
)

var PsqlDB *gorm.DB

// Init initializes the database connection.
//
// It takes three parameters: user, password, and dbname, which are used to connect to the database.
// The function does not return anything.
func Init(user string, password string, dbname string) {
	var err error
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow", user, password, dbname)
	PsqlDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}
	// Creating tables from models
	PsqlDB.AutoMigrate(&models.Client{}, &models.ClientCard{},
		&models.Operation{}, &models.Conductor{}, &models.ConductorCard{},
		&models.Manager{}, &models.Shift{}, &models.Benefit{}, &models.Partner{}, &models.Promotion{},
		&models.City{}, &models.State{}, &models.Route{}, &models.Transport{}, &models.ClientBenefit{})
}

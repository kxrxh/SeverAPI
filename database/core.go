package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sever.hack/api/database/models"
)

var Database *gorm.DB

func Init(user string, password string, dbname string) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow", user, password, dbname)
	Database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Database.AutoMigrate(&models.Client{}, &models.ClientCard{},
		&models.Operation{}, &models.Conductor{}, &models.ConductorCard{},
		&models.Manager{}, &models.Shift{}, &models.Benefit{}, &models.Partner{}, &models.Promotion{},
		&models.City{}, &models.State{}, &models.Route{}, &models.Transport{}, &models.ClientBenefit{})

}

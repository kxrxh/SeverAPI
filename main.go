package main

import (
	"os"

	"sever.hack/api/api"
	"sever.hack/api/database"
)

func main() {
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	api.Init(true)
	database.Init(dbUser, dbPassword, dbName)
	dbPort := os.Getenv("db_port")
	if dbPort == "" {
		dbPort = "5001"
	}
	api.Listen(dbPort)
}

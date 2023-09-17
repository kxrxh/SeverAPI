package main

import (
	"sever.hack/api/api"
	"sever.hack/api/database"
)

func main() {
	api.Init()
	database.Init("kxrxh", "0228", "sever-hack")
	api.Listen("5000")
}

package api

import "github.com/gofiber/fiber/v2"

var FiberApp *fiber.App

func Init() {
	FiberApp = fiber.New()
}

func Listen(port string) {
	FiberApp.Listen(":" + port)
}

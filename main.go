package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/web3santa/JWT-Authentication-Golang/database"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	app.Listen(":3000")

}

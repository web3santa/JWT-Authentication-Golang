package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/web3santa/JWT-Authentication-Golang/database"
	"github.com/web3santa/JWT-Authentication-Golang/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3000")

}

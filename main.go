package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/web3santa/JWT-Authentication-Golang/database"
	"github.com/web3santa/JWT-Authentication-Golang/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:3000",
	}))

	routes.Setup(app)

	app.Listen(":3000")

}

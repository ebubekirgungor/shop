package main

import (
	"log"
	"shop/database"
	"shop/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Static("/images", "./images")

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))
}

package main

import (
	"log"
	"shop/config"
	"shop/database"
	"shop/router"

	"github.com/ebubekirgungor/http2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
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

	s := &fasthttp.Server{
		Handler: app.Handler(),
	}

	if config.Config("HTTP2") == "true" {
		http2.ConfigureServer(s, http2.ServerConfig{})
	}

	log.Fatal(s.ListenAndServeTLS(":8080", "certs/RootCA.crt", "certs/RootCA.key"))
}

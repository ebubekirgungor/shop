package main

import (
	"log"
	"shop/config"
	"shop/database"
	"shop/router"

	"github.com/dgrr/http2"
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

	s := &fasthttp.Server{
		Handler: app.Handler(),
	}

	if config.Config("HTTP2") == "true" {
		http2.ConfigureServer(s, http2.ServerConfig{})
	}

	log.Fatal(s.ListenAndServeTLS(":8080", "/etc/letsencrypt/live/go-nuxt.shop/fullchain.pem", "/etc/letsencrypt/live/go-nuxt.shop/privkey.pem"))
}

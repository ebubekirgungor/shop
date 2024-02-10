package router

import (
	"shop/controllers"
	"shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/check_user", controllers.CheckUser)

	users := api.Group("/users")
	users.Get("/:id", middleware.Protected(), controllers.User)
	users.Post("/", controllers.AddUser)
	users.Patch("/:id", middleware.Protected(), controllers.UpdateUser)
	users.Patch("/password/:id", middleware.Protected(), controllers.UpdatePassword)
	users.Delete("/:id", middleware.Protected(), controllers.DeleteUser)

	addresses := api.Group("/addresses")
	addresses.Get("/:id", middleware.Protected(), controllers.Addresses)
	addresses.Post("/", middleware.Protected(), controllers.AddAddress)
	addresses.Patch("/:id", middleware.Protected(), controllers.UpdateAddress)
	addresses.Delete("/:id", middleware.Protected(), controllers.DeleteAddress)

	categories := api.Group("/categories")
	categories.Get("/", controllers.AllCategories)
	categories.Get("/:id", controllers.Category)
	categories.Post("/", middleware.Protected(), controllers.AddCategory)
	categories.Delete("/:id", middleware.Protected(), controllers.DeleteCategory)

	products := api.Group("/products")
	products.Get("/", controllers.AllProducts)
	products.Get("/:url", controllers.Product)
	products.Post("/", middleware.Protected(), controllers.AddProduct)
	products.Patch("/:id", middleware.Protected(), controllers.UpdateProduct)
	products.Delete("/:id", middleware.Protected(), controllers.DeleteProduct)

	orders := api.Group("/orders", middleware.Protected())
	orders.Get("/", controllers.AllOrders)
	orders.Get("/:id", controllers.Order)
	orders.Post("/", controllers.AddOrder)
	orders.Delete("/:id", controllers.DeleteOrder)
}

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
	users.Delete("/:id", middleware.Protected(), controllers.DeleteUser)

	categories := api.Group("/categories")
	categories.Get("/", controllers.AllCategories)
	categories.Get("/:id", controllers.Category)
	categories.Post("/", middleware.Protected(), controllers.AddCategory)
	categories.Delete("/:id", middleware.Protected(), controllers.DeleteCategory)

	products := api.Group("/products")
	products.Get("/", controllers.AllProducts)
	products.Get("/:id", controllers.Product)
	products.Post("/", middleware.Protected(), controllers.AddProduct)
	products.Delete("/:id", middleware.Protected(), controllers.DeleteProduct)

	orders := api.Group("/orders")
	orders.Get("/", middleware.Protected(), controllers.AllOrders)
	orders.Get("/:id", middleware.Protected(), controllers.Order)
	orders.Post("/", middleware.Protected(), controllers.AddOrder)
	orders.Delete("/:id", middleware.Protected(), controllers.DeleteOrder)
}

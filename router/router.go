package router

import (
	"shop/config"
	"shop/controllers"
	"shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group(config.Config("API_BASE"))

	api.Get("/search", controllers.Search)

	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/check_user", controllers.CheckUser)
	auth.Get("/logout", controllers.Logout)

	users := api.Group("/users")
	users.Get("/", middleware.Protected(), controllers.User)
	users.Post("/", controllers.AddUser)
	users.Patch("/", middleware.Protected(), controllers.UpdateUser)
	users.Patch("/password", middleware.Protected(), controllers.UpdatePassword)
	users.Delete("/", middleware.Protected(), controllers.DeleteUser)
	users.Get("/cart", middleware.Protected(), controllers.Cart)
	users.Patch("/cart", middleware.Protected(), controllers.UpdateCart)
	api.Get("/cart", controllers.UnregisteredCart)

	favorites := api.Group("/favorites")
	favorites.Get("/", middleware.Protected(), controllers.Favorites)
	favorites.Post("/:id", middleware.Protected(), controllers.AddFavorite)
	favorites.Delete("/:id", middleware.Protected(), controllers.DeleteFavorite)

	addresses := api.Group("/addresses", middleware.Protected())
	addresses.Get("/", controllers.Addresses)
	addresses.Post("/", controllers.AddAddress)
	addresses.Patch("/:id", controllers.UpdateAddress)
	addresses.Delete("/:id", controllers.DeleteAddress)

	categories := api.Group("/categories")
	categories.Get("/", controllers.AllCategories)
	categories.Get("/:id", controllers.Category)
	categories.Post("/", middleware.Protected(), controllers.AddCategory)
	categories.Delete("/:id", middleware.Protected(), controllers.DeleteCategory)

	products := api.Group("/products")
	products.Get("/", controllers.AllProducts)
	products.Get("/:id", controllers.Product)
	products.Post("/", middleware.Protected(), controllers.AddProduct)
	products.Patch("/:id", middleware.Protected(), controllers.UpdateProduct)
	products.Delete("/:id", middleware.Protected(), controllers.DeleteProduct)

	orders := api.Group("/orders", middleware.Protected())
	orders.Get("/", controllers.AllOrders)
	orders.Get("/:id", controllers.Order)
	orders.Post("/", controllers.AddOrder)
	orders.Delete("/:id", controllers.DeleteOrder)
}

package controllers

import (
	"shop/database"
	"shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Preload("Category").Preload("Orders").Find(&products)

	var all_products []fiber.Map
	for _, product := range products {
		all_products = append(all_products, fiber.Map{
			"id":             product.ID,
			"title":          product.Title,
			"category":       product.Category,
			"list_price":     product.ListPrice,
			"stock_quantity": product.StockQuantity,
			"orders":         product.Orders,
		})
	}
	return c.Status(200).JSON(all_products)
}

func Product(c *fiber.Ctx) error {
	product := models.Product{}
	database.DB.Db.Preload("Category").Preload("Orders").First(&product, c.Params("id"))
	return c.Status(200).JSON(fiber.Map{
		"id":             product.ID,
		"title":          product.Title,
		"category":       product.Category,
		"list_price":     product.ListPrice,
		"stock_quantity": product.StockQuantity,
		"orders":         product.Orders,
	})
}

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id := c.Query("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(id) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	database.DB.Db.Create(&product)

	return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	product := []models.Product{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&product)

	return c.Status(200).JSON("Product deleted")
}

package controllers

import (
	"shop/database"
	"shop/models"

	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Preload("Category").Where("title ILIKE ? AND stock_quantity > ?", "%"+c.Query("q")+"%", 0).Find(&products)

	var all_products []fiber.Map
	for _, product := range products {
		all_products = append(all_products, fiber.Map{
			"title":    product.Title,
			"url":      product.Url,
			"category": product.Category.Title,
		})
	}
	if all_products == nil {
		return c.Status(200).JSON([]string{})
	}
	return c.Status(200).JSON(all_products)
}

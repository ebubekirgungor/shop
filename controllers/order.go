package controllers

import (
	"shop/database"
	"shop/models"

	"github.com/gofiber/fiber/v2"
)

func AllOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Db.Preload("User").Preload("Products").Find(&orders)

	var all_orders []fiber.Map
	for _, order := range orders {
		all_orders = append(all_orders, fiber.Map{
			"id":       order.ID,
			"price":    order.Price,
			"user":     order.User,
			"products": order.Products,
		})
	}
	return c.Status(200).JSON(all_orders)
}

func Order(c *fiber.Ctx) error {
	order := models.Order{}
	database.DB.Db.Preload("User").Preload("Products").First(&order, c.Params("id"))
	return c.Status(200).JSON(fiber.Map{
		"id":       order.ID,
		"price":    order.Price,
		"user":     order.User,
		"products": order.Products,
	})
}

func AddOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&order)

	return c.Status(200).JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	order := []models.Order{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&order)

	return c.Status(200).JSON("Order deleted")
}

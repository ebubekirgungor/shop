package controllers

import (
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AllOrders(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.Preload("Orders.Products", func(db *gorm.DB) *gorm.DB {
		return db.Order("orders.created_at ASC")
	}).First(&user, userid)

	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	return c.Status(200).JSON(user.Orders)
}

func Order(c *fiber.Ctx) error {
	order := models.Order{}
	database.Db.Preload("Products").First(&order, c.Params("id"))

	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, order.UserId) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	return c.Status(200).JSON(fiber.Map{
		"id":       order.ID,
		"price":    order.Price,
		"products": order.Products,
	})
}

func AddOrder(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	order.UserId = uint(userid)

	database.Db.Create(&order)

	return c.Status(200).JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	order := models.Order{}

	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	database.Db.First(&order, c.Params("id"))

	if uint(userid) != order.UserId {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user id"})
	}

	database.Db.Where("id = ?", c.Params("id")).Delete(&order)

	return c.Status(200).JSON("Order deleted")
}

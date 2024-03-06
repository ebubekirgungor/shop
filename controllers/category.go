package controllers

import (
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AllCategories(c *fiber.Ctx) error {
	categories := []models.Category{}
	database.DB.Db.Find(&categories)
	return c.Status(200).JSON(categories)
}

func Category(c *fiber.Ctx) error {
	category := models.Category{}
	database.DB.Db.First(&category, c.Params("id"))
	return c.Status(200).JSON(category)
}

func AddCategory(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, userid) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(userid) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&category)

	return c.Status(200).JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, userid) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	category := models.Category{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&category)

	return c.Status(200).JSON("Address deleted")
}

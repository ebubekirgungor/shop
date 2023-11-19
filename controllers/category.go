package controllers

import (
	"shop/database"
	"shop/models"

	"github.com/gofiber/fiber/v2"
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
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&category)

	return c.Status(200).JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	category := []models.Category{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&category)

	return c.Status(200).JSON("Category deleted")
}

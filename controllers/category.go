package controllers

import (
	"fmt"
	"shop/database"
	"shop/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AllCategories(c *fiber.Ctx) error {
	categories := []models.Category{}
	database.Db.Find(&categories)
	return c.Status(200).JSON(categories)
}

func Category(c *fiber.Ctx) error {
	category := models.Category{}
	database.Db.First(&category, c.Params("id"))
	return c.Status(200).JSON(category)
}

func AddCategory(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(userid) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	category := new(models.Category)

	if form, err := c.MultipartForm(); err == nil {
		file := form.File["file"]
		if strings.HasPrefix(file[0].Header["Content-Type"][0], "image/") &&
			(strings.HasSuffix(file[0].Filename, ".jpg") ||
				strings.HasSuffix(file[0].Filename, ".jpeg") ||
				strings.HasSuffix(file[0].Filename, ".png")) {
			if err := c.SaveFile(file[0], fmt.Sprintf("./images/categories/%s", file[0].Filename)); err != nil {
				return err
			}
		} else {
			return c.Status(500).JSON(fiber.Map{"error": "Invalid file format"})
		}
	}

	category.Title = c.FormValue("title")
	category.Image = c.FormValue("image")
	database.Db.Create(&category)

	return c.Status(200).JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	category := models.Category{}
	database.Db.Where("id = ?", c.Params("id")).Delete(&category)

	return c.Status(200).JSON("Address deleted")
}

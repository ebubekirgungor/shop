package controllers

import (
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Favorites(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.Preload("Favorites").First(&user, userid)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	return c.Status(200).JSON(fiber.Map{
		"favorites": user.Favorites,
	})
}

func AddFavorite(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.First(&user, userid)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	product := models.Product{}
	product_id, _ := c.ParamsInt("id")

	if err := database.Db.First(&product, product_id).Error; err != nil {
		return err
	}

	if err := database.Db.Model(&user).Association("Favorites").Append(&product); err != nil {
		return err
	}

	return c.Status(200).JSON("ok")
}

func DeleteFavorite(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.First(&user, userid)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	product := models.Product{}
	product_id, _ := c.ParamsInt("id")

	if err := database.Db.First(&product, product_id).Error; err != nil {
		return err
	}

	if err := database.Db.Model(&user).Association("Favorites").Delete(&product); err != nil {
		return err
	}

	return c.Status(200).JSON("ok")
}

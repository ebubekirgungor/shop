package controllers

import (
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Addresses(c *fiber.Ctx) error {
	user := models.User{}
	database.DB.Db.Preload("Addresses").First(&user, c.Params("id"))
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	return c.Status(200).JSON(fiber.Map{
		"addresses": user.Addresses,
	})
}

func AddAddress(c *fiber.Ctx) error {
	address := new(models.Address)
	if err := c.BodyParser(address); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id := c.Query("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	address.UserId, _ = strconv.Atoi(c.Query("id"))

	database.DB.Db.Create(&address)

	return c.Status(200).JSON(address)
}

func UpdateAddress(c *fiber.Ctx) error {
	type UpdateAddressInput struct {
		Title   string `json:"title"`
		Address string `json:"address"`
	}
	var uai UpdateAddressInput
	if err := c.BodyParser(&uai); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}

	address := new(models.Address)
	database.DB.Db.First(&address, c.Params("id"))

	userid, _ := strconv.Atoi(c.Query("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, c.Query("userid")) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if userid != address.UserId {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user id"})
	}

	address.Title = uai.Title
	address.Address = uai.Address
	database.DB.Db.Save(&address)

	return c.JSON(address)
}

func DeleteAddress(c *fiber.Ctx) error {
	address := models.Address{}

	userid, _ := strconv.Atoi(c.Query("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, c.Query("userid")) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	database.DB.Db.First(&address, c.Params("id"))

	if userid != address.UserId {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user id"})
	}

	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&address)

	return c.Status(200).JSON("Address deleted")
}

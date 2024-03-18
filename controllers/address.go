package controllers

import (
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Addresses(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.Preload("Addresses", func(db *gorm.DB) *gorm.DB {
		return db.Order("addresses.created_at DESC")
	}).First(&user, userid)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	return c.Status(200).JSON(fiber.Map{
		"addresses": user.Addresses,
	})
}

func AddAddress(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	address := new(models.Address)
	if err := c.BodyParser(address); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	address.UserId = uint(userid)

	database.Db.Create(&address)

	return c.Status(200).JSON(address)
}

func UpdateAddress(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	type UpdateAddressInput struct {
		Title   string `json:"title"`
		Address string `json:"address"`
	}
	var uai UpdateAddressInput
	if err := c.BodyParser(&uai); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}

	address := new(models.Address)
	database.Db.First(&address, c.Params("id"))

	if uint(userid) != address.UserId {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user id"})
	}

	address.Title = uai.Title
	address.Address = uai.Address
	database.Db.Save(&address)

	return c.JSON(address)
}

func DeleteAddress(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	address := models.Address{}

	database.Db.First(&address, c.Params("id"))

	if uint(userid) != address.UserId {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user id"})
	}

	database.Db.Where("id = ?", c.Params("id")).Delete(&address)

	return c.Status(200).JSON("Address deleted")
}

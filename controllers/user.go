package controllers

import (
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {
	user := models.User{}
	database.DB.Db.First(&user, id)
	if user.Email == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

func User(c *fiber.Ctx) error {
	user := models.User{}
	database.DB.Db.First(&user, c.Params("email"))
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}
	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Couldn't hash password"})

	}

	user.Password = hash
	if err := database.DB.Db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Couldn't create user"})
	}

	return c.Status(200).JSON(fiber.Map{"email": user.Email})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Name string `json:"name"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := new(models.User)

	database.DB.Db.First(&user, id)
	user.Name = uui.Name
	database.DB.Db.Save(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}

	var password_input PasswordInput
	if err := c.BodyParser(&password_input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})

	}

	if !validUser(id, password_input.Password) {
		return c.Status(500).JSON(fiber.Map{"error": "Not valid user"})

	}

	user := new(models.User)

	database.DB.Db.First(&user, id)
	database.DB.Db.Delete(&user)

	return c.JSON("User deleted")
}

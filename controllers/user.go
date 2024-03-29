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

func validUser(id int, p string) bool {
	user := models.User{}
	database.Db.First(&user, id)
	if user.Email == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

func User(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(id)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.First(&user, id)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"error": "No user found with email"})
	}

	return c.Status(200).JSON(fiber.Map{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"phone":      user.Phone,
		"birthdate":  user.BirthDate,
		"gender":     user.Gender,
		"cart":       user.Cart,
	})
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "hash"})
	}

	user.Password = hash
	if err := database.Db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "create"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "email": user.Email})
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(id)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	type UpdateUserInput struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		BirthDate string `json:"birthdate"`
		Gender    string `json:"gender"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}

	user := new(models.User)
	database.Db.First(&user, id)

	user.FirstName = uui.FirstName
	user.LastName = uui.LastName
	user.Phone = uui.Phone
	user.BirthDate = uui.BirthDate
	user.Gender = uui.Gender
	database.Db.Save(&user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(id)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	type UpdatePasswordInput struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	var upi UpdatePasswordInput
	if err := c.BodyParser(&upi); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}

	user := new(models.User)
	database.Db.First(&user, id)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(upi.OldPassword)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid old password"})
	}

	new_hash, err := hashPassword(upi.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "hash"})
	}

	user.Password = new_hash
	database.Db.Save(&user)

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

	id, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(id)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})

	}

	if !validUser(id, password_input.Password) {
		return c.Status(500).JSON(fiber.Map{"error": "Not valid user"})
	}

	user := new(models.User)

	database.Db.First(&user, id)
	database.Db.Delete(&user)

	return c.JSON("User deleted")
}

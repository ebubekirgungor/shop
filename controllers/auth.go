package controllers

import (
	"errors"
	"strconv"
	"time"

	"shop/config"
	"shop/database"
	"shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserData struct {
	ID       uint   `json:"id"`
	Password string `json:"password"`
	Role     int    `json:"role"`
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

func isAdmin(id string) bool {
	user := models.User{}
	database.DB.Db.First(&user, id)
	return user.Role != 0
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUser(e string) (*models.User, error) {
	var user models.User
	if err := database.DB.Db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CheckUser(c *fiber.Ctx) error {
	type LoginInput struct {
		Email string `json:"email"`
	}
	input := new(LoginInput)
	user := models.User{}

	if err := c.BodyParser(&input); err != nil {
		return c.JSON("error")
	}

	database.DB.Db.Where("email = ?", input.Email).First(&user)

	if user.Email == "" {
		return c.JSON(false)
	} else {
		return c.JSON(true)
	}
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "server", "data": err})
	}

	email := input.Email
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	userModel, err = getUser(email)

	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "login", "data": err})
	} else {
		userData = UserData{
			ID:       userModel.ID,
			Role:     userModel.Role,
			Password: userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.JSON(fiber.Map{"status": "error", "message": "login", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": fiber.Map{
		"id":    userData.ID,
		"token": t,
		"role":  userData.Role,
	}})
}

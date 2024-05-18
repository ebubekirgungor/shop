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
	Role     uint8  `json:"role"`
}

func validToken(t *jwt.Token, id uint) bool {
	claims := t.Claims.(jwt.MapClaims)
	uid := uint(claims["user_id"].(float64))

	return uid == id
}

func isAdmin(id int) bool {
	user := models.User{}
	database.Db.First(&user, id)
	return user.Role != 0
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUser(e string) (*models.User, error) {
	var user models.User
	if err := database.Db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
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

	database.Db.Where("email = ?", input.Email).First(&user)

	if user.Email == "" {
		return c.JSON(false)
	} else {
		return c.JSON(true)
	}
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
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
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "server", "data": err})
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

	var expires time.Time

	if input.RememberMe {
		expires = time.Now().Add(time.Hour * 24 * 30)
	} else {
		expires = time.Now().Add(time.Hour * 1)
	}

	claims["exp"] = expires.Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "userid",
		Value:    strconv.Itoa(int(userData.ID)),
		Expires:  expires,
		Secure:   true,
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "role",
		Value:   strconv.Itoa(int(userData.Role)),
		Expires: expires,
		Secure:  true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  expires,
		Secure:   true,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"status": "success", "message": "ok"})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "userid",
		Value:    "",
		Expires:  time.Now(),
		Secure:   true,
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "role",
		Value:   "",
		Expires: time.Now(),
		Secure:  true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		Secure:   true,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"status": "success", "message": "ok"})
}

package controllers

import (
	"encoding/json"
	"fmt"
	"shop/database"
	"shop/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/datatypes"
)

func AllCategories(c *fiber.Ctx) error {
	categories := []models.Category{}
	database.Db.Find(&categories)
	return c.Status(200).JSON(categories)
}

func Category(c *fiber.Ctx) error {
	type Image struct {
		Name  string `json:"name"`
		Order uint   `json:"order"`
	}
	category := models.Category{}
	database.Db.Preload("Products").First(&category, "url = ?", c.Params("url"))
	var all_products []fiber.Map
	for _, product := range category.Products {
		var images []Image
		err := json.Unmarshal(product.Images, &images)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		image := "product.png"
		if len(images) != 0 {
			image = images[0].Name
		}
		all_products = append(all_products, fiber.Map{
			"id":             product.ID,
			"title":          product.Title,
			"url":            product.Url,
			"image":          image,
			"list_price":     product.ListPrice,
			"stock_quantity": product.StockQuantity,
			"filters":        product.Filters,
		})
	}
	if all_products == nil {
		return c.Status(200).JSON([]string{})
	}
	return c.Status(200).JSON(all_products)
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
	category.Url = c.FormValue("url")
	category.Filters = datatypes.JSON([]byte(c.FormValue("filters")))
	category.Image = c.FormValue("image")
	database.Db.Create(&category)

	return c.Status(200).JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(userid) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	category := new(models.Category)
	database.Db.First(&category, c.Params("id"))

	if c.FormValue("image") != "" {
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
		category.Image = c.FormValue("image")
	}
	category.Title = c.FormValue("title")
	category.Url = c.FormValue("url")
	category.Filters = datatypes.JSON([]byte(c.FormValue("filters")))
	database.Db.Save(&category)

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

	return c.Status(200).JSON("ok")
}

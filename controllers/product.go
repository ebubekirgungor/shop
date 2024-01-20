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

func AllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Preload("Category").Preload("Orders").Find(&products)

	var all_products []fiber.Map
	for _, product := range products {
		all_products = append(all_products, fiber.Map{
			"id":             product.ID,
			"title":          product.Title,
			"category":       product.Category,
			"list_price":     product.ListPrice,
			"stock_quantity": product.StockQuantity,
			"orders":         product.Orders,
		})
	}
	return c.Status(200).JSON(all_products)
}

func Product(c *fiber.Ctx) error {
	product := models.Product{}
	database.DB.Db.Preload("Category").Preload("Orders").First(&product, c.Params("id"))
	return c.Status(200).JSON(fiber.Map{
		"id":             product.ID,
		"title":          product.Title,
		"category":       product.Category,
		"list_price":     product.ListPrice,
		"stock_quantity": product.StockQuantity,
		"orders":         product.Orders,
	})
}

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if form, err := c.MultipartForm(); err == nil {
		files := form.File["file"]
		for _, file := range files {
			if strings.HasPrefix(file.Header["Content-Type"][0], "image/") &&
				(strings.HasSuffix(file.Filename, ".jpg") ||
					strings.HasSuffix(file.Filename, ".jpeg") ||
					strings.HasSuffix(file.Filename, ".png")) {
				if err := c.SaveFile(file, fmt.Sprintf("./images/products/%s", file.Filename)); err != nil {
					return err
				}
			} else {
				return c.Status(500).JSON(fiber.Map{"error": "Invalid file format"})
			}
		}
	}

	product.Title = c.FormValue("title")
	product.CategoryId, _ = strconv.Atoi(c.FormValue("category_id"))
	listprice, err := strconv.ParseFloat(c.FormValue("list_price"), 32)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product.ListPrice = float32(listprice)
	product.StockQuantity, _ = strconv.Atoi(c.FormValue("stock_quantity"))
	product.Images = c.FormValue("images")

	id := c.Query("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(id) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	database.DB.Db.Create(&product)

	return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	product := []models.Product{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&product)

	return c.Status(200).JSON("Product deleted")
}

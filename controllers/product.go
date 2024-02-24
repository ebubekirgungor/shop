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
	database.DB.Db.Preload("Category").Preload("Orders").Order("created_at").Find(&products)

	var all_products []fiber.Map
	for _, product := range products {
		all_products = append(all_products, fiber.Map{
			"id":             product.ID,
			"title":          product.Title,
			"url":            product.Url,
			"category":       product.Category.Title,
			"list_price":     product.ListPrice,
			"stock_quantity": product.StockQuantity,
			"orders":         product.Orders,
		})
	}
	if all_products == nil {
		return c.Status(200).JSON(products)
	}
	return c.Status(200).JSON(all_products)
}

func Product(c *fiber.Ctx) error {
	product := models.Product{}
	database.DB.Db.Preload("Category").Preload("Orders").Where("url = ?", c.Params("url")).First(&product)
	return c.Status(200).JSON(fiber.Map{
		"ID":             product.ID,
		"title":          product.Title,
		"category_id":    product.Category.ID,
		"list_price":     product.ListPrice,
		"stock_quantity": product.StockQuantity,
		"images":         product.Images,
		"orders":         product.Orders,
	})
}

func AddProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(id) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	check_product := models.Product{}
	database.DB.Db.First(&check_product, c.FormValue("url"))
	if check_product.Title != "" {
		return c.Status(500).JSON(fiber.Map{"error": "Product with given title is already exists"})
	}

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
	product.Url = c.FormValue("url")
	product.CategoryId, _ = strconv.Atoi(c.FormValue("category_id"))
	listprice, err := strconv.ParseFloat(c.FormValue("list_price"), 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid price value"})
	}
	product.ListPrice = float32(listprice)
	product.StockQuantity, _ = strconv.Atoi(c.FormValue("stock_quantity"))
	product.Images = c.FormValue("images")

	database.DB.Db.Create(&product)

	return c.Status(200).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(id) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	product := new(models.Product)
	database.DB.Db.First(&product, c.Params("id"))

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
	product.Url = c.FormValue("url")
	product.CategoryId, _ = strconv.Atoi(c.FormValue("category_id"))
	listprice, err := strconv.ParseFloat(c.FormValue("list_price"), 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid price value"})
	}
	product.ListPrice = float32(listprice)
	product.StockQuantity, _ = strconv.Atoi(c.FormValue("stock_quantity"))
	product.Images = c.FormValue("images")

	database.DB.Db.Save(&product)

	return c.Status(200).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	product := models.Product{}

	id := c.Query("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, c.Query("userid")) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	if !isAdmin(id) {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid role"})
	}

	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&product)

	return c.Status(200).JSON("Product deleted")
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/datatypes"
)

type CartData struct {
	ID       int  `json:"id"`
	Quantity int  `json:"quantity"`
	Selected bool `json:"selected"`
}

func Cart(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)
	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	products := []models.Product{}

	database.Db.First(&user, userid)

	var cartData []CartData
	err := json.Unmarshal(user.Cart, &cartData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(cartData) == 0 {
		return c.Status(200).JSON([]string{})
	}

	ids := []string{}
	for _, item := range cartData {
		ids = append(ids, fmt.Sprintf("%d", item.ID))
	}
	database.Db.Find(&products, ids)

	cartDataByID := make(map[int]CartData)
	for _, item := range cartData {
		cartDataByID[item.ID] = item
	}

	var found_products []fiber.Map
	for _, product := range products {
		cartItem, ok := cartDataByID[int(product.ID)]
		if !ok {
			cartItem = CartData{ID: int(product.ID), Quantity: 0, Selected: false}
		}
		found_products = append(found_products, fiber.Map{
			"id":             product.ID,
			"title":          product.Title,
			"url":            product.Url,
			"list_price":     product.ListPrice,
			"stock_quantity": product.StockQuantity,
			"images":         product.Images,
			"cart": fiber.Map{
				"id":       cartItem.ID,
				"quantity": cartItem.Quantity,
				"selected": cartItem.Selected,
			},
		})
	}
	if found_products == nil {
		return c.Status(200).JSON([]string{})
	}
	return c.Status(200).JSON(found_products)
}

func UnregisteredCart(c *fiber.Ctx) error {
	products := []models.Product{}
	var cartData []CartData

	cookie, urlerr := url.QueryUnescape(c.Cookies("cart"))
	if urlerr != nil {
		return c.Status(500).JSON(fiber.Map{"error": urlerr.Error()})
	}

	err := json.Unmarshal([]byte(cookie), &cartData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(cartData) == 0 {
		return c.Status(200).JSON([]string{})
	}

	ids := []string{}
	for _, item := range cartData {
		ids = append(ids, fmt.Sprintf("%d", item.ID))
	}
	database.Db.Find(&products, ids)

	cartDataByID := make(map[int]CartData)
	for _, item := range cartData {
		cartDataByID[item.ID] = item
	}

	var found_products []fiber.Map
	for _, product := range products {
		cartItem, ok := cartDataByID[int(product.ID)]
		if !ok {
			cartItem = CartData{ID: int(product.ID), Quantity: 0, Selected: false}
		}
		found_products = append(found_products, fiber.Map{
			"id":             product.ID,
			"title":          product.Title,
			"url":            product.Url,
			"list_price":     product.ListPrice,
			"stock_quantity": product.StockQuantity,
			"images":         product.Images,
			"cart": fiber.Map{
				"id":       cartItem.ID,
				"quantity": cartItem.Quantity,
				"selected": cartItem.Selected,
			},
		})
	}
	if found_products == nil {
		return c.Status(200).JSON([]string{})
	}
	return c.Status(200).JSON(found_products)
}

func UpdateCart(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	type CartInput struct {
		Cart datatypes.JSON `json:"cart"`
	}
	var ci CartInput
	if err := c.BodyParser(&ci); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error in input"})
	}

	database.Db.Model(&models.User{}).Where("id = ?", userid).Update("cart", ci.Cart)

	return c.JSON("ok")
}

func AddProductToCart(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	user := models.User{}
	database.Db.First(&user, userid)

	var cartData []CartData
	err := json.Unmarshal(user.Cart, &cartData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	id, _ := c.ParamsInt("id")

	for index, item := range cartData {
		if item.ID == id {
			cartData[index].Quantity = item.Quantity + 1
			cart, _ := json.Marshal(cartData)
			database.Db.Model(&models.User{}).Where("id = ?", userid).Update("cart", datatypes.JSON([]byte(cart)))
			return c.JSON("ok")
		}
	}

	var newCartData = CartData{
		ID:       id,
		Quantity: 1,
		Selected: true,
	}
	cartData = append(cartData, newCartData)

	cart, _ := json.Marshal(cartData)
	database.Db.Model(&models.User{}).Where("id = ?", userid).Update("cart", datatypes.JSON([]byte(cart)))

	return c.JSON("ok")
}

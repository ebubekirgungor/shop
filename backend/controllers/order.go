package controllers

import (
	"encoding/json"
	"shop/database"
	"shop/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/maps"
)

type OrderProduct struct {
	Title     string  `json:"title"`
	Url       string  `json:"url"`
	ListPrice float32 `json:"list_price"`
	Image     string  `json:"image"`
	Quantity  uint16  `json:"quantity"`
}

func AllOrders(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	orders := []models.Order{}
	database.Db.Order("created_at desc").Where("user_id = ?", userid).Find(&orders)

	var all_orders []fiber.Map
	for _, order := range orders {
		var products []OrderProduct
		err := json.Unmarshal(order.Products, &products)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		var all_products []fiber.Map
		for _, product := range products {
			all_products = append(all_products, fiber.Map{
				"url":      product.Url,
				"title":    product.Title,
				"image":    product.Image,
				"quantity": product.Quantity,
			})
		}
		all_orders = append(all_orders, fiber.Map{
			"id":              order.ID,
			"date":            order.CreatedAt,
			"price":           order.Price,
			"customer_name":   order.CustomerName,
			"products":        all_products,
			"delivery_status": order.DeliveryStatus,
		})
	}
	if all_orders == nil {
		return c.Status(200).JSON([]string{})
	}

	return c.Status(200).JSON(all_orders)
}

func Order(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	order := models.Order{}
	database.Db.First(&order, c.Params("id"))

	return c.Status(200).JSON(fiber.Map{
		"id":               order.ID,
		"date":             order.CreatedAt,
		"price":            order.Price,
		"customer_name":    order.CustomerName,
		"delivery_address": order.DeliveryAddress,
		"products":         order.Products,
		"delivery_status":  order.DeliveryStatus,
	})
}

func AddOrder(c *fiber.Ctx) error {
	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	type Cart struct {
		ID       uint `json:"id"`
		Quantity uint `json:"quantity"`
		Selected bool `json:"selected"`
	}
	type Image struct {
		Name  string `json:"name"`
		Order uint   `json:"order"`
	}
	user := models.User{}
	products := []models.Product{}

	database.Db.First(&user, userid)

	var cartData []Cart
	err := json.Unmarshal(user.Cart, &cartData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	ids := make(map[uint]uint)
	for _, item := range cartData {
		if item.Selected {
			ids[item.ID] = item.Quantity
		}
	}
	database.Db.Find(&products, maps.Keys(ids))

	var order_products []fiber.Map

	var price float32
	for _, product := range products {
		var images []Image
		err := json.Unmarshal(product.Images, &images)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		image := "product.png"
		if len(images) != 0 {
			image = images[0].Name
		}
		order_products = append(order_products, fiber.Map{
			"url":        product.Url,
			"title":      product.Title,
			"list_price": product.ListPrice,
			"image":      image,
			"quantity":   ids[product.ID],
		})
		price += product.ListPrice * float32(ids[product.ID])
	}

	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	order.UserId = uint(userid)
	order.Price = price
	order.Products, _ = json.Marshal(order_products)
	order.DeliveryStatus = models.InProgress
	database.Db.Create(&order)

	temp := cartData[:0]
	for _, item := range cartData {
		if !item.Selected {
			temp = append(temp, item)
		}
	}
	user.Cart, _ = json.Marshal(temp)
	database.Db.Save(&user)

	return c.Status(200).JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	order := models.Order{}

	userid, _ := strconv.Atoi(c.Cookies("userid"))
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, uint(userid)) {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid token id"})
	}

	database.Db.First(&order, c.Params("id"))

	if uint(userid) != order.UserId {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user id"})
	}

	database.Db.Where("id = ?", c.Params("id")).Delete(&order)

	return c.Status(200).JSON("Order deleted")
}

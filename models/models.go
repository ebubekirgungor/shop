package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null"`
	FirstName string         `gorm:"not null" json:"first_name"`
	LastName  string         `gorm:"not null" json:"last_name"`
	Phone     string         `json:"phone"`
	BirthDate string         `gorm:"type:date;default:null" json:"birthdate"`
	Gender    string         `gorm:"type:char(1)" json:"gender"`
	Role      uint8          `gorm:"size:1;default:0" json:"role"`
	Addresses []Address      `json:"addresses"`
	Orders    []Order        `json:"orders"`
	Cart      datatypes.JSON `gorm:"not null" json:"cart"`
	Favorites []*Product     `gorm:"many2many:user_products;" json:"favorites"`
}

type Address struct {
	gorm.Model
	Title        string `gorm:"uniqueIndex;not null" json:"title"`
	CustomerName string `gorm:"not null" json:"customer_name"`
	Address      string `gorm:"not null" json:"address"`
	UserId       uint   `gorm:"not null" json:"user_id"`
	User         User
}

type Category struct {
	gorm.Model
	Title    string    `gorm:"uniqueIndex;not null" json:"title"`
	Products []Product `json:"products"`
}

type Product struct {
	gorm.Model
	ID            uint   `gorm:"primarykey"`
	Title         string `gorm:"not null" json:"title"`
	Url           string `gorm:"not null" json:"url"`
	CategoryId    uint   `gorm:"not null" json:"category_id"`
	Category      Category
	ListPrice     float32        `gorm:"not null" json:"list_price"`
	StockQuantity uint16         `gorm:"not null" json:"stock_quantity"`
	Images        datatypes.JSON `gorm:"not null" json:"images"`
	Users         []*User        `gorm:"many2many:user_products;"`
}

type DeliveryStatus uint8

const (
	Delivered  DeliveryStatus = 0
	InProgress DeliveryStatus = 1
	Returned   DeliveryStatus = 2
	Canceled   DeliveryStatus = 3
)

type Order struct {
	gorm.Model
	Price           float32 `gorm:"not null" json:"price"`
	UserId          uint    `gorm:"not null" json:"user_id"`
	User            User
	Products        datatypes.JSON `gorm:"not null" json:"products"`
	CustomerName    string         `gorm:"not null" json:"customer_name"`
	DeliveryAddress string         `gorm:"not null" json:"delivery_address"`
	DeliveryStatus  DeliveryStatus `gorm:"not null" json:"delivery_status"`
}

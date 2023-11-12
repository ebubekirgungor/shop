package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Phone     int       `json:"phone"`
	BirthDate time.Time `json:"birthdate"`
	Gender    string    `gorm:"type:char(1)" json:"gender"`
	Orders    []Order   `json:"orders"`
}

type Category struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title"`
	Products []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Title         string `gorm:"not null" json:"title"`
	CategoryId    int    `gorm:"not null" json:"category_id"`
	Category      Category
	ListPrice     float32  `gorm:"not null" json:"list_price"`
	StockQuantity int      `gorm:"not null" json:"stock_quantity"`
	Orders        []*Order `gorm:"many2many:order_products;"`
}

type Order struct {
	gorm.Model
	Price    float32 `gorm:"not null" json:"price"`
	UserId   int     `gorm:"not null" json:"user_id"`
	User     User
	Products []*Product `gorm:"many2many:order_products;"`
}

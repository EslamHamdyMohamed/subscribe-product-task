package seeders

import (
	"subscribe-product/database"
	"subscribe-product/database/models/Product"
	"subscribe-product/database/models/user"
)

func Seed() {
	userSeeder()
	productSeeder()
}

func productSeeder() {
	products := []Product.Product{
		{
			Name: "product1",
		},
		{
			Name: "product2",
		},
	}

	db := database.DB
	db.Create(&products)
}

func userSeeder() {
	users := []user.User{
		{
			Name:  "user1",
			Email: "user1@u.com",
		},
		{
			Name:  "user2",
			Email: "user2@u.com",
		},
	}
	db := database.DB
	db.Create(&users)
}

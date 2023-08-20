package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"subscribe-product/database/models/Product"
	"subscribe-product/database/models/user"
)

var DB *gorm.DB
var err error

func ConnectToDatabase() {
	dsn := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") +
		")/" + os.Getenv("DATABASE_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println("Connected To Database")
	}
}

func MigrateAll() {
	DB.AutoMigrate(&user.User{})
	DB.AutoMigrate(&Product.Product{})
	DB.AutoMigrate(&Product.Subscription{})
}

package models

import (
	"subscribe-product/database"
	"subscribe-product/database/models/Product"
)

func FindOrFailProduct(productId uint) (bool, error, Product.Product) {
	var product Product.Product
	db := database.DB
	err := db.Where("id", productId).
		Find(&product).Error
	if product.ID == 0 {
		return false, err, product
	}
	return true, err, product
}

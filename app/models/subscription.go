package models

import (
	appHelper "subscribe-product/app/helpers"
	"subscribe-product/database"
	"subscribe-product/database/models/Product"
)

func CreateSubscription(subscription Product.Subscription) error {
	db := database.DB
	err := db.Create(&subscription).Error
	return err
}

func UpdateSubscriptionStatus(oldSubscription, UpdatedSubscription Product.Subscription) error {
	err := database.DB.Model(&oldSubscription).Updates(&UpdatedSubscription).Error
	return err
}

func FindOrFailSubscription(userId, productId uint, timeInterval appHelper.TimeInterval, status appHelper.Status) (bool, error, Product.Subscription) {
	var subscription Product.Subscription
	db := database.DB
	err := db.Where("user_id", userId).
		Where("product_id", productId).
		Where("status", status).
		Where("time_interval", timeInterval).
		Find(&subscription).Error
	if subscription.TransactionID == "" {
		return false, err, subscription
	}
	return true, err, subscription
}

func FindOrFailProduct(productId uint) (bool, error, Product.Product) {
	var product Product.Product
	db := database.DB
	err := db.Where("product_id", productId).
		Find(&product).Error
	if product.ID == 0 {
		return false, err, product
	}
	return true, err, product
}

func GetRequiredSubscriptions(status appHelper.Status, limit, offset int) (error, []Product.Subscription) {
	var subscriptions []Product.Subscription
	db := database.DB
	err := db.Where("status", appHelper.PAID).
		Where("end_date", ">=", "start_date").
		Limit(limit).Offset(offset).
		Find(&subscriptions).Error
	return err, subscriptions
}

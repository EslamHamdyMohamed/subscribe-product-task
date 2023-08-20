package models

import (
	appHelper "subscribe-product/app/exchange"
	"subscribe-product/database"
	"subscribe-product/database/models/Product"
	"time"
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

func GetRequiredSubscriptions(status appHelper.Status, limit, offset int) (error, []Product.Subscription) {
	var subscriptions []Product.Subscription
	currentDate := time.Now().Format("2006-01-02")
	db := database.DB
	err := db.Where("status", status).
		Where("end_date = ?", currentDate).
		Limit(limit).Offset(offset).
		Find(&subscriptions).Error
	return err, subscriptions
}

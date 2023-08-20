package services

import (
	"subscribe-product/app/exchange"
	"subscribe-product/app/exchange/requests"
	appHelper "subscribe-product/app/helpers"
	"subscribe-product/app/models"
	"subscribe-product/database/models/Product"
	"subscribe-product/helpers"
)

func SubscribeService(request requests.SubscriptionRequest) error {
	check, err, subscription := models.FindOrFailSubscription(request.UserID, request.ProductID, request.TimeInterval, exchange.PENDING)
	if helpers.CheckError(err) {
		return err
	}
	if check {
		err = UpdateBill(subscription, exchange.PAID)
	} else {
		err = GenerateBill(request, exchange.PAID)
	}
	return err
}

func GenerateBill(request requests.SubscriptionRequest, status exchange.Status) error {
	subscriptionRow := appHelper.SetSubscriptionRow(request, status)
	err := models.CreateSubscription(subscriptionRow)
	return err
}

func UpdateBill(subscription Product.Subscription, status exchange.Status) error {
	updated := appHelper.SetSubscriptionWithStatus(subscription, status)
	err := models.UpdateSubscriptionStatus(subscription, updated)
	return err
}

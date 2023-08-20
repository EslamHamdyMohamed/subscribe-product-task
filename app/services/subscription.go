package services

import (
	"subscribe-product/app/exchange/requests"
	appHelper "subscribe-product/app/helpers"
	"subscribe-product/app/models"
	"subscribe-product/database/models/Product"
	"subscribe-product/helpers"
)

func SubscribeService(request requests.SubscriptionRequest) error {
	check, err, subscription := models.FindOrFailSubscription(request.UserID, request.ProductID, request.TimeInterval, appHelper.PENDING)
	if helpers.CheckError(err) {
		return err
	}
	if check {
		err = PayBill(subscription)
	} else {
		err = GenerateBill(request, appHelper.PAID)
	}
	return err
}

func GenerateBill(request requests.SubscriptionRequest, status appHelper.Status) error {
	subscriptionRow := appHelper.SetSubscriptionRow(request, status)
	err := models.CreateSubscription(subscriptionRow)
	return err
}

func PayBill(subscription Product.Subscription) error {
	updated := appHelper.SetSubscriptionWithStatus(subscription, appHelper.PAID)
	err := models.UpdateSubscriptionStatus(subscription, updated)
	return err
}

package cron

import (
	"subscribe-product/app/exchange/requests"
	"subscribe-product/database/models/Product"
)

func SetSubscriptionRequest(subscription Product.Subscription) requests.SubscriptionRequest {
	return requests.SubscriptionRequest{
		UserID:       subscription.UserID,
		ProductID:    subscription.ProductID,
		TimeInterval: subscription.TimeInterval,
	}
}

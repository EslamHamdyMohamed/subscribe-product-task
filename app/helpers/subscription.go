package helpers

import (
	"subscribe-product/app/exchange/requests"
	"subscribe-product/database/models/Product"
	"time"
)

func SetSubscriptionRow(request requests.SubscriptionRequest, status Status) Product.Subscription {
	start, end := convertTimeIntervalToDates(request.TimeInterval)
	return Product.Subscription{
		ProductID:     request.ProductID,
		UserID:        request.UserID,
		TransactionID: request.TransactionID,
		TimeInterval:  request.TimeInterval,
		StartDate:     start,
		EndDate:       end,
		Status:        status,
	}
}

func SetSubscriptionWithStatus(subscription Product.Subscription, status Status) Product.Subscription {
	subscription.Status = status
	return subscription
}

func convertTimeIntervalToDates(timeInterval TimeInterval) (startDate, endDate time.Time) {
	var start = time.Now()
	var end time.Time
	switch timeInterval {
	case MONTHLY:
		end = start.AddDate(0, 1, 0)
	case YEARLY:
		end = start.AddDate(1, 0, 0)
	}
	return start, end
}

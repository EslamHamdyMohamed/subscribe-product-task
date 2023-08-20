package helpers

import (
	"subscribe-product/app/exchange"
	"subscribe-product/app/exchange/requests"
	"subscribe-product/database/models/Product"
	"time"
)

func SetSubscriptionRow(request requests.SubscriptionRequest, status exchange.Status) Product.Subscription {
	start, end := convertTimeIntervalToDates(request.TimeInterval)
	return Product.Subscription{
		ProductID:     request.ProductID,
		UserID:        request.UserID,
		TransactionID: request.TransactionID,
		TimeInterval:  request.TimeInterval,
		StartDate:     start,
		EndDate:       setTimeWithZeroHoursAndMinutes(end),
		Status:        status,
	}
}

func SetSubscriptionWithStatus(subscription Product.Subscription, status exchange.Status) Product.Subscription {
	subscription.Status = status
	return subscription
}

func convertTimeIntervalToDates(timeInterval exchange.TimeInterval) (startDate, endDate time.Time) {
	var start = time.Now()
	var end time.Time
	switch timeInterval {
	case exchange.MONTHLY:
		end = start.AddDate(0, 1, 0)
	case exchange.YEARLY:
		end = start.AddDate(1, 0, 0)
	}
	return start, end
}

func setTimeWithZeroHoursAndMinutes(t time.Time) time.Time {
	date := t.Format("2006-01-02")
	timeObj, _ := time.Parse("2006-01-02", date)
	return timeObj
}

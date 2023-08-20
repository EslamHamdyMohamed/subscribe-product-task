package requests

import (
	"subscribe-product/app/exchange"
)

type SubscriptionRequest struct {
	UserID        uint                  `json:"user_id" binding:"required"`
	ProductID     uint                  `json:"product_id" binding:"required"`
	TransactionID string                `json:"transaction_id" binding:"required"`
	TimeInterval  exchange.TimeInterval `json:"time_interval" binding:"required"`
}

package requests

import "subscribe-product/app/helpers"

type SubscriptionRequest struct {
	UserID        uint                 `json:"user_id" binding:"required"`
	ProductID     uint                 `json:"product_id" binding:"required"`
	TransactionID string               `json:"transaction_id" binding:"required"`
	TimeInterval  helpers.TimeInterval `json:"time_interval" binding:"required"`
}

package Product

import (
	"subscribe-product/app/helpers"
	"subscribe-product/database/models/user"
	"time"
)

type Subscription struct {
	UserID        uint                 `json:"user_id"`
	User          user.User            `json:"user"`
	ProductID     uint                 `json:"product_id"`
	Product       Product              `json:"product"`
	TransactionID string               `json:"transaction_id"`
	TimeInterval  helpers.TimeInterval `json:"time_interval"`
	StartDate     time.Time            `json:"start_date"`
	EndDate       time.Time            `json:"end_date"`
	Status        helpers.Status       `json:"status"`
}

package Product

import (
	"gorm.io/gorm"
	"subscribe-product/app/exchange"
	"subscribe-product/database/models/user"
	"time"
)

type Subscription struct {
	gorm.Model
	UserID        uint                  `json:"user_id"`
	User          user.User             `json:"user"`
	ProductID     uint                  `json:"product_id"`
	Product       Product               `json:"product"`
	TransactionID string                `json:"transaction_id"`
	TimeInterval  exchange.TimeInterval `json:"time_interval"`
	StartDate     time.Time             `json:"start_date"`
	EndDate       time.Time             `json:"end_date"`
	Status        exchange.Status       `json:"status"`
}

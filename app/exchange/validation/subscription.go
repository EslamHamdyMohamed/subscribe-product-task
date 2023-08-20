package validation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"subscribe-product/app/exchange/requests"
	"subscribe-product/app/models"
	"subscribe-product/helpers"
)

func ValidateSubscriptionRequest(c *gin.Context) (error, requests.SubscriptionRequest) {
	var request requests.SubscriptionRequest
	err := c.ShouldBindJSON(&request)
	if helpers.CheckError(err) {
		return err, request
	}

	ok, _, _ := models.FindOrFailProduct(request.ProductID)
	if !ok {
		return errors.New("product_id not found"), request
	}
	ok, _, _ = models.FindOrFailUser(request.UserID)
	if !ok {
		return errors.New("user_id not found"), request
	}
	if _, ok = TimeIntervalsMap[request.TimeInterval]; !ok {
		return errors.New("invalid time_interval"), request
	}
	return nil, request
}

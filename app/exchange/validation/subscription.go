package validation

import (
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
		return err, request
	}
	ok, _, _ = models.FindOrFailProduct(request.ProductID)
	if !ok {
		return err, request
	}
	if _, ok = TimeIntervalsMap[request.TimeInterval]; !ok {
		return err, request
	}
	return nil, request
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscribe-product/app/exchange/validation"
	"subscribe-product/app/services"
	"subscribe-product/helpers"
)

func Subscribe(c *gin.Context) {
	err, request := validation.ValidateSubscriptionRequest(c)
	if helpers.CheckError(err) {
		helpers.Response(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}
	err = services.SubscribeService(request)
	if helpers.CheckError(err) {
		helpers.Response(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}
	helpers.Response(c, http.StatusOK, "Process run successfully", "", nil)
}

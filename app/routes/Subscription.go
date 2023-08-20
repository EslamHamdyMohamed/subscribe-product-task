package routes

import (
	"github.com/gin-gonic/gin"
	"subscribe-product/app/controllers"
)

func SubscriptionRoutes(r *gin.Engine) {
	r.POST("/subscribe", controllers.Subscribe)
}

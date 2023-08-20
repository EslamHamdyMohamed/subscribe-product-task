package providers

import (
	"github.com/gin-gonic/gin"
	"os"
	"subscribe-product/app/routes"
	"subscribe-product/helpers"
)

func Run() {
	r := gin.Default()

	routes.SubscriptionRoutes(r)

	err := r.Run(":" + os.Getenv("APP_PORT_LOCAL"))
	helpers.CheckError(err)
}

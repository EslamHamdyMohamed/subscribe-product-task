package helpers

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status int, msg string, data, errors interface{}) {
	c.JSON(status, gin.H{
		"message": msg,
		"payload": data,
		"errors":  errors,
		"status":  status,
	})
}

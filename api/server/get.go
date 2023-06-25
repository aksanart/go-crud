package server

import (
	"aksan-go-crud/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func get(s service.ServiceIntf) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := s.GetService()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"status":  "FAILED",
				"message": fmt.Sprintf("%v", err.Error()),
			})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "message": "success", "data": data})
	}
}

package server

import (
	"aksan-go-crud/internal/entity"
	"aksan-go-crud/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func save(s service.ServiceIntf) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param entity.SaveRequest
		if err := c.ShouldBind(&param); err != nil {
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{
					"status":  "FAILED",
					"message": fmt.Sprintf("%v", err.Error()),
				})
				return
			}
		}
		err := s.SaveService(&param)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"status":  "FAILED",
				"message": fmt.Sprintf("%v", err.Error()),
			})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "message": "success"})
	}
}

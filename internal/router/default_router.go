package router

import (
	"github.com/gin-gonic/gin"
)

func SetupDefaultRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) { 
		c.JSON(200, gin.H{"status": "pong"}) 
	})
}

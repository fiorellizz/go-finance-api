package router

import (
	"github.com/gin-gonic/gin"
	"github.com/fiorellizz/go-finance-api/internal/handler"
)

func SetupRoutes(r *gin.Engine, th *handler.TransactionHandler) {
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"status": "pong"}) })

	api := r.Group("/api")
	{
		api.POST("/transactions", th.Create)
		api.GET("/transactions", th.List)
	}
}

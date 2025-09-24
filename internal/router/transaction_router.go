package router

import (
	"github.com/fiorellizz/go-finance-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupTransactionRoutes(r *gin.Engine, th *handler.TransactionHandler) {
	api := r.Group("/api")
	{
		api.POST("/transactions", th.Create)
		api.GET("/transactions", th.List)
		api.GET("/users/:id/transactions", th.ListByUser)
	}
}

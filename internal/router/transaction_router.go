package router

import (
	"github.com/fiorellizz/go-finance-api/internal/handler"
	"github.com/fiorellizz/go-finance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupTransactionRoutes(r *gin.Engine, th *handler.TransactionHandler) {
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/transactions", th.Create)
		api.GET("/transactions", th.List)
		api.GET("/users/:id/transactions", th.ListByUser)
		api.PUT("/transactions/:id", th.Update)
		api.DELETE("/transactions/:id", th.Delete)
		api.GET("/reports/balance", th.GetBalance)
    	api.GET("/reports/expenses-by-category", th.GetExpensesByCategory)
	}
}

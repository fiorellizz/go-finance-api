package router

import (
	"github.com/fiorellizz/go-finance-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, uh *handler.UserHandler) {
	api := r.Group("/api")
	{
		api.POST("/users", uh.Create)
		api.GET("/users", uh.List)
	}
}

package router

import (
	"github.com/fiorellizz/go-finance-api/internal/handler"
	"github.com/fiorellizz/go-finance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, uh *handler.UserHandler) {
	api := r.Group("/api")
	{
		api.POST("/register", uh.Create)
		api.POST("/login", uh.Login)
		api.GET("/users", middleware.AuthMiddleware(), uh.List) // opcional: proteger
		// api.GET("/users", uh.List)
	}
}

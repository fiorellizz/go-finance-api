package router

import (
	"github.com/fiorellizz/go-finance-api/internal/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, app *app.Application) {
	// rotas globais (ping, healthcheck etc.)
	SetupDefaultRoutes(r)

	// rotas de usuários
	SetupUserRoutes(r, app.UserHandler)

	// rotas de transações
	SetupTransactionRoutes(r, app.TransactionHandler)
}

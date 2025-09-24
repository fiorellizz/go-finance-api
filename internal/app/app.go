package app

import (
	"gorm.io/gorm"
	"github.com/fiorellizz/go-finance-api/internal/handler"
	"github.com/fiorellizz/go-finance-api/internal/repository"
	"github.com/fiorellizz/go-finance-api/internal/service"
)

type Application struct {
	TransactionHandler *handler.TransactionHandler
	UserHandler        *handler.UserHandler
}

func New(db *gorm.DB) *Application {
	// Transaction
	transactionRepo := repository.NewTransactionRepository(db)
	transactionSvc := service.NewTransactionService(transactionRepo)
	th := handler.NewTransactionHandler(transactionSvc)

	// User
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	uh := handler.NewUserHandler(userSvc)

	return &Application{
		TransactionHandler: th,
		UserHandler:        uh,
	}
}
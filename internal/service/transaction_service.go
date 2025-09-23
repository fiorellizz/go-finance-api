package service

import (
	"github.com/fiorellizz/go-finance-api/internal/domain"
	"github.com/fiorellizz/go-finance-api/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(r *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) Create(tx *domain.Transaction) error {
	return s.repo.Create(tx)
}

func (s *TransactionService) List() ([]domain.Transaction, error) {
	return s.repo.List()
}

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

// Create - cria uma transação
func (s *TransactionService) Create(tx *domain.Transaction) error {
	// regra de negócio poderia entrar aqui (ex: validar valores negativos, etc.)
	return s.repo.Create(tx)
}

func (s *TransactionService) List() ([]domain.Transaction, error) {
	return s.repo.List()
}

// ListByUser - lista transações de um usuário
func (s *TransactionService) ListByUser(userID uint) ([]domain.Transaction, error) {
	return s.repo.ListByUser(userID)
}

func (s *TransactionService) Update(id string, tx *domain.Transaction) error {
    return s.repo.Update(id, tx)
}

func (s *TransactionService) Delete(id string) error {
    return s.repo.Delete(id)
}

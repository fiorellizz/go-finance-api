package repository

import (
	"github.com/fiorellizz/go-finance-api/internal/domain"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(tx *domain.Transaction) error {
	return r.db.Create(tx).Error
}

// lista todas as transações
func (r *TransactionRepository) List() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Order("created_at desc").Find(&transactions).Error
	return transactions, err
}

// lista transações de um usuário específico
func (r *TransactionRepository) ListByUser(userID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&transactions).Error
	return transactions, err
}

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

func (r *TransactionRepository) Update(id string, tx *domain.Transaction) error {
    return r.db.Model(&domain.Transaction{}).Where("id = ?", id).Updates(tx).Error
}

func (r *TransactionRepository) Delete(id string) error {
    return r.db.Delete(&domain.Transaction{}, id).Error
}

// GetBalance - soma de receitas e despesas de um usuário
func (r *TransactionRepository) GetBalance(userID uint) (float64, error) {
    var income, expense float64
    if err := r.db.Model(&domain.Transaction{}).
        Where("user_id = ? AND type = ?", userID, "income").
        Select("COALESCE(SUM(amount),0)").Scan(&income).Error; err != nil {
        return 0, err
    }
    if err := r.db.Model(&domain.Transaction{}).
        Where("user_id = ? AND type = ?", userID, "expense").
        Select("COALESCE(SUM(amount),0)").Scan(&expense).Error; err != nil {
        return 0, err
    }
    return income - expense, nil
}

// GetExpensesByCategory - soma despesas agrupadas por categoria
func (r *TransactionRepository) GetExpensesByCategory(userID uint) (map[string]float64, error) {
    results := make(map[string]float64)
    rows, err := r.db.Model(&domain.Transaction{}).
        Select("category, COALESCE(SUM(amount),0) as total").
        Where("user_id = ? AND type = ?", userID, "expense").
        Group("category").Rows()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var category string
    var total float64
    for rows.Next() {
        rows.Scan(&category, &total)
        results[category] = total
    }
    return results, nil
}

package repository

import (
	"github.com/fiorellizz/go-finance-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(tx *domain.User) error {
	return r.db.Create(tx).Error
}

func (r *UserRepository) List() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Order("created_at desc").Find(&users).Error
	return users, err
}

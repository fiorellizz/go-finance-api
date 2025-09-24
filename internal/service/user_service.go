package service

import (
	"github.com/fiorellizz/go-finance-api/internal/domain"
	"github.com/fiorellizz/go-finance-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Criação de usuário com hash de senha
func (s *UserService) CreateUser(name, email, password string) (*domain.User, error) {
	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Listagem de usuários
func (s *UserService) ListUsers() ([]domain.User, error) {
	return s.repo.List()
}

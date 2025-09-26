package service

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/fiorellizz/go-finance-api/internal/domain"
	"github.com/fiorellizz/go-finance-api/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists         = errors.New("user with email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email, password string) (*domain.User, error) {
	// checar existência
	if u, _ := s.repo.FindByEmail(email); u != nil {
		return nil, ErrUserExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashed),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	// não retornar hash no objeto
	user.PasswordHash = ""
	return user, nil
}

func (s *UserService) ListUsers() ([]domain.User, error) {
	users, err := s.repo.List()
	// limpar password antes de retornar (segurança)
	for i := range users {
		users[i].PasswordHash = ""
	}
	return users, err
}

// Login gera JWT e retorna token + usuário (sem password)
func (s *UserService) Login(email, password string) (string, *domain.User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev_secret" // só dev; force usar env em prod
	}

	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(uint64(u.ID), 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	// limpar hash antes de devolver
	u.PasswordHash = ""
	return signed, u, nil
}

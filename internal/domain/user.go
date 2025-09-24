package domain

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" gorm:"unique" binding:"required,email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`

	Transactions []Transaction `json:"transactions,omitempty"`
}

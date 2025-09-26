package domain

import "time"

type User struct {
	ID           uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string        `json:"name" gorm:"size:100;not null"`
	Email        string        `json:"email" gorm:"size:150;uniqueIndex;not null"`
	PasswordHash string        `json:"-" gorm:"not null"`
	CreatedAt    time.Time     `json:"created_at" gorm:"autoCreateTime"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}

package domain

import "time"

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Amount    float64   `json:"amount" binding:"required"`
	Type      string    `json:"type" binding:"required,oneof=income expense"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

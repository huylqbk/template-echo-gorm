package models

import (
	"red-coins/config"
	"time"
)

type Transaction struct {
	ID        int    	`gorm:"primary_key;auto_increment" json:"id"`
	Cotation  float64   `validate:"required" gorm:"not null" json:"cotation"`
	Amount    float64   `validate:"required" gorm:"not null;" json:"amount"`
	Type      string    `validate:"required" gorm:"not null;" json:"type"`
	UserId    int    	`validate:"required" gorm:"not null" json:"user_id"`
	Date      string    `validate:"required" gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func TransactionStore(transaction *Transaction) bool {
	res := config.DB.Debug().Create(&transaction)

	if res.Error == nil {
		return true
	}

	return false
}
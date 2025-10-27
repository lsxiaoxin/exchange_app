package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FromCurrency string    `json:"from_currency" binding:"required"`
	ToCurrency   string    `json:"to_currency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"` 

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ExchangeRate) TableName() string {
	return "exchange_rates"
}

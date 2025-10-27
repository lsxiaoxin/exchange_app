package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"unique" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
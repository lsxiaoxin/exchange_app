package models

import "time"

type Article struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time
}


func (Article) TableName() string {
	return "article"
}
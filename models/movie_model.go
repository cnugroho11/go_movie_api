package models

import "time"

type Movie struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Rating      float32   `gorm:"not null" json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MovieUpdate struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string  `gorm:"type:varchar(255);not null" json:"title"`
	Description string  `gorm:"type:varchar(255);not null" json:"description"`
	Rating      float32 `gorm:"not null" json:"rating"`
	Image       string  `json:"image"`
}

type MovieInput struct {
	Title       string  `gorm:"type:varchar(255);not null" json:"title"`
	Description string  `gorm:"type:varchar(255);not null" json:"description"`
	Rating      float32 `gorm:"not null" json:"rating"`
	Image       string  `json:"image"`
}

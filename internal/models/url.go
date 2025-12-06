package models

import "time"

type URL struct {
	ID          uint      `gorm:"primaryKey"`
	OriginalURL string    `gorm:"not null"`
	ShortCode   string    `gorm:"uniqueIndex;not null"`
	Visits      int       `gorm:"default:0"`
	ExpiresAt   time.Time `gorm:"default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

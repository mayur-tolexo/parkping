package model

import "time"

type User struct {
	ID          uint   `gorm:"primaryKey"`
	PhoneNumber string `gorm:"uniqueIndex;size:15;not null"`
	CreatedAt   time.Time
}

package model

import "time"

type Vehicle struct {
	ID              uint   `gorm:"primaryKey"`
	UserID          uint   `gorm:"index;not null"`
	User            User   `gorm:"constraint:OnDelete:CASCADE"`
	VehicleNumber   string `gorm:"size:20;null"`
	VehicleType     string `gorm:"size:10;null"`
	FastagNumber    string `gorm:"uniqueIndex"`
	CallsEnabled    bool   `gorm:"default:true"`
	MessagesEnabled bool   `gorm:"default:true"`
	CreatedAt       time.Time
}

package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id     uint `gorm:"primaryKey"`
	Text   string
	UserID uint
	User   User `gorm:"foreignKey:Id"`
}

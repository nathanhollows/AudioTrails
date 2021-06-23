package models

import "gorm.io/gorm"

// Library keeps track of the pages that each player has found
type Library struct {
	gorm.Model
	Page Page `gorm:"foreignKey:ID"`
	User User `gorm:"foreignKey:ID"`
}

package models

import "gorm.io/gorm"

// Library keeps track of the pages that each player has found
type Library struct {
	gorm.Model
	Page Geosite `gorm:"foreignKey:ID"`
	User string
}

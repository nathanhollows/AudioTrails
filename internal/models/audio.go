package models

import (
	"gorm.io/gorm"
)

// Audio keeps track of uploaded audio
type Audio struct {
	gorm.Model
	Title  string `gorm:"unique"`
	Author string
	File   string
}

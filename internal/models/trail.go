package models

import "gorm.io/gorm"

// Trail stores tags that can group pages
type Trail struct {
	gorm.Model
	Trail string `gorm:"index:unique"`
}

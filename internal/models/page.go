package models

import (
	"gorm.io/gorm"
)

// Page stores a static page
type Page struct {
	gorm.Model
	Code      string `gorm:"unique"`
	Title     string
	Text      string
	Author    string
	Published bool `sql:"DEFAULT:false"`
}

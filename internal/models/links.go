package models

import (
	"gorm.io/gorm"
)

// Link stores a static page that can be accessed via QR code
type Link struct {
	gorm.Model
	Code      string `gorm:"unique"`
	Title     string
	Author    string
	URL       string
	Hits      uint
	Published bool `sql:"DEFAULT:false"`
	System    bool `sql:"DEFAULT:false"`
}

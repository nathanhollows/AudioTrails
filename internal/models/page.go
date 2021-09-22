package models

import (
	"gorm.io/gorm"
)

// Page stores a static page that can be accessed via QR code
type Page struct {
	gorm.Model
	Code      string `gorm:"unique"`
	Title     string
	Text      string
	Author    string
	Published bool    `sql:"DEFAULT:false"`
	System    bool    `sql:"DEFAULT:false"`
	TrailID   int     `sql:"DEFAULT:NULL"`
	Trail     Trail   `gorm:"references:ID"`
	GalleryID int     `sql:"DEFAULT:NULL"`
	Gallery   Gallery `gorm:"references:ID"`
}

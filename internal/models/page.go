package models

import (
	"github.com/nathanhollows/Argon/internal/helpers"
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

// BeforeCreate will generate a new code for the page
func (p *Page) BeforeCreate(tx *gorm.DB) (err error) {
	p.Code = helpers.NewCode(5)
	return nil
}

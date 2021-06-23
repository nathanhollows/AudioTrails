package models

import (
	"github.com/nathanhollows/AmazingTrace/pkg/helpers"
	"gorm.io/gorm"
)

// Page stores a static page that can be accessed via QR code
type Page struct {
	gorm.Model
	Code   string `gorm:"unique"`
	Title  string
	Text   string
	Author string
}

// BeforeCreate will generate a new code for the page
func (c *Page) BeforeCreate(tx *gorm.DB) (err error) {
	c.Code = helpers.NewCode(5)
	return nil
}

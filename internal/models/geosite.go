package models

import (
	"gorm.io/gorm"
)

// Geosite stores a static page that can be accessed via QR code
type Geosite struct {
	gorm.Model
	Code      string `gorm:"unique"`
	Title     string
	Text      string
	Author    string
	Published bool `sql:"DEFAULT:false"`
	System    bool `sql:"DEFAULT:false"`
	CoverID   int
	Cover     Media `gorm:"foreignKey:CoverID"`
	AudioID   int
	Audio     Media `gorm:"foreignKey:AudioID"`
}

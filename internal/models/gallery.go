package models

import "gorm.io/gorm"

// Gallery stores the gallery that a page can belong to
type Gallery struct {
	gorm.Model
	Gallery string
}

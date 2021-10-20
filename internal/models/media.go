package models

import (
	"fmt"

	"github.com/nathanhollows/Argon/internal/helpers"
	"gorm.io/gorm"
)

// Media keeps track of uploaded files
type Media struct {
	gorm.Model
	Title  string
	File   string
	Type   string
	Format string
	Hash   string
}

// URL returns the URL for the given media object
func (m Media) URL() string {
	return helpers.URL(fmt.Sprint("public/uploads/", m.Type, "/", m.File))
}

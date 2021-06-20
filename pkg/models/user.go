package models

import "gorm.io/gorm"

// User represents both players and admins
type User struct {
	gorm.Model
	Email string
	Code  string
	Admin bool
}

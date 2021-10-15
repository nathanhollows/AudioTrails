package models

import "gorm.io/gorm"

// ScanEvent tracks the scan events
type ScanEvent struct {
	gorm.Model
	UserID      string `sql:"DEFAULT:NULL"`
	UserID   string `sql:"DEFAULT:NULL"`
	PageCode string
	Page     Page `gorm:"foreignKey:PageCode;references:Code"`
	Link        Link `gorm:"foreignKey:LinkCode;references:Code"`
}

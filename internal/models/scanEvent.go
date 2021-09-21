package models

import "gorm.io/gorm"

// ScanEvent tracks the scan events
type ScanEvent struct {
	gorm.Model
	UserID string `sql:"DEFAULT:NULL"`
	PageID int
	Page   Page `gorm:"foreignKey:PageID"`
}

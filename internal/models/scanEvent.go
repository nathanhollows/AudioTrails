package models

import "gorm.io/gorm"

// ScanEvent tracks the scan events
type ScanEvent struct {
	gorm.Model
	UserID      string `sql:"DEFAULT:NULL"`
	GeositeCode string
	Geosite     Geosite `gorm:"foreignKey:GeositeCode;references:Code"`
	LinkCode    string
	Link        Link `gorm:"foreignKey:LinkCode;references:Code"`
	UserAgent   string
}

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

// FoundPage holds the values of a custom query
type FoundPage struct {
	Code    string
	Title   string
	Gallery string
	Trail   string
	Seen    bool
}

// FindPagesByUserQuery returns []FoundPage for a given user
var FindPagesByUserQuery = `SELECT pages.code, pages.title, gallery, trails.trail, scan.seen
FROM galleries
JOIN pages ON pages.gallery_id = galleries.id
JOIN trails ON trails.id = pages.trail_id
LEFT JOIN
	(SELECT page_code, true AS seen
		FROM scan_events
		WHERE user_id = ?)
	AS scan ON scan.page_code = pages.code
WHERE pages.deleted_at IS NULL
AND pages.published IS TRUE
GROUP BY pages.code
ORDER BY trail, gallery;`

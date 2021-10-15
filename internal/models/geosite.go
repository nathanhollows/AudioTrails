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
}

// ResultFoundPages holds the values of a custom query
type ResultFoundPages struct {
	Code    string
	Title   string
	Gallery string
	Trail   string
	Seen    bool
}

// QueryFindPagesByUser returns []FoundPage for a given user
var QueryFindPagesByUser = `SELECT geosites.code, geosites.title, gallery, trails.trail, scan.seen
FROM galleries
JOIN geosites ON geosites.gallery_id = galleries.id
JOIN trails ON trails.id = pages.trail_id
LEFT JOIN
	(SELECT page_code, true AS seen
		FROM scan_events
		WHERE user_id = ?)
	AS scan ON scan.page_code = pages.code
WHERE geosites.deleted_at IS NULL
AND geosites.published IS TRUE
GROUP BY geosites.code
ORDER BY trail, gallery;`

// ResultsTrailCounts holds the values of a custom query
type ResultsTrailCounts struct {
	Gallery string
	Trail   string
	Found   int
	Unfound int
}

// QueryTrailCountByUser is a query that returns the number of trails found / unfound
var QueryTrailCountByUser = `SELECT gallery, trails.trail, count(scan.seen) as found, count(trails.trail) as unfound
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
GROUP BY gallery, trail
ORDER BY trail, gallery;`

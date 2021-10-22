package admin

import (
	"net/http"

	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Analytics shows player stats
func Analytics(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	data := make(map[string]interface{})
	data["title"] = "Analytics | Admin"
	data["section"] = "analytics"

	geoScans := []models.ScanEvent{}
	env.DB.Where("geosite_code != ''").Preload(clause.Associations).Order("created_at DESC").Limit(15).Find(&geoScans)
	data["geoScans"] = geoScans

	linkScans := []models.ScanEvent{}
	env.DB.Model(linkScans).Where("link_code != ''").Preload(clause.Associations).Order("created_at DESC").Limit(15).Find(&linkScans)
	data["linkScans"] = linkScans

	var today int64
	env.DB.Raw("SELECT count(*) FROM scan_events WHERE date(created_at) = date('now')").Find(&today)
	data["today"] = today
	var week int64
	env.DB.Raw("SELECT count(*) FROM scan_events WHERE date(created_at) > date('now', '-7 day')").Find(&week)
	data["week"] = week
	var month int64
	env.DB.Raw("SELECT count(*) FROM scan_events WHERE date(created_at) > date('now', '-1 month')").Find(&month)
	data["month"] = month
	var year int64
	env.DB.Raw("SELECT count(*) FROM scan_events WHERE date(created_at) > date('now', '-12 month')").Find(&year)
	data["year"] = year

	return render(w, data, "analytics/index.html")
}

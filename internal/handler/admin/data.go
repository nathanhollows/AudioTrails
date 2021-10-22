package admin

import (
	"fmt"
	"net/http"

	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// DataDump generates a CSV of all scan data
func DataDump(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/csv")

	data := make(map[string]interface{})

	scans := []models.ScanEvent{}
	env.DB.Model(scans).Preload(clause.Associations).Order("created_at DESC").Limit(15).Find(&scans)
	data["linkScans"] = scans

	fmt.Fprint(w, "time,user,geosite.code,geosite.title,link.code,link.title,userAgent\r\n")
	for _, scan := range scans {
		fmt.Fprintf(w, "\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\r\n", scan.CreatedAt, scan.UserID, scan.GeositeCode, scan.Geosite.Title, scan.LinkCode, scan.Link.Title, scan.UserAgent)
	}

	return nil
}

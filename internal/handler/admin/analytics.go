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

	scans := []models.ScanEvent{}
	env.DB.Preload(clause.Associations).Order("created_at DESC").Find(&scans)
	data["dump"] = scans

	return render(w, data, "analytics/index.html")
}

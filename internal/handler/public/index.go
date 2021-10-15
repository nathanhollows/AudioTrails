package public

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Index is the homepage of the game.
// Prints a very simple page asking only for a team code.
func Index(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "index"

	session, err := env.Session.Get(r, "uid")
	if err != nil || session.Values["id"] == nil {
		session, err = env.Session.New(r, "uid")
		session.Options.HttpOnly = true
		session.Options.SameSite = http.SameSiteStrictMode
		session.Options.Secure = true
		id := uuid.New()
		session.Values["id"] = id.String()
		session.Save(r, w)
	}

	var found []models.ScanEvent
	env.DB.Model(&models.ScanEvent{}).Where("user_id = ?", session.Values["id"]).Preload(clause.Associations).Distinct("geosite_code").Scan(&found)
	data["found"] = found

	var geosites []models.Geosite
	env.DB.Model(&models.Geosite{}).Scan(&geosites)
	data["geosites"] = geosites

	return render(w, data, "index/index.html")
}

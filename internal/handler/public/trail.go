package public

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Trail shows the user all the unlocked pages they have found
func Trail(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "trail"
	data["title"] = "QR Audio Trail"

	session, err := env.Session.Get(r, "uid")
	if err != nil || session.Values["id"] == nil {
		fmt.Println(err)
		session, err = env.Session.New(r, "uid")
		session.Options.HttpOnly = true
		session.Options.SameSite = http.SameSiteStrictMode
		session.Options.Secure = true
		id := uuid.New()
		session.Values["id"] = id.String()
		session.Save(r, w)
	}

	geosites := []models.Geosite{}
	env.DB.Raw("SELECT DISTINCT geosites.* FROM geosites LEFT JOIN  (select geosite_code, user_id from scan_events WHERE user_id = ?) AS scan_events ON scan_events.geosite_code = geosites.code WHERE user_id IS NULL AND deleted_at IS NULL AND published", session.Values["id"]).Distinct().Preload(clause.Associations).Find(&geosites)
	data["geosites"] = geosites

	found := []models.Geosite{}
	env.DB.Raw("SELECT DISTINCT geosites.* FROM geosites LEFT JOIN scan_events ON scan_events.geosite_code = geosites.code WHERE geosites.deleted_at IS NULL AND user_id = ? AND published;", session.Values["id"]).Preload(clause.Associations).Find(&found)
	data["found"] = found

	return render(w, data, "trail/index.html")
}

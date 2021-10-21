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

	geosites := []models.Geosite{}
	env.DB.Where("published = true").Preload(clause.Associations).Find(&geosites)
	data["geosites"] = geosites

	library := []models.Library{}
	env.DB.Find(&library)
	data["library"] = library

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

	return render(w, data, "trail/index.html")
}

package public

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Page delivers the page relating to a particular code.
// This function does not track scan events.
func Page(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["section"] = "library"

	code := chi.URLParam(r, "code")
	page := models.Geosite{}
	env.DB.Where("Code = ?", code).Preload(clause.Associations).Find(&page)
	if page.Code == "" {
		flash.Set(w, r, flash.Message{Message: "That's not a valid code"})
		http.Redirect(w, r, "/404", http.StatusFound)
		return nil
	}

	if !page.Published {
		flash.Set(w, r, flash.Message{Message: "This page is not yet public", Style: "warning"})
	}

	var count int64
	env.DB.Model(models.Geosite{}).Where("published = true").Count(&count)
	data["count"] = count

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
	var trails []models.ResultsTrailCounts
	env.DB.Raw(models.QueryTrailCountByUser, session.Values["id"]).Scan(&trails)
	data["trails"] = trails

	data["title"] = page.Title
	data["md"] = parseMD(page.Text)
	data["page"] = page

	data["messages"] = flash.Get(w, r)
	return render(w, data, "page/discovered.html")
}

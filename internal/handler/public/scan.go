package public

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
)

// ScanGeosite handles the scanned URL.
// This functions track scan events
func ScanGeosite(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	code := strings.ToUpper(chi.URLParam(r, "code"))
	geosite := models.Geosite{}
	env.DB.Where("Code = ?", code).Find(&geosite)
	if geosite.Code == "" {
		flash.Set(w, r, flash.Message{Message: "That's not a valid code"})
		http.Redirect(w, r, "/404", http.StatusFound)
		return nil
	}

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

	scan := models.ScanEvent{}
	scan.Geosite = geosite
	scan.UserID = fmt.Sprint(session.Values["id"])
	scan.UserAgent = r.UserAgent()
	env.DB.Model(&models.ScanEvent{}).Create(&scan)

	http.Redirect(w, r, fmt.Sprintf("/%s", geosite.Code), http.StatusTemporaryRedirect)
	return nil
}

// ScanLink handles the scanned URL.
// This functions track scan events
func ScanLink(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	code := strings.ToUpper(chi.URLParam(r, "code"))

	link := models.Link{}
	env.DB.Where("Code = ?", code).Find(&link)
	if link.Code == "" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return nil
	}

	scan := models.ScanEvent{}
	scan.Link = link
	scan.UserID = "untracked"
	scan.UserAgent = r.UserAgent()
	env.DB.Model(&models.ScanEvent{}).Create(&scan)
	link.Hits++
	env.DB.Save(link)

	http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
	return nil
}

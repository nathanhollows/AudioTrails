package public

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
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

	code := chi.URLParam(r, "code")
	page := models.Page{}
	env.DB.Where("Code = ?", code).Preload(clause.Associations).Find(&page)
	if page.Code == "" {
		flash.Set(w, r, flash.Message{Message: "That's not a valid code"})
		http.Redirect(w, r, "/404", 302)
		return nil
	}

	if !page.Published {
		flash.Set(w, r, flash.Message{Message: "This page is not yet public", Style: "warning"})
	}

	var count int64
	env.DB.Model(models.Page{}).Where("published = true AND gallery_id = ?", page.GalleryID).Count(&count)
	data["count"] = count

	data["title"] = page.Title
	data["md"] = parseMD(page.Text)
	data["page"] = page

	data["messages"] = flash.Get(w, r)
	return render(w, data, "page/index.html")
}

// Scan handles the scanned URL.
// This functions track scan events
func Scan(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	code := chi.URLParam(r, "code")
	page := models.Page{}
	env.DB.Where("Code = ?", code).Find(&page)
	if page.Code == "" {
		flash.Set(w, r, flash.Message{Message: "That's not a valid code"})
		http.Redirect(w, r, "/404", 302)
		return nil
	}

	http.Redirect(w, r, fmt.Sprintf("/%s", page.Code), http.StatusTemporaryRedirect)
	return nil
}

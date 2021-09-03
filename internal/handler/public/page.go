package public

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
)

// Page handles the scanned URL.
func Page(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)

	code := chi.URLParam(r, "code")
	page := models.Page{}
	env.DB.Where("Code = ?", code).Find(&page)
	if page.Code == "" {
		flash.Set(w, r, flash.Message{Message: "That's not a valid code"})
		http.Redirect(w, r, "/404", 302)
		return nil
	}

	data["title"] = page.Title
	data["md"] = parseMD(page.Text)
	data["page"] = page

	return render(w, data, "page/index.html")
}

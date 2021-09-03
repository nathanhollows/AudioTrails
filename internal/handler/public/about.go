package public

import (
	"net/http"

	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
)

// About tells visitors about the project
func About(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	page := models.Page{}
	env.DB.Where("code = ?", "about").Find(&page)
	data["md"] = parseMD(page.Text)
	data["messages"] = flash.Get(w, r)
	return render(w, data, "about/index.html")
}

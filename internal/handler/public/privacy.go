package public

import (
	"net/http"

	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
)

// Privacy tells visitors about the project terms and privacy conditions
func Privacy(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	page := models.Page{}
	env.DB.Where("code = ?", "privacy").Find(&page)
	data["md"] = parseMD(page.Text)
	data["messages"] = flash.Get(w, r)
	return render(w, data, "privacy/index.html")
}

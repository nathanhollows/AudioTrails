package admin

import (
	"net/http"

	"github.com/nathanhollows/Argon/internal/handler"
)

// Media manages assests, go figure.
// E.g. images, videos, audio
func Media(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	data := make(map[string]interface{})
	data["title"] = "Media"

	return render(w, data, "media/index.html")
}

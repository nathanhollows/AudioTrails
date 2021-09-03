package admin

import (
	"net/http"

	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
)

// Dashboard handles the teams and shows the current status
func Dashboard(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	flash.Set(w, r, flash.Message{Message: "You need to be logged in to play.", Style: "warning"})

	data["messages"] = flash.Get(w, r)
	data["title"] = "Admin"

	return render(w, data, "dashboard/index.html")
}

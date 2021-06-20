package public

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

// Library shows the user all the unlocked pages they have found
func Library(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/templates/index.html",
		"../web/templates/flash.html",
		"../web/views/library/index.html"))

	if err := templates.ExecuteTemplate(w, "base", env.Data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

package public

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

// Index is the homepage of the game.
// Prints a very simple page asking only for a team code.
func Index(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	// flash.Set(w, r, flash.Message{Title: "Test", Message: "This is a test!", Style: "danger"})

	templates := template.Must(template.ParseFiles(
		"../web/templates/index.html",
		"../web/templates/flash.html",
		"../web/views/index/index.html"))

	if err := templates.ExecuteTemplate(w, "base", env.Data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

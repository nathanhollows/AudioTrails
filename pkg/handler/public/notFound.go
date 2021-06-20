package public

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

// NotFound is the 404 handler.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/templates/index.html",
		"../web/views/errors/notfound.html"))

	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
}

func Error404(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/templates/index.html",
		"../web/views/errors/notfound.html"))

	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

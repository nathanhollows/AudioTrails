package public

import (
	"html/template"
	"log"
	"net/http"
)

// NotFound is the 404 handler.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		"../web/template/errors/notfound.html"))

	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
}

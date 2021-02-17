package handler

import (
	"html/template"
	"log"
	"net/http"
)

// Clue handles the scanned URL.
// Either shows an error, the next clue, an opportunity, or a challenge.
func Clue(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	url := r.URL.String()[1:6]

	var page string = "error"

	clue, err := env.Manager.GetClue(url)
	if err == nil {
		page = "clue"
	}

	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		"../web/template/clue/"+page+".html"))

	if err := templates.ExecuteTemplate(w, "base", clue); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

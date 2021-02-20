package handler

import (
	"html/template"
	"log"
	"net/http"
)

// Admin handles the teams and shows the current status
func Admin(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		"../web/template/admin/index.html"))

	if err := templates.ExecuteTemplate(w, "base", env.Manager); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

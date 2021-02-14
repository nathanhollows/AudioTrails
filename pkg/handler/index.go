package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		"../web/template/index/index.html"))

	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

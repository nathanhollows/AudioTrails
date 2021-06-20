package admin

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

// Codes list all available codes
func Codes(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	templates := template.Must(template.ParseFiles(
		"../web/templates/admin.html",
		"../web/templates/flash.html",
		"../web/views/admin/codes.html"))

	if err := templates.ExecuteTemplate(w, "base", env.Manager); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

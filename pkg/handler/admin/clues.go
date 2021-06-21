package admin

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/handler"
	"github.com/nathanhollows/AmazingTrace/pkg/models"
)

// Clues lists all available clues
func Clues(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	clues := []models.Clue{}
	result := env.DB.Find(&clues)
	if result.RowsAffected > 0 {
		env.Data["clues"] = clues
	}

	templates := template.Must(template.ParseFiles(
		"../web/templates/admin.html",
		"../web/templates/flash.html",
		"../web/views/admin/clues.html"))

	if err := templates.ExecuteTemplate(w, "base", env.Data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

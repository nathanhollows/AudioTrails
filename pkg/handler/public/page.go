package public

import (
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/nathanhollows/AmazingTrace/pkg/flash"
	"github.com/nathanhollows/AmazingTrace/pkg/handler"
	"github.com/nathanhollows/AmazingTrace/pkg/models"
)

// Page handles the scanned URL.
func Page(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	code := chi.URLParam(r, "code")
	page := models.Page{}
	env.DB.Where("Code = ?", code).Find(&page)
	if page.Code == "" {
		flash.Set(w, r, flash.Message{Message: "That's not a valid code"})
		http.Redirect(w, r, "/404", 302)
		return nil
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.HardLineBreak | parser.Footnotes | parser.Mmark
	parser := parser.NewWithExtensions(extensions)
	md := string(markdown.ToHTML([]byte(page.Text), parser, nil))

	data := make(map[string]interface{})
	data["title"] = page.Title
	data["md"] = md
	data["page"] = page

	templates := template.Must(template.ParseFiles(
		"../web/templates/index.html",
		"../web/templates/flash.html",
		"../web/views/public/page/index.html"))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

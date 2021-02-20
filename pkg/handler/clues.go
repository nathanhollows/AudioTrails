package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/game"
)

// Clues shows a team all of the clues they have unlocked
func Clues(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	type Data struct {
		Code  string
		Team  game.Team
		Clues []game.Clue
	}

	r.ParseForm()
	team := r.PostForm.Get("team")
	data := Data{
		Code:  team,
		Team:  env.Manager.GetTeam(team),
		Clues: game.Clues,
	}

	var page string
	if env.Manager.CheckTeam(team) {
		page = "../web/template/clues/index.html"
	} else {
		page = "../web/template/clues/error.html"
	}

	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		page,
	))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

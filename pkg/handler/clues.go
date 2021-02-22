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
		Clue game.Clue
		Team game.Team
	}
	var page string

	r.ParseForm()
	teamCode := r.PostForm.Get("code")
	team := &game.Team{}
	index, err := env.Manager.GetTeam(teamCode)
	if err != nil {
		page = "../web/template/clues/error.html"
	} else {
		page = "../web/template/clues/index.html"
		team = &env.Manager.Teams[index]
		team.CheckIn()
	}

	data := Data{
		Clue: game.Clue{},
		Team: *team,
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

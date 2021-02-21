package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/nathanhollows/AmazingTrace/pkg/game"
)

// Clue handles the scanned URL.
func Clue(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	clueCode := (r.URL.String()[1:6])
	clueCode = strings.ToUpper(clueCode)
	var page string = "404"

	clue, err := env.Manager.GetClue(clueCode)
	if err == nil {
		page = "index"
	}

	// When a team is entered
	if r.Method == "POST" {
		r.ParseForm()
		teamCode := r.Form.Get("code")
		index, err := env.Manager.GetTeam(teamCode)
		team := &game.Team{}
		if err != nil {
			page = "notateam"
		} else {
			page = "index"
			team = &env.Manager.Teams[index]
			team.LastSeen.Local().Hour()
			if team.Solve(clueCode) != nil {
				page = "error"
			}
		}
		team.CheckIn()

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

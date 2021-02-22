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

	type Data struct {
		Clue game.Clue
		Team game.Team
	}
	data := Data{}

	if len(r.URL.String()) != 6 {
		templates := template.Must(template.ParseFiles(
			"../web/template/index.html",
			"../web/template/errors/notfound.html"))

		if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
			http.Error(w, err.Error(), 0)
			log.Print("Template executing error: ", err)
		}

		return nil
	}
	clueCode := (r.URL.String()[1:6])
	clueCode = strings.ToUpper(clueCode)
	var page string = "clue/404"

	clue, err := env.Manager.GetClue(clueCode)
	if err == nil {
		page = "clue/index"
	}
	data.Clue = clue

	// When a team is entered
	if r.Method == "POST" {
		r.ParseForm()
		teamCode := r.Form.Get("code")
		index, err := env.Manager.GetTeam(teamCode)
		team := &game.Team{}
		if err != nil {
			page = "clue/notateam"
		} else {
			team = &env.Manager.Teams[index]
			team.LastSeen.Local().Hour()
			if team.Solve(clueCode) != nil {
				page = "clue/dungoofed"
			} else {
				page = "clue/yougotit"
			}
		}
		team.CheckIn()
		data.Team = *team
	}

	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		"../web/template/"+page+".html"))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

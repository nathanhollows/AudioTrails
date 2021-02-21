package handler

import (
	"html/template"
	"log"
	"net/http"
)

// Start begins the game for the team. Prints out their first clue
func Start(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	type Data struct {
		Team string
	}

	r.ParseForm()
	teamCode := r.Form.Get("code")
	data := Data{
		Team: teamCode,
	}

	var page string
	index, err := env.Manager.GetTeam(teamCode)
	if err != nil {
		page = "../web/template/index/error.html"
	} else {
		team := &env.Manager.Teams[index]
		team.CheckIn()
		page = "../web/template/start/index.html"
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

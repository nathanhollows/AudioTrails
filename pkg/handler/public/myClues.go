package public

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/nathanhollows/AmazingTrace/pkg/game"
	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

// Clues shows a team all of the clues they have unlocked
func Clues(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	type Data struct {
		Clue     game.Clue
		Team     game.Team
		TimeLeft float64
	}
	data := Data{}

	var page string

	r.ParseForm()
	teamCode := r.PostForm.Get("code")
	team := &game.Team{}
	index, err := env.Manager.GetTeam(teamCode)
	if err != nil {
		page = "../web/views/clues/error.html"
	} else {
		page = "../web/views/clues/index.html"
		team = &env.Manager.Teams[index]
		team.CheckIn()
		team.Status = ""
	}

	env.Data["clue"] = Data{
		Clue: game.Clue{},
		Team: *team,
	}
	end := time.Date(2021, time.February, 24, 12, 0, 0, 0, time.Local)
	fmt.Println(end)
	data.TimeLeft = math.Floor(end.Sub(time.Now()).Minutes())

	templates := template.Must(template.ParseFiles(
		"../web/templates/index.html",
		"../web/templates/flash.html",
		page,
	))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

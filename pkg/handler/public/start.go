package public

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

// Start begins the game for the team. Prints out their first clue
func Start(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
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
		session, _ := env.Session.Get(r, "trace")
		session.Values["code"] = team.Code
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
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

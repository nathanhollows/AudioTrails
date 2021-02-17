package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Start begins the game for the team. Prints out their first clue
func Start(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	r.ParseForm()
	team := r.PostForm.Get("code")

	fmt.Println("The team code is: ", team)
	if team != "" {
		fmt.Println("The team code is: ", team)
		// Check if the team exists
	}

	templates := template.Must(template.ParseFiles(
		"../web/template/index.html",
		"../web/template/start/index.html"))

	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

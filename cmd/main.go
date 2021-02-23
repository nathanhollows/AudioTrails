// AmazingTrace is a QR code based scavenger hunt
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
	"github.com/nathanhollows/AmazingTrace/pkg/filesystem"
	"github.com/nathanhollows/AmazingTrace/pkg/game"
	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

var router *chi.Mux
var env handler.Env

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))

	var store = sessions.NewCookieStore([]byte("trace"))

	env = handler.Env{
		Manager: game.Manager{},
		Session: store,
	}
	env.Manager.CreateTeams(50)
}

func main() {
	routes()
	fmt.Println(http.ListenAndServe(":8000", router))
}

func routes() {
	router.Handle("/", handler.Handler{Env: &env, H: handler.Index})
	router.Handle("/wxvan", handler.Handler{Env: &env, H: handler.Index})
	router.Handle("/WXVAN", handler.Handler{Env: &env, H: handler.Index})
	router.Handle("/start", handler.Handler{Env: &env, H: handler.Start})
	router.Handle("/admin", handler.Handler{Env: &env, H: handler.Admin})
	router.Handle("/admin/ff", handler.Handler{Env: &env, H: handler.FastForward})
	router.Handle("/admin/hinder", handler.Handler{Env: &env, H: handler.Hinder})
	router.Handle("/admin/codes", handler.Handler{Env: &env, H: handler.Codes})
	router.Handle("/clues", handler.Handler{Env: &env, H: handler.Clues})
	router.Handle("/{/[A-z0-9]{5}}", handler.Handler{Env: &env, H: handler.Clue})
	router.NotFound(handler.NotFound)

	workDir, _ := os.Getwd()
	filesDir := filesystem.Myfs{http.Dir(filepath.Join(workDir, "../web/static"))}
	filesystem.FileServer(router, "/public", filesDir)
}

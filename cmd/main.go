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
	"github.com/nathanhollows/AmazingTrace/pkg/handler/admin"
	"github.com/nathanhollows/AmazingTrace/pkg/handler/public"
)

var router *chi.Mux
var env handler.Env

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))

	var store = sessions.NewCookieStore([]byte("trace"))

	env = handler.Env{
		Manager: game.Manager{},
		Session: store,
		Data:    make(map[string]interface{}),
	}
}

func main() {
	routes()
	fmt.Println(http.ListenAndServe(":8000", router))
}

// Set up the routes needed for the game.
func routes() {
	router.Handle("/", handler.Handler{Env: &env, H: public.Index})

	router.Handle("/start", handler.Handler{Env: &env, H: public.Start})
	router.Handle("/library", handler.Handler{Env: &env, H: public.Library})
	router.Handle("/clues", handler.Handler{Env: &env, H: public.Clues})

	router.Handle("/{[A-z0-9]{5}}", handler.Handler{Env: &env, H: public.Clue})

	router.Handle("/login", handler.Handler{Env: &env, H: public.Login})
	router.Handle("/register", handler.Handler{Env: &env, H: public.Register})

	router.Handle("/admin", handler.Admin{Env: &env, H: admin.Admin})
	router.Handle("/admin/ff", handler.Admin{Env: &env, H: admin.FastForward})
	router.Handle("/admin/hinder", handler.Admin{Env: &env, H: admin.Hinder})
	router.Handle("/admin/codes", handler.Admin{Env: &env, H: admin.Codes})

	router.Handle("/404", handler.Handler{Env: &env, H: public.Error404})
	router.NotFound(public.NotFound)

	workDir, _ := os.Getwd()
	filesDir := filesystem.Myfs{Dir: http.Dir(filepath.Join(workDir, "../web/static"))}
	filesystem.FileServer(router, "/public", filesDir)
}

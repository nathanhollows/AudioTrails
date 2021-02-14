// AmazingTrace is a QR code based scavenger hunt
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nathanhollows/AmazingTrace/pkg/filesystem"
	"github.com/nathanhollows/AmazingTrace/pkg/handler"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))
}

func main() {
	routes()
	fmt.Println(http.ListenAndServe(":8000", router))
}

func routes() {
	router.Handle("/", handler.Handler{H: handler.Index})

	workDir, _ := os.Getwd()
	filesDir := filesystem.Myfs{http.Dir(filepath.Join(workDir, "../web/static"))}
	filesystem.FileServer(router, "/public", filesDir)
}

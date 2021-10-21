// Argon is a QR code based education platform
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
	"github.com/nathanhollows/Argon/internal/filesystem"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/handler/admin"
	"github.com/nathanhollows/Argon/internal/handler/public"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *chi.Mux
var env handler.Env

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))

	var store sessions.Store
	if key, ok := os.LookupEnv("GEOTRACE_SESSION_KEY"); ok {
		store = sessions.NewCookieStore([]byte(key))
	} else {
		panic("env var GEOTRACE_SESSION_KEY must be set")
	}

	db, err := gorm.Open(sqlite.Open("trace.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	env = handler.Env{
		Session: store,
		DB:      *db,
		Data:    make(map[string]interface{}),
	}
}

func main() {
	env.DB.AutoMigrate(
		&models.Media{},
		&models.Page{},
		&models.User{},
		&models.Admin{},
		&models.Geosite{},
		&models.Link{},
		&models.ScanEvent{},
		&models.Library{},
	)
	routes()
	if key, ok := os.LookupEnv("GEOTRACE_PORT"); ok {
		fmt.Println(http.ListenAndServe(":"+key, router))
	} else {
		fmt.Println(http.ListenAndServe(":8050", router))
	}
}

// Set up the routes needed for the game.
func routes() {
	router.Handle("/", handler.HandlePublic{Env: &env, H: public.Index})

	var pages = []models.Page{}
	env.DB.Model(models.Page{}).Find(&pages)
	for _, page := range pages {
		router.Handle(fmt.Sprintf("/%v", page.Code), handler.HandlePublic{Env: &env, H: public.RegularPage})
	}

	router.Handle("/trail", handler.HandlePublic{Env: &env, H: public.Trail})

	router.Handle("/{code:[A-z]{5}}", handler.HandlePublic{Env: &env, H: public.Page})
	router.Handle("/s/{code:[A-z]{5}}", handler.HandlePublic{Env: &env, H: public.ScanGeosite})
	router.Handle("/l/{code:[A-z]{5}}", handler.HandlePublic{Env: &env, H: public.ScanLink})
	router.Handle("/qr/{location:[A-z]{1}}/{code:[A-z]{5}} - {fluff}.{format:[A-z]{3}}", handler.HandlePublic{Env: &env, H: public.QR})

	router.Handle("/login", handler.HandlePublic{Env: &env, H: public.Login})
	router.Handle("/logout", handler.HandlePublic{Env: &env, H: public.Logout})

	router.Handle("/admin", handler.HandleAdmin{Env: &env, H: admin.Geosites})
	router.Handle("/admin/media", handler.HandleAdmin{Env: &env, H: admin.Media})
	router.Handle("/admin/upload", handler.HandleAdmin{Env: &env, H: admin.Upload})
	router.Handle("/admin/analytics", handler.HandleAdmin{Env: &env, H: admin.Analytics})
	router.Handle("/admin/geosites", handler.HandleAdmin{Env: &env, H: admin.Geosites})
	router.Handle("/admin/geosites/delete", handler.HandleAdmin{Env: &env, H: admin.DeleteGeosite})
	router.Handle("/admin/geosites/delete/{code:[A-z]{5}}", handler.HandleAdmin{Env: &env, H: admin.DeleteGeosite})
	router.Handle("/admin/geosites/restore", handler.HandleAdmin{Env: &env, H: admin.Restore})
	router.Handle("/admin/geosites/edit/{code}", handler.HandleAdmin{Env: &env, H: admin.EditGeosite})
	router.Handle("/admin/geosites/create", handler.HandleAdmin{Env: &env, H: admin.CreateGeosite})
	router.Handle("/admin/geosites/preview", handler.HandleAdmin{Env: &env, H: admin.PreviewMD})
	router.Handle("/admin/links", handler.HandleAdmin{Env: &env, H: admin.Links})
	router.Handle("/admin/links/delete", handler.HandleAdmin{Env: &env, H: admin.DeleteLink})
	router.Handle("/admin/links/restore", handler.HandleAdmin{Env: &env, H: admin.RestoreLink})
	router.Handle("/admin/links/edit/{code}", handler.HandleAdmin{Env: &env, H: admin.EditLink})
	router.Handle("/admin/links/create", handler.HandleAdmin{Env: &env, H: admin.CreateLink})
	router.Handle("/admin/links/preview", handler.HandleAdmin{Env: &env, H: admin.PreviewMD})

	router.Handle("/404", handler.HandlePublic{Env: &env, H: public.Error404})
	router.NotFound(public.NotFound)

	workDir, _ := os.Getwd()
	filesDir := filesystem.Myfs{Dir: http.Dir(filepath.Join(workDir, "web/static"))}
	filesystem.FileServer(router, "/public", filesDir)
}

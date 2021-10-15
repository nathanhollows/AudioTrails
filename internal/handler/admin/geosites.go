package admin

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/helpers"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Geosites shows admin the geosites
func Geosites(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["title"] = "Geosites"
	data["section"] = "geosites"

	pages := []models.Geosite{}
	result := env.DB.Preload(clause.Associations).Find(&pages)
	if result.RowsAffected > 0 {
		data["pages"] = pages
	}

	return render(w, data, "geosites/index.html")
}

// DeleteGeosite removes the given geosite from the database
func DeleteGeosite(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		r.ParseForm()
		code := r.PostFormValue("page")
		result := env.DB.Where("Code = ?", code).Delete(&models.Geosite{})
		if result.RowsAffected != 0 {
			flash.Set(w, r, flash.Message{Message: "Deleted page", Style: "success"})
			http.Redirect(w, r, helpers.URL("admin/pages"), http.StatusFound)
		} else {
			flash.Set(w, r, flash.Message{Message: "Could not delete page", Style: "warning"})
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
	} else {
		flash.Set(w, r, flash.Message{Message: "Invalid request", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
	return nil
}

// Restore will undo soft delete on the last deleted geosite
func Restore(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	var page = models.Geosite{}
	env.DB.Unscoped().Where("deleted_at IS NOT NULL").Order("deleted_at DESC").Limit(1).Find(&page)
	result := env.DB.Model(&page).Updates(map[string]interface{}{"deleted_at": nil})
	if result.RowsAffected == 0 {
		flash.Set(w, r, flash.Message{Message: "Could not restore page", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		flash.Set(w, r, flash.Message{Message: "Page restored", Style: "success"})
		http.Redirect(w, r, helpers.URL("admin/geosites/edit/"+page.Code), http.StatusFound)
	}
	return nil
}

// EditGeosite allows the user to view and edit the given geosites
func EditGeosite(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})

	code := chi.URLParam(r, "code")
	page := models.Geosite{}
	result := env.DB.Where("code = ?", code).Preload(clause.Associations).Find(&page)

	if result.RowsAffected == 0 {
		return handler.StatusError{Code: http.StatusNotFound, Err: errors.New("page cannot be found")}
	}

	if r.Method == http.MethodPost || r.Method == http.MethodPatch {
		r.ParseForm()
		if _, ok := r.PostForm["delete"]; ok {
			result = env.DB.Delete(&page)
			if result.RowsAffected == 0 {
				flash.Set(w, r, flash.Message{Message: "Could not delete page", Style: "danger"})
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			} else {
				flash.Set(w, r, flash.Message{Message: "Page deleted. <a href=\"" + helpers.URL("admin/geosites/restore") + "\">Undo</a>", Style: "success"})
				http.Redirect(w, r, helpers.URL("admin/pages"), http.StatusFound)
			}
			return nil
		}

		if val, ok := r.PostForm["title"]; ok {
			page.Title = val[0]
		}
		if val, ok := r.PostForm["content"]; ok {
			page.Text = val[0]
		}
		if val, ok := r.PostForm["author"]; ok {
			page.Author = val[0]
		}
		if _, ok := r.PostForm["publish"]; ok {
			page.Published = true
		} else {
			page.Published = false
		}
		result = env.DB.Save(&page)
		if result.RowsAffected == 0 {
			if r.Method == http.MethodPost {
				flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			}
			http.Error(w, "could not save page", http.StatusBadGateway)
			return nil
		}
		if r.Method == http.MethodPost {
			flash.Set(w, r, flash.Message{Message: "Saved!", Style: "success"})
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
		return nil
	}

	data["page"] = page
	data["title"] = "Editing " + page.Title
	data["section"] = "geosites"

	data["messages"] = flash.Get(w, r)

	return render(w, data, "geosites/edit.html")

}

// CreateGeosite removes the given geosite from the database
func CreateGeosite(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["title"] = "Create a Geosite"
	data["section"] = "geosites"

	if r.Method == http.MethodPost {
		r.ParseForm()
		page := models.Geosite{}
		page.Title = r.FormValue("title")
		page.Text = r.FormValue("content")
		for {
			page.Code = helpers.NewCode(5)
			check := models.Geosite{}
			env.DB.Model(models.Geosite{}).Where("code = ?", page.Code).Find(&check)
			if check.Code != page.Code {
				break
			}
		}

		result := env.DB.Model(&models.Geosite{}).Create(&page)
		if result.Error != nil {
			flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			return nil
		}

		flash.Set(w, r, flash.Message{Message: "Created page!", Style: "success"})
		http.Redirect(w, r, helpers.URL("admin/geosites/edit/"+page.Code), http.StatusFound)
		return nil
	}

	return render(w, data, "geosites/create.html")

}

// PreviewMD accepts MD and returns HTML
func PreviewMD(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {

		type markdown struct {
			Md string `json:"md"`
		}
		var response markdown
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}

		t := template.Must(template.New("md").Parse("{{.}}"))
		t.Execute(w, parseMD(response.Md))

		return nil
	}
	return handler.StatusError{Code: http.StatusBadRequest, Err: errors.New("must be POST")}
}

package admin

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/helpers"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Links lists all of the tracked redirects
func Links(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["title"] = "Links"
	data["section"] = "links"

	links := []models.Link{}
	result := env.DB.Preload(clause.Associations).Find(&links)
	if result.RowsAffected > 0 {
		data["links"] = links
	}

	return render(w, data, "links/index.html")
}

// DeleteLink removes the given link from the database
func DeleteLink(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		r.ParseForm()
		code := r.PostFormValue("link")
		result := env.DB.Where("Code = ?", code).Delete(&models.Geosite{})
		if result.RowsAffected != 0 {
			flash.Set(w, r, flash.Message{Message: "Linked deleted successfully", Style: "success"})
			http.Redirect(w, r, helpers.URL("admin/links"), http.StatusFound)
		} else {
			flash.Set(w, r, flash.Message{Message: "Something went wrong and the link wasn't deleted. Check if the link was already deleted.", Style: "warning"})
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
	} else {
		flash.Set(w, r, flash.Message{Message: "Invalid request", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	}
	return nil
}

// RestoreLink will undo soft delete on the last deleted link
func RestoreLink(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	var link = models.Link{}
	env.DB.Unscoped().Where("deleted_at IS NOT NULL").Order("deleted_at DESC").Limit(1).Find(&link)
	result := env.DB.Model(&link).Updates(map[string]interface{}{"deleted_at": nil})
	if result.RowsAffected == 0 {
		flash.Set(w, r, flash.Message{Message: "Could not restore link", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	} else {
		flash.Set(w, r, flash.Message{Message: "Link restored successfully", Style: "success"})
		http.Redirect(w, r, helpers.URL("admin/links/edit/"+link.Code), http.StatusFound)
	}
	return nil
}

// EditLink allows the user to view and edit the given link
func EditLink(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})

	code := chi.URLParam(r, "code")
	link := models.Link{}
	result := env.DB.Where("code = ?", code).Preload(clause.Associations).Find(&link)

	if result.RowsAffected == 0 {
		return handler.StatusError{Code: http.StatusNotFound, Err: errors.New("link cannot be found")}
	}

	if r.Method == http.MethodPost || r.Method == http.MethodPatch {
		r.ParseForm()
		if _, ok := r.PostForm["delete"]; ok {
			result = env.DB.Delete(&link)
			if result.RowsAffected == 0 {
				flash.Set(w, r, flash.Message{Message: "Could not delete the link. It may have already been deleted.", Style: "danger"})
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			} else {
				flash.Set(w, r, flash.Message{Message: "Link deleted. <a href=\"" + helpers.URL("admin/links/restore") + "\">Undo</a>", Style: "success"})
				http.Redirect(w, r, helpers.URL("admin/links"), http.StatusFound)
			}
			return nil
		}

		if val, ok := r.PostForm["title"]; ok {
			link.Title = val[0]
		}
		if val, ok := r.PostForm["link"]; ok {
			link.URL = val[0]
		}
		if val, ok := r.PostForm["author"]; ok {
			link.Author = val[0]
		}
		if _, ok := r.PostForm["publish"]; ok {
			link.Published = true
		} else {
			link.Published = false
		}
		result = env.DB.Save(&link)
		if result.RowsAffected == 0 {
			if r.Method == http.MethodPost {
				flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			}
			http.Error(w, "could not save link", http.StatusBadGateway)
			return nil
		}
		if r.Method == http.MethodPost {
			flash.Set(w, r, flash.Message{Message: "Saved!", Style: "success"})
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		}
		return nil
	}

	data["page"] = link
	data["title"] = "Editing " + link.Title
	data["section"] = "links"

	data["messages"] = flash.Get(w, r)

	return render(w, data, "links/edit.html")

}

// CreateLink shows the form, and stores a link in the database
func CreateLink(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["title"] = "Create a Page"
	data["section"] = "links"

	if r.Method == http.MethodPost {
		r.ParseForm()
		link := models.Link{}
		link.Title = r.FormValue("title")
		link.URL = r.FormValue("link")
		for {
			link.Code = helpers.NewCode(5)
			check := models.Link{}
			env.DB.Model(models.Link{}).Where("code = ?", link.Code).Find(&check)
			if check.Code != link.Code {
				break
			}
		}

		result := env.DB.Model(&models.Link{}).Create(&link)
		if result.Error != nil {
			flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
			return nil
		}

		flash.Set(w, r, flash.Message{Message: "Created link!", Style: "success"})
		http.Redirect(w, r, helpers.URL("admin/links/edit/"+link.Code), http.StatusFound)
		return nil
	}

	return render(w, data, "links/create.html")

}

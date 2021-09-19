package admin

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
	"gorm.io/gorm/clause"
)

// Pages allows admin to add content
func Pages(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["title"] = "Pages"

	pages := []models.Page{}
	result := env.DB.Order("created_at desc").Preload(clause.Associations).Find(&pages)
	if result.RowsAffected > 0 {
		data["pages"] = pages
	}

	return render(w, data, "pages/index.html")
}

// DeletePage removes the given page from the database
func DeletePage(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		r.ParseForm()
		code := r.PostFormValue("page")
		result := env.DB.Where("Code = ?", code).Delete(&models.Page{})
		if result.RowsAffected != 0 {
			flash.Set(w, r, flash.Message{Message: "Deleted page", Style: "success"})
			http.Redirect(w, r, "/admin/pages", 302)
		} else {
			flash.Set(w, r, flash.Message{Message: "Could not delete page", Style: "warning"})
			http.Redirect(w, r, r.Header.Get("Referer"), 302)
		}
	} else {
		flash.Set(w, r, flash.Message{Message: "Invalid request", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}
	return nil
}

// Restore will undo soft delete on the last deleted page
func Restore(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	var page = models.Page{}
	env.DB.Unscoped().Where("deleted_at IS NOT NULL").Order("deleted_at DESC").Limit(1).Find(&page)
	result := env.DB.Model(&page).Updates(map[string]interface{}{"deleted_at": nil})
	if result.RowsAffected == 0 {
		flash.Set(w, r, flash.Message{Message: "Could not delete page", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	} else {
		flash.Set(w, r, flash.Message{Message: "Page restored", Style: "success"})
		http.Redirect(w, r, "/admin/pages/edit/"+page.Code, 302)
	}
	return nil
}

// EditPage allows the user to view and edit the given page
func EditPage(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})

	code := chi.URLParam(r, "code")
	page := models.Page{}
	result := env.DB.Where("code = ?", code).Preload(clause.Associations).Find(&page)

	if result.RowsAffected == 0 {
		return handler.StatusError{Code: 404, Err: errors.New("page cannot be found")}
	}

	if r.Method == http.MethodPost || r.Method == http.MethodPatch {
		r.ParseForm()
		if _, ok := r.PostForm["delete"]; ok {
			result = env.DB.Delete(&page)
			if result.RowsAffected == 0 {
				flash.Set(w, r, flash.Message{Message: "Could not delete page", Style: "danger"})
				http.Redirect(w, r, r.Header.Get("Referer"), 302)
			} else {
				flash.Set(w, r, flash.Message{Message: "Page deleted. <a href=\"/admin/pages/restore\">Undo</a>", Style: "success"})
				http.Redirect(w, r, "/admin/pages", 302)
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
		if val, ok := r.PostForm["trail"]; ok {
			if v, err := strconv.ParseInt(val[0], 10, 32); err == nil {
				trail := models.Trail{}
				env.DB.Where("id = ?", v).Find(&trail)
				page.Trail = trail
			}
		}
		if val, ok := r.PostForm["gallery"]; ok {
			if v, err := strconv.ParseInt(val[0], 10, 32); err == nil {
				gallery := models.Gallery{}
				env.DB.Where("id = ?", v).Find(&gallery)
				page.Gallery = gallery
			}
		}
		result = env.DB.Save(&page)
		if result.RowsAffected == 0 {
			if r.Method == http.MethodPost {
				flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
				http.Redirect(w, r, r.Header.Get("Referer"), 302)
			}
			http.Error(w, "could not save page", http.StatusBadGateway)
			return nil
		}
		if r.Method == http.MethodPost {
			flash.Set(w, r, flash.Message{Message: "Saved!", Style: "success"})
			http.Redirect(w, r, r.Header.Get("Referer"), 302)
		}
		return nil
	}

	trails := []models.Trail{}
	env.DB.Find(&trails)
	data["trails"] = trails

	galleries := []models.Gallery{}
	env.DB.Find(&galleries)
	data["galleries"] = galleries

	data["page"] = page
	data["title"] = "Editing " + page.Title + " | Admin"

	data["messages"] = flash.Get(w, r)

	return render(w, data, "pages/edit.html")

}

// CreatePage removes the given page from the database
func CreatePage(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["title"] = "Create a Page | Admin"

	if r.Method == http.MethodPost {
		r.ParseForm()
		page := models.Page{}
		page.Title = r.FormValue("title")
		page.Text = r.FormValue("content")
		result := env.DB.Model(&models.Page{}).Create(&page)
		if result.RowsAffected == 0 {
			flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
			http.Redirect(w, r, r.Header.Get("Referer"), 302)
			return nil
		}
		flash.Set(w, r, flash.Message{Message: "Created page!", Style: "success"})
		http.Redirect(w, r, "/admin/pages/edit/"+page.Code, 302)
		return nil
	}

	return render(w, data, "pages/create.html")

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
	return errors.New("This is not a POST request")
}

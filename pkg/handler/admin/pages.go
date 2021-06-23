package admin

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/AmazingTrace/pkg/flash"
	"github.com/nathanhollows/AmazingTrace/pkg/handler"
	"github.com/nathanhollows/AmazingTrace/pkg/models"
)

// Pages allows admin to add content
func Pages(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	data := make(map[string]interface{})
	data["title"] = "Codes | Admin"

	pages := []models.Page{}
	result := env.DB.Order("created_at desc").Find(&pages)
	if result.RowsAffected > 0 {
		data["pages"] = pages
	}

	data["messages"] = flash.Get(w, r)

	templates := template.Must(template.ParseFiles(
		"../web/templates/admin.html",
		"../web/templates/flash.html",
		"../web/views/admin/pages/index.html"))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

// ImportPages populates the database with dummy pages
func ImportPages(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	results := env.DB.Find(&models.Page{})
	if results.RowsAffected > 0 {
		flash.Set(w, r, flash.Message{Message: "Pages already exist. Cannot import dummy data.", Style: "danger"})
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
		return nil
	}
	page := []models.Page{
		{
			Title:  "Test Page",
			Text:   "Lorem ipsum",
			Author: "Nathan Hollows",
		},
		{
			Title:  "Another Page",
			Text:   "Can you believe it? Even more lorem ipsum. Well, hardly.",
			Author: "Nathan Hollows",
		},
	}
	result := env.DB.Model(&models.Page{}).Create(&page)
	if result.RowsAffected > 0 {
		flash.Set(w, r, flash.Message{Message: "Successfully imported 2 dummy pages", Style: "success"})
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	} else {
		flash.Set(w, r, flash.Message{Message: "Something went wrong", Style: "warning"})
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
	}

	return nil
}

// DeletePage removes the given page from the database
func DeletePage(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		r.ParseForm()
		code := r.PostFormValue("page")
		result := env.DB.Where("Code = ?", code).Delete(&models.Page{})
		if result.RowsAffected > 0 {
			flash.Set(w, r, flash.Message{Message: "Deleted page", Style: "success"})
			http.Redirect(w, r, r.Header.Get("Referer"), 302)
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

// EditPage allows the user to view and edit the given page
func EditPage(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	code := chi.URLParam(r, "code")
	page := models.Page{}
	result := env.DB.Where("Code = ?", code).Find(&page)

	if result.RowsAffected == 0 {
		return handler.StatusError{Code: 404, Err: errors.New("page cannot be found")}
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		page.Title = r.FormValue("title")
		page.Text = r.FormValue("content")
		result = env.DB.Save(&page)
		if result.RowsAffected == 0 {
			flash.Set(w, r, flash.Message{Message: "Could not save", Style: "danger"})
			http.Redirect(w, r, r.Header.Get("Referer"), 302)
			return nil
		}
		flash.Set(w, r, flash.Message{Message: "Saved!", Style: "success"})
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
		return nil
	}

	data := make(map[string]interface{})
	data["page"] = page
	data["title"] = "Editing " + page.Title + " | Admin"
	data["messages"] = flash.Get(w, r)

	templates := template.Must(template.ParseFiles(
		"../web/templates/admin.html",
		"../web/templates/flash.html",
		"../web/views/admin/pages/edit.html"))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

// CreatePage removes the given page from the database
func CreatePage(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
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

	data := make(map[string]interface{})
	data["title"] = "Create a Page | Admin"
	data["messages"] = flash.Get(w, r)

	templates := template.Must(template.ParseFiles(
		"../web/templates/admin.html",
		"../web/templates/flash.html",
		"../web/views/admin/pages/create.html"))

	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return nil
}

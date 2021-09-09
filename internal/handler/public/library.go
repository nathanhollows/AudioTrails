package public

import (
	"net/http"

	"github.com/nathanhollows/Argon/internal/flash"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
)

// Library shows the user all the unlocked pages they have found
func Library(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)

	pages := []models.Page{}
	env.DB.Where("published = true").Find(&pages)
	data["pages"] = pages

	library := []models.Library{}
	env.DB.Find(&library)
	data["library"] = library

	var galleries []map[string]interface{}
	env.DB.Table("galleries").Joins("left join pages on pages.gallery_id = galleries.id").Where("published = ?", true).Select("gallery, count(*) as total").Group("gallery_id").Find(&galleries)
	data["galleries"] = galleries

	var trails []map[string]interface{}
	env.DB.Table("trails").
		Joins("left join pages on pages.trail_id = trails.id").
		Joins("left join galleries on galleries.id = pages.gallery_id").
		Where("published = ?", true).
		Select("gallery, trail, count(*) as total").
		Group("trail_id, gallery").
		Find(&trails)
	data["trails"] = trails

	return render(w, data, "library/index.html")
}

package admin

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/models"
)

// Media manages assests, go figure.
// E.g. images, videos, audio
func Media(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	data := make(map[string]interface{})
	data["title"] = "Media"

	return render(w, data, "media/index.html")
}

// Upload saves audio and images to the server
// Only accepts POST requests
func Upload(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return handler.StatusError{Code: http.StatusMethodNotAllowed, Err: errors.New("Request must be POST")}
	}

	err := r.ParseMultipartForm(20000000)
	if err != nil {
		log.Println(err.Error())
		return handler.StatusError{Code: http.StatusBadRequest, Err: errors.New("The file could not be uploaded")}
	}

	formdata := r.MultipartForm
	files := formdata.File["file"]

	if len(files) == 0 {
		return handler.StatusError{Code: http.StatusBadRequest, Err: errors.New("No files available to read. Try again")}
	}

	for i := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return handler.StatusError{Code: http.StatusBadRequest, Err: errors.New("No files available to read. Try again")}
		}

		filetype := strings.Split(files[i].Header.Get("Content-Type"), "/")[0]
		format := strings.Split(files[i].Header.Get("Content-Type"), "/")[1]
		title := strings.TrimSuffix(files[i].Filename, filepath.Ext(files[i].Filename))
		filename := fmt.Sprint(time.Now().Nanosecond(), "-", files[i].Filename)
		filename = strings.Replace(filename, " ", "-", -1)
		filepath := fmt.Sprint("web/static/uploads/", filetype, "/", filename)

		hash := sha256.New()
		tr := io.TeeReader(file, hash)

		if filetype != "audio" && filetype != "image" {
			return handler.StatusError{Code: http.StatusNotAcceptable, Err: errors.New("No files available to read. Try again")}
		}

		out, err := os.Create(filepath)

		defer out.Close()
		if err != nil {
			return handler.StatusError{Code: http.StatusInternalServerError, Err: errors.New("Unable to create the file for writing. Check your write access privilege")}
		}

		_, err = io.Copy(out, tr) // file not files[i] !

		if err != nil {
			return handler.StatusError{Code: http.StatusInternalServerError, Err: errors.New("Unable to write the file. Something went wrong")}
		}

		filehash := fmt.Sprintf("%x", hash.Sum(nil))

		media := models.Media{}
		env.DB.Model(media).Where("hash = ?", filehash).Limit(1).Find(&media)
		if media.File == "" {
			media = models.Media{
				Title:  title,
				File:   filename,
				Type:   filetype,
				Format: format,
				Hash:   filehash,
			}
			env.DB.Create(&media)
		} else {
			path := filepath
			os.Remove(path)
		}

		if len(formdata.Value["geosite"]) != 0 {
			geosite := models.Geosite{}
			env.DB.Where("code = ?", formdata.Value["geosite"][0]).Find(&geosite)
			if geosite.Code != "" {
				if filetype == "audio" {
					geosite.Audio = media
				} else {
					geosite.Cover = media
				}
				env.DB.Updates(&geosite)
			}
		}

		if filetype == "image" {
			cmd := exec.Command("convert", filename, "-resize", "576x576^", "-quality", "90", "-define", "png:compression-filter=5", "small/"+filename)
			cmd.Dir = "./web/static/uploads/image/"
			cmd.Run()
			cmd = exec.Command("convert", filename, "-resize", "1200x300^", "-quality", "90", "-define", "png:compression-filter=5", "medium/"+filename)
			cmd.Dir = "./web/static/uploads/image/"
			cmd.Run()
			cmd = exec.Command("convert", filename, "-resize", "2000x700^", "-quality", "90", "-define", "png:compression-filter=5", "large/"+filename)
			cmd.Dir = "./web/static/uploads/image/"
			cmd.Run()

			fmt.Fprintf(w, media.ImgURL("small"))
			return nil
		}

		fmt.Fprintf(w, media.URL())
	}
	return nil
}

package public

import (
	"net/http"

	"github.com/aaronarduino/goqrsvg"
	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
	"github.com/go-chi/chi"
	"github.com/nathanhollows/Argon/internal/handler"
	"github.com/nathanhollows/Argon/internal/helpers"
)

// QR returns an SVG qr code
func QR(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "max-age=36000;")
	code := chi.URLParam(r, "code")
	location := chi.URLParam(r, "location")

	svgQR(helpers.URL(location+"/"+code), w)

	return nil
}

func svgQR(link string, w http.ResponseWriter) {
	s := svg.New(w)
	qrCode, _ := qr.Encode(link, qr.Q, qr.Auto)

	qs := goqrsvg.NewQrSVG(qrCode, 20)
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)

	s.End()
}

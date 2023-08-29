package greeting

import (
	"bytes"
	"image/png"
	"net/http"

	"github.com/aofei/cameron"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}
	png.Encode(&buf, cameron.Identicon([]byte(r.RequestURI), 540, 60))
	w.Header().Set("Content-Type", "image/png")
	buf.WriteTo(w)
	// fmt.Fprint(w, "Hello World")
}

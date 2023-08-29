package greeting

import (
	"bytes"
	"fmt"
	"image/png"
	"net/http"

	"github.com/aofei/cameron"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}
	png.Encode(&buf, cameron.Identicon([]byte(r.RequestURI), 540, 60))
	w.Header().Set("Content-Type", "image/png")
	buf.WriteTo(w)

	// if you just want to print the following text, you can use the following code
	// fmt.Fprint(w, greet("Jenny")
}

// private function, have a think about how to test it
func greet(name string) string {
	return fmt.Sprintf("Hello World, %s", name)
}

// public function
func Greet(name string) string {
	return fmt.Sprintf("Hello World, %s", name)
}

package main 

import (
	"net/http"
	"io"
	"os"
	"log"
	"strings"
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		// comment these , you'd get texts, no routing..
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"		
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}
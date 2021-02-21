package renderer

import (
	"net/http"
)

func LoadServer(dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := http.FileServer(http.Dir(dir))
		h.ServeHTTP(w, r)
	})
}

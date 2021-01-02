package renderer

import (
	"net/http"
)

func LoadServer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := http.FileServer(http.Dir("./renderer/assets"))
		h.ServeHTTP(w, r)
	})
}

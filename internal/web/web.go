package web

import (
	"net/http"
	"os"
	"path"

	"ecksbee.com/telefacts/internal/cache"
	"github.com/gorilla/mux"
)

func Catalog() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]
		if len(id) <= 0 {
			http.Error(w, "Error: invalid id '"+id+"'", http.StatusBadRequest)
			return
		}
		workingDir := path.Join(".", "projects", id)
		_, err := os.Stat(workingDir)
		if os.IsNotExist(err) {
			http.Error(w, "Error: "+err.Error(), http.StatusNotFound)
			return
		}
		data, err := cache.MarshalCatalog(workingDir)
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func Renderable() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)
		id := vars["id"]
		if len(id) <= 0 {
			http.Error(w, "Error: invalid id '"+id+"'", http.StatusBadRequest)
			return
		}
		workingDir := path.Join(".", "projects", id)
		_, err := os.Stat(workingDir)
		if os.IsNotExist(err) {
			http.Error(w, "Error: "+err.Error(), http.StatusNotFound)
			return
		}
		hash := vars["hash"]
		if len(hash) <= 0 {
			http.Error(w, "Error: invalid roote", http.StatusBadRequest)
			return
		}

		data, err := cache.Marshal(workingDir, hash)
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/renderables", Catalog()).Methods("GET")
	renderablesRoute := r.PathPrefix("/renderables").Subrouter()
	projectIDPrefix := renderablesRoute.PathPrefix("/{id}")
	projectIDRoute := projectIDPrefix.Subrouter()
	projectIDRoute.HandleFunc("/{hash}", Renderable()).Methods("GET")
	r.Handle("/", http.FileServer(http.Dir((path.Join(".", "renderer")))))

	//todo serve open api spec at root
	return r
}

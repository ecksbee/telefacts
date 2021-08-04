package web

import (
	"net/http"
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
		data, err := cache.MarshalCatalog(id)
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
		hash := vars["hash"]
		if len(hash) <= 0 {
			http.Error(w, "Error: invalid roote", http.StatusBadRequest)
			return
		}

		data, err := cache.Marshal(id, hash)
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
	foldersRoute := r.PathPrefix("/folders").Subrouter()
	foldersRoute.HandleFunc("/{id}", Catalog()).Methods("GET")
	projectIDRoute := foldersRoute.PathPrefix("/{id}").Subrouter()
	projectIDRoute.HandleFunc("/{hash}", Renderable()).Methods("GET")
	r.Handle("/", http.FileServer(http.Dir((path.Join(".", "renderer")))))

	//todo serve open api spec at root
	return r
}

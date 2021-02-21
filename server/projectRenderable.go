package server

import (
	"net/http"
	"os"
	"path"

	"ecksbee.com/telefacts/sec"
	"github.com/gorilla/mux"
)

func getProjectRenderable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error: incorrect verb", http.StatusInternalServerError)
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
	slug := vars["slug"]
	if len(slug) <= 0 {
		http.Error(w, "Error: invalid roote", http.StatusBadRequest)
		return
	}
	//todo check underscore and determine if sec
	data, err := sec.Marshal(workingDir, slug)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func postProjectRenderable(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error: incorrect verb", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) <= 0 {
		http.Error(w, "Error: invalid id '"+id+"'", http.StatusBadRequest)
		return
	}
	r.ParseMultipartForm(32 << 20)
	workingDir := path.Join(".", "projects", id)
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		http.Error(w, "Error: "+err.Error(), http.StatusNotFound)
		return
	}
	slug := vars["slug"]
	if len(slug) <= 0 {
		http.Error(w, "Error: invalid network ", http.StatusBadRequest)
		return
	}
	//todo check underscore and determine if sec
	//todo clear sec cache
	data, err := sec.Marshal(workingDir, slug)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ProjectRenderable() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProjectRenderable(w, r)
		} else if r.Method == http.MethodPost {
			postProjectRenderable(w, r)
		} else {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
		}
	}
}

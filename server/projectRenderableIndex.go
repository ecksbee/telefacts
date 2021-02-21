package server

import (
	"net/http"
	"os"
	"path"

	"ecksbee.com/telefacts/sec"
	"github.com/gorilla/mux"
)

func getProjectRenderableIndex(w http.ResponseWriter, r *http.Request) {
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
	//todo check underscore and determine if sec
	data, err := sec.MarshalCatalog(workingDir)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func postProjectRenderableIndex(w http.ResponseWriter, r *http.Request) {
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
	//todo check underscore and determine if sec
	//todo clear sec cache
	data, err := sec.MarshalCatalog(workingDir)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ProjectRenderableIndex() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProjectRenderableIndex(w, r)
		} else if r.Method == http.MethodPost {
			postProjectRenderableIndex(w, r)
		} else {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
		}
	}
}

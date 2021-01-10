package server

import (
	"net/http"
	"os"
	"path"
	"strconv"

	"ecks-bee.com/telefacts/sec"
	"github.com/gorilla/mux"
	gocache "github.com/patrickmn/go-cache"
)

func getProjectRenderable(cache *gocache.Cache, w http.ResponseWriter, r *http.Request) {
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
	l := vars["l"]
	if len(l) <= 0 {
		http.Error(w, "Error: invalid network '"+l+"'", http.StatusBadRequest)
		return
	}
	i, err := strconv.Atoi(vars["i"])
	if err != nil {
		http.Error(w, "Error: invalid entity index '"+vars["i"]+"'", http.StatusBadRequest)
		return
	}
	j, err := strconv.Atoi(vars["j"])
	if err != nil {
		http.Error(w, "Error: invalid relationship set index '"+vars["j"]+"'", http.StatusBadRequest)
		return
	}
	//todo check underscore and determine if sec
	secProject := sec.SECProject{
		ID:       id,
		AppCache: cache,
	}
	data, err := secProject.RenderDataGrid(workingDir, l, i, j)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func postProjectRenderable(cache *gocache.Cache, w http.ResponseWriter, r *http.Request) {
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
	l := vars["l"]
	if len(l) <= 0 {
		http.Error(w, "Error: invalid network '"+l+"'", http.StatusBadRequest)
		return
	}
	i, err := strconv.Atoi(vars["i"])
	if err != nil {
		http.Error(w, "Error: invalid entity index '"+vars["i"]+"'", http.StatusBadRequest)
		return
	}
	j, err := strconv.Atoi(vars["j"])
	if err != nil {
		http.Error(w, "Error: invalid relationship set index '"+vars["j"]+"'", http.StatusBadRequest)
		return
	}
	//todo import taxonomies
	//todo check underscore and determine if sec
	secProject := sec.SECProject{
		ID:       id,
		AppCache: cache,
	}
	data, err := secProject.RenderDataGrid(workingDir, l, i, j)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ProjectRenderable(cache *gocache.Cache) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProjectRenderable(cache, w, r)
		} else if r.Method == http.MethodPost {
			postProjectRenderable(cache, w, r)
		} else {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
		}
	}
}

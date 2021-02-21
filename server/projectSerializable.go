package server

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"ecksbee.com/telefacts/actions"
	"ecksbee.com/telefacts/sec"
	"github.com/gorilla/mux"
)

func getProjectSerializable(w http.ResponseWriter, r *http.Request) {
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
	data, err := actions.Download(workingDir)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", id))
	w.Write(data)
}

func postProjectSerializable(w http.ResponseWriter, r *http.Request) {
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
	zipFile, handler, err := r.FormFile("zip")
	defer zipFile.Close()
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME header: %+v\n", handler.Header)
	//todo check underscore and determine if sec
	err = sec.Upload(zipFile, *handler, workingDir, true)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ProjectSerializable() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProjectSerializable(w, r)
		} else if r.Method == http.MethodPost {
			postProjectSerializable(w, r)
		} else {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
		}
	}
}

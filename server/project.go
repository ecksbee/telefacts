package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"ecksbee.com/telefacts/sec"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error: incorrect verb", http.StatusInternalServerError)
		return
	}
	workingDir := path.Join(".", "projects")
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) > 0 {
		pathStr := path.Join(workingDir, id)
		_, err := os.Stat(pathStr)
		if os.IsNotExist(err) {
			http.Error(w, "Error: "+err.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "<div><h1>todo show list of documents for %s</h1></div>", id) //todo show list of documents
		return
	}
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		http.Error(w, "Error: incorrect verb", http.StatusInternalServerError)
		return
	}
	list := ""
	for _, f := range files {
		if f.IsDir() {
			//todo show list of project ids and there metadata
			list = list + f.Name() + ","
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, list)
}

func postProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error: incorrect verb", http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	workingDir := path.Join(".", "projects")
	id := uuid.New()
	pathStr := path.Join(workingDir, id.String())
	_, err := os.Stat(pathStr)
	for os.IsExist(err) {
		id = uuid.New()
		pathStr = path.Join(workingDir, id.String())
		_, err = os.Stat(pathStr)
	}
	err = os.Mkdir(pathStr, 0755)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	filingURL := r.FormValue("SEC")
	if len(filingURL) > 0 {
		log.Printf("Scraping SEC")
		err = sec.Import(filingURL, pathStr)
		if err != nil {
			log.Printf("SEC scraping error: %v+\n", err)
		}
		meta := path.Join(workingDir, "_")
		file, _ := os.OpenFile(meta, os.O_CREATE, 0755)
		defer file.Close()
		encoder := json.NewEncoder(file)
		go encoder.Encode(struct {
			SEC string
		}{
			SEC: filingURL,
		})
		go sec.Hydratable(id.String())
	}
	fmt.Fprintf(w, id.String())
}

func Project() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProject(w, r)
		} else if r.Method == http.MethodPost {
			postProject(w, r)
		} else {
			http.Error(w, "Error: incorrect verb, "+r.Method, http.StatusInternalServerError)
		}
	}
}

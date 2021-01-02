// +build js
// +build wasm

package renderer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"ecks-bee.com/telefacts/renderables"
)

func fetchCatalog() (*renderables.Catalog, error) {
	targetURL := currentURL.Scheme + "://" + currentURL.Host + "/projects/" + id + "/renderables"
	consoleLog("fetching from " + targetURL)
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var catalog renderables.Catalog
	err = json.Unmarshal(bodyBytes, &catalog)
	return &catalog, err
}

func fetchPGrid(entityIndex int, relationshipSetIndex int) (*renderables.PGrid, error) {
	i := strconv.Itoa(entityIndex)
	j := strconv.Itoa(relationshipSetIndex)
	targetURL := currentURL.Scheme + "://" + currentURL.Host + "/projects/" + id + "/renderables" +
		"/pre/" + i + "/" + j
	consoleLog("fetching from " + targetURL)
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var pGrid renderables.PGrid
	err = json.Unmarshal(bodyBytes, &pGrid)
	return &pGrid, err
}

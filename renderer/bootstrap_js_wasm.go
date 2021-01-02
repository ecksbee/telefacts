// +build js
// +build wasm

package renderer

import (
	"fmt"
	"net/url"
)

var (
	currentURL *url.URL
	id         string
)

func Run() {
	fmt.Println("bootstrapping TeleFacts renderer")
	var err error
	err = initializeBrowser()
	if err != nil {
		return
	}

	connectView()
	bindController()
	initializeRenderer()
	<-make(chan bool)
}

func initializeRenderer() error {
	setIsLoading(true)
	defer setIsLoading(false)
	catalog, err := fetchCatalog()
	if err != nil {
		msg := "cannot connect to the view"
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	setEntities(catalog.Entities)
	setRelationshipSets(catalog.RelationshipSets)
	return nil
}

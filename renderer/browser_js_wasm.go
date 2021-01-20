// +build js
// +build wasm

package renderer

import (
	"fmt"
	"net/url"
	"strconv"
	"syscall/js"
)

func initializeBrowser() error {
	navigator = js.Global().Get("navigator")
	if !navigator.Truthy() {
		return fmt.Errorf("Unable to get navigator object")
	}
	document = js.Global().Get("document")
	if !document.Truthy() {
		return fmt.Errorf("Unable to get document object")
	}
	console = js.Global().Get("console")
	if !console.Truthy() {
		return fmt.Errorf("Unable to get console object")
	}
	var deviceMemory = navigator.Get("deviceMemory")
	if !deviceMemory.Truthy() {
		msg := "Cannot bootstrap renderer to an incompatible browser"
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	ram := deviceMemory.Int()
	if ram < 2 {
		msg := "Cannot bootstrap renderer in a device with less than 2 gigabytes of memory: " + strconv.Itoa(ram)
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	//todo check screen size if its a small screens and touchscreen
	var err error
	href := js.Global().Get("location").Get("href").String()
	currentURL, err = url.Parse(href)
	if err != nil {
		msg := err.Error()
		alert(msg)
		consoleError(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	queries := currentURL.Query()
	idQueries := queries["id"]
	if len(idQueries) <= 0 {
		msg := "Invalid query string, id"
		alert(msg)
		consoleError(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	id = idQueries[0]
	if len(id) <= 0 {
		msg := "Invalid query id"
		alert(msg)
		consoleError(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	return nil
}

func alert(text string) {
	js.Global().Call("alert", text)
}

var navigator js.Value

var document js.Value

var console js.Value

func consoleLog(text string) {
	if !console.Truthy() {
		fmt.Println("Log: " + text)
		return
	}
	console.Call("log", text)
}

func consoleError(text string) {
	msg := "TeleFacts error: " + text
	if !console.Truthy() {
		fmt.Println(msg)
		return
	}
	console.Call("error", msg)
}

// +build js
// +build wasm

package main

import (
	"ecksbee.com/telefacts/renderer"
)

// First, make sure the javascript glue code is served with "cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./renderer/assets"
// Then compile with "GOOS=js GOARCH=wasm go build -o ./renderer/assets/renderer.wasm ./renderer/wasm/main.go"
func main() {
	renderer.Run()
}

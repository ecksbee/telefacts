// +build js
// +build wasm

package renderer

import (
	"fmt"
	"syscall/js"
)

func bindController() {
	fmt.Println("binding to controller")
	view.Set("changeLanguage", changeLanguage())
	view.Set("changeNetwork", changeNetwork())
	view.Set("changeEntity", changeEntity())
	view.Set("changeRelationshipSet", changeRelationshipSet())
}

func changeLanguage() js.Func {
	var ret js.Func
	ret = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if isLoading() {
			return ret
		}
		v := args[0].String()
		consoleLog(v)
		//todo
		return ret
	})
	return ret
}

func changeNetwork() js.Func {
	var ret js.Func
	ret = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if isLoading() {
			return ret
		}
		e := args[0].String()
		selectNetwork(e)
		if selectedRelationshipSet() != "" && selectedEntity() != "" {
			go func() {
				refreshDataGrid()
			}()
		}
		return ret
	})
	return ret
}

func changeEntity() js.Func {
	var ret js.Func
	ret = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if isLoading() {
			return ret
		}
		e := args[0].String()
		selectEntity(e)
		if selectedRelationshipSet() != "" {
			go func() {
				refreshDataGrid()
			}()
		}
		return ret
	})
	return ret
}

func changeRelationshipSet() js.Func {
	var ret js.Func
	ret = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if isLoading() {
			return ret
		}
		rset := args[0].String()
		selectRelationshipSet(rset)
		if selectedEntity() != "" {
			go func() {
				refreshDataGrid()
			}()
		}
		return ret
	})
	return ret
}

func refreshDataGrid() error {
	network := selectedNetwork()
	switch network {
	case "pre":
	case "cal":
		break
	default:
		return fmt.Errorf("invalid concept network")
	}
	if isLoading() {
		return nil
	}
	setIsLoading(true)
	defer setIsLoading(false)
	catalog, err := fetchCatalog() //todo
	if err != nil {
		msg := "cannot fetch catalog"
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	i := func() int {
		value := selectedEntity()
		slice := catalog.Entities
		for p, v := range slice {
			if v == value {
				return p
			}
		}
		return -1
	}()
	if i < 0 {
		msg := "cannot select entity " + selectedEntity()
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	j := func() int {
		value := selectedRelationshipSet()
		slice := catalog.RelationshipSets
		for p, v := range slice {
			if v == value {
				return p
			}
		}
		return -1
	}()
	if j < 0 {
		msg := "cannot select relationship set " + selectedRelationshipSet()
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	switch network {
	case "pre":
		pGrid, err := fetchPGrid(i, j)
		if err != nil {
			msg := "cannot fetch pGrid"
			consoleError(msg)
			alert(msg)
			//todo render error UI
			return fmt.Errorf(msg)
		}
		setPGrid(pGrid)
		setPGrid(nil)
		return nil
		break
	case "cal":
		cGrid, err := fetchCGrid(i, j)
		if err != nil {
			msg := "cannot fetch cGrid"
			consoleError(msg)
			alert(msg)
			//todo render error UI
			return fmt.Errorf(msg)
		}
		setCGrid(cGrid)
		setPGrid(nil)
		return nil
		break
	default:
		return fmt.Errorf("invalid concept network")
	}
	return fmt.Errorf("data grid refresh failed")
}

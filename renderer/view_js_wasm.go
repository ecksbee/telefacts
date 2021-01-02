// +build js
// +build wasm

package renderer

import (
	"fmt"
	"syscall/js"

	"ecks-bee.com/telefacts/renderables"
)

var (
	view js.Value
)

func connectView() error {
	fmt.Println("connecting to view")
	view = js.Global().Get("telefacts-renderer")
	if !view.Truthy() {
		msg := "cannot connect to the view"
		consoleError(msg)
		alert(msg)
		//todo render error UI
		return fmt.Errorf(msg)
	}
	return nil
}

func isLoading() bool {
	return view.Get("isLoading").Bool()
}

func setIsLoading(v bool) {
	view.Set("isLoading", v)
}

func selectedEntity() string {
	return view.Get("selectedEntity").String()
}

func selectEntity(v string) {
	view.Set("selectedEntity", v)
}

func entities() []string {
	jsEntities := view.Get("entities")
	var ret []string
	len := jsEntities.Length()
	for i := 0; i < len; i++ {
		entity := jsEntities.Index(i).String()
		ret = append(ret, entity)
	}
	return ret
}

func setEntities(v []string) {
	s := make([]interface{}, len(v))
	for i, t := range v {
		s[i] = t
	}
	view.Set("entities", s)
}

func selectedRelationshipSet() string {
	return view.Get("selectedRelationshipSet").String()
}

func selectRelationshipSet(v string) {
	view.Set("selectedRelationshipSet", v)
}

func relationshipSets() []string {
	jsRSets := view.Get("relationshipSets")
	var ret []string
	len := jsRSets.Length()
	for i := 0; i < len; i++ {
		rset := jsRSets.Index(i).String()
		ret = append(ret, rset)
	}
	return ret
}

func setRelationshipSets(v []string) {
	s := make([]interface{}, len(v))
	for i, t := range v {
		s[i] = t
	}
	view.Set("relationshipSets", s)
}

func selectedPGrid() *renderables.PGrid {
	var ret renderables.PGrid
	jsVal := view.Get("pGrid")
	indentedLabels := jsVal.Get("IndentedLabels")
	ilLen := indentedLabels.Length()
	for i := 0; i < ilLen; i++ {
		jsIL := indentedLabels.Index(i)
		ret.IndentedLabels = append(ret.IndentedLabels, renderables.IndentedLabel{
			Href:        jsIL.Get("Href").String(),
			Indentation: jsIL.Get("Indentation").Int(),
		})
	}
	ret.MaxIndentation = jsVal.Get("MaxIndentation").Int()
	relevantContexts := jsVal.Get("RelevantContexts")
	rcLen := relevantContexts.Length()
	for i := 0; i < rcLen; i++ {
		jsRC := relevantContexts.Index(i)
		jsRCDM := jsRC.Get("DomainMemberHeaders")
		jsRCDMLen := jsRCDM.Length()
		domainMemberHeaders := make([]string, jsRCDMLen)
		for j := 0; j < jsRCDMLen; j++ {
			jsDM := jsRCDM.Index(j)
			domainMemberHeaders = append(domainMemberHeaders, jsDM.String())
		}
		ret.RelevantContexts = append(ret.RelevantContexts, renderables.RelevantContext{
			ContextRef:          jsRC.Get("ContextRef").String(),
			PeriodHeader:        jsRC.Get("PeriodHeader").String(),
			DomainMemberHeaders: domainMemberHeaders,
		})
	}
	ret.MaxDepth = jsVal.Get("MaxDepth").Int()
	jsFQ := jsVal.Get("FactualQuadrant")
	jsFQLen := jsFQ.Length()
	goFQ := make([][]string, jsFQLen)
	for i := 0; i < jsFQLen; i++ {
		jsFQRow := jsFQ.Index(i)
		jsFQRowLen := jsFQRow.Length()
		goFQ[i] = make([]string, jsFQRowLen)
		for j := 0; j < jsFQRowLen; j++ {
			jsFQCell := jsFQRow.Index(j)
			goFQ[i][j] = jsFQCell.String()
		}
	}
	ret.FactualQuadrant = goFQ
	return &ret
}

func setPGrid(v *renderables.PGrid) {
	jsMap := make(map[string]interface{})
	indentedLabels := make([]interface{}, len(v.IndentedLabels))
	for j, goIL := range v.IndentedLabels {
		jsIL := make(map[string]interface{})
		jsIL["Href"] = goIL.Href
		jsIL["Indentation"] = goIL.Indentation
		indentedLabels[j] = jsIL
	}
	jsMap["IndentedLabels"] = indentedLabels
	jsMap["MaxIndentation"] = v.MaxIndentation
	relevantContexts := make([]interface{}, len(v.RelevantContexts))
	for j, goRC := range v.RelevantContexts {
		jsRC := make(map[string]interface{})
		jsRC["ContextRef"] = goRC.ContextRef
		jsRC["PeriodHeader"] = goRC.PeriodHeader
		jsRCDM := make([]interface{}, len(goRC.DomainMemberHeaders))
		for jj, goDM := range goRC.DomainMemberHeaders {
			jsRCDM[jj] = goDM
		}
		jsRC["DomainMemberHeaders"] = jsRCDM
		relevantContexts[j] = jsRC
	}
	jsMap["RelevantContexts"] = relevantContexts
	jsMap["MaxDepth"] = v.MaxDepth
	factualQuadrant := make([]interface{}, len(v.FactualQuadrant))
	for j, goFQRow := range v.FactualQuadrant {
		jsFQRow := make([]interface{}, len(goFQRow))
		for jj, fqCell := range goFQRow {
			jsFQRow[jj] = fqCell
		}
		factualQuadrant[j] = jsFQRow
	}
	jsMap["FactualQuadrant"] = factualQuadrant
	view.Set("pGrid", jsMap)
}

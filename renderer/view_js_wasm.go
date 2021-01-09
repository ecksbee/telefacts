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

func selectedNetwork() string {
	return view.Get("selectedNetwork").String()
}

func selectNetwork(v string) {
	view.Set("selectedNetwork", v)
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
	ret.RelevantContexts = convertJSRelevantContexts(jsVal.Get("RelevantContexts"))
	ret.FactualQuadrant = convertJSFactualQuadrant(jsVal.Get("FactualQuadrant"))
	ret.MaxDepth = jsVal.Get("MaxDepth").Int()
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
	jsMap["RelevantContexts"] = convertRelevantContextsToJS(v.RelevantContexts)
	jsMap["MaxDepth"] = v.MaxDepth
	jsMap["FactualQuadrant"] = convertFactualQuadrantToJS(v.FactualQuadrant)
	view.Set("pGrid", jsMap)
}
func selectedCGrid() *renderables.CGrid {
	var ret renderables.CGrid
	jsVal := view.Get("cGrid")
	jsSummationItems := jsVal.Get("SummationItems")
	siLen := jsSummationItems.Length()
	for i := 0; i < siLen; i++ {
		jsSI := jsSummationItems.Index(i)
		jsContributingConcepts := jsSI.Get("ContributingConcepts")
		ccLen := jsContributingConcepts.Length()
		contributingConcepts := make([]renderables.ContributingConcept, ccLen)
		for j := 0; j < ccLen; j++ {
			jsCC := jsContributingConcepts.Index(j)
			contributingConcepts[j] = renderables.ContributingConcept{
				Sign:            jsCC.Get("Sign").String(),
				Scale:           jsCC.Get("Scale").String(),
				Href:            jsCC.Get("Href").String(),
				IsSummationItem: jsCC.Get("IsSummationItem").Bool(),
			}
		}
		relevantContexts := convertJSRelevantContexts(jsSI.Get("RelevantContexts"))
		maxDepth := jsSI.Get("MaxDepth").Int()
		factualQuadrant := convertJSFactualQuadrant(jsSI.Get("FactualQuadrant"))
		ret.SummationItems = append(ret.SummationItems, renderables.SummationItem{
			Href:                 jsSI.Get("Href").String(),
			ContributingConcepts: contributingConcepts,
			MaxDepth:             maxDepth,
			RelevantContexts:     relevantContexts,
			FactualQuadrant:      factualQuadrant,
		})
	}
	return &ret
}
func setCGrid(v *renderables.CGrid) {
	jsMap := make(map[string]interface{})
	summationItems := make([]interface{}, len(v.SummationItems))
	for j, goSI := range v.SummationItems {
		jsSI := make(map[string]interface{})
		jsSI["Href"] = goSI.Href
		contributingConcepts := make([]interface{}, len(goSI.ContributingConcepts))
		for j, goCC := range goSI.ContributingConcepts {
			jsCC := make(map[string]interface{})
			jsCC["Href"] = goCC.Href
			jsCC["Scale"] = goCC.Scale
			jsCC["Sign"] = goCC.Sign
			jsCC["IsSummationItem"] = goCC.IsSummationItem
			contributingConcepts[j] = jsCC
		}
		jsSI["ContributingConcepts"] = contributingConcepts
		jsSI["MaxDepth"] = goSI.MaxDepth
		jsSI["RelevantContexts"] = convertRelevantContextsToJS(goSI.RelevantContexts)
		jsSI["FactualQuadrant"] = convertFactualQuadrantToJS(goSI.FactualQuadrant)
		summationItems[j] = jsSI
	}
	jsMap["SummationItems"] = summationItems
	view.Set("cGrid", jsMap)
}
func convertJSRelevantContexts(jsRelevantContexts js.Value) []renderables.RelevantContext {
	var ret []renderables.RelevantContext
	rcLen := jsRelevantContexts.Length()
	ret = make([]renderables.RelevantContext, 0, rcLen)
	for i := 0; i < rcLen; i++ {
		jsRC := jsRelevantContexts.Index(i)
		jsRCDM := jsRC.Get("DomainMemberHeaders")
		jsRCDMLen := jsRCDM.Length()
		domainMemberHeaders := make([]string, jsRCDMLen)
		for j := 0; j < jsRCDMLen; j++ {
			jsDM := jsRCDM.Index(j)
			domainMemberHeaders = append(domainMemberHeaders, jsDM.String())
		}
		ret[i] = renderables.RelevantContext{
			ContextRef:          jsRC.Get("ContextRef").String(),
			PeriodHeader:        jsRC.Get("PeriodHeader").String(),
			DomainMemberHeaders: domainMemberHeaders,
		}
	}
	return ret
}
func convertJSFactualQuadrant(jsFactualQuadrant js.Value) [][]string {
	jsFQLen := jsFactualQuadrant.Length()
	goFQ := make([][]string, jsFQLen)
	for i := 0; i < jsFQLen; i++ {
		jsFQRow := jsFactualQuadrant.Index(i)
		jsFQRowLen := jsFQRow.Length()
		goFQ[i] = make([]string, jsFQRowLen)
		for j := 0; j < jsFQRowLen; j++ {
			jsFQCell := jsFQRow.Index(j)
			goFQ[i][j] = jsFQCell.String()
		}
	}
	return goFQ
}
func convertRelevantContextsToJS(relevantContexts []renderables.RelevantContext) []interface{} {
	ret := make([]interface{}, len(relevantContexts))
	for j, goRC := range relevantContexts {
		jsRC := make(map[string]interface{})
		jsRC["ContextRef"] = goRC.ContextRef
		jsRC["PeriodHeader"] = goRC.PeriodHeader
		jsRCDM := make([]interface{}, len(goRC.DomainMemberHeaders))
		for jj, goDM := range goRC.DomainMemberHeaders {
			jsRCDM[jj] = goDM
		}
		jsRC["DomainMemberHeaders"] = jsRCDM
		ret[j] = jsRC
	}
	return ret
}
func convertFactualQuadrantToJS(factualQuadrant [][]string) []interface{} {
	ret := make([]interface{}, len(factualQuadrant))
	for j, goFQRow := range factualQuadrant {
		jsFQRow := make([]interface{}, len(goFQRow))
		for jj, fqCell := range goFQRow {
			jsFQRow[jj] = fqCell
		}
		ret[j] = jsFQRow
	}
	return ret
}

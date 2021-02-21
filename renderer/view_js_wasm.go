// +build js
// +build wasm

package renderer

import (
	"fmt"
	"syscall/js"

	"ecksbee.com/telefacts/renderables"
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
			Label:       jsIL.Get("Label").String(),
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
	if v == nil {
		view.Set("pGrid", js.Null())
		return
	}
	jsMap := make(map[string]interface{})
	indentedLabels := make([]interface{}, len(v.IndentedLabels))
	for j, goIL := range v.IndentedLabels {
		jsIL := make(map[string]interface{})
		jsIL["Href"] = goIL.Href
		jsIL["Label"] = goIL.Label
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
				Label:           jsCC.Get("Label").String(),
				IsSummationItem: jsCC.Get("IsSummationItem").Bool(),
			}
		}
		relevantContexts := convertJSRelevantContexts(jsSI.Get("RelevantContexts"))
		maxDepth := jsSI.Get("MaxDepth").Int()
		factualQuadrant := convertJSFactualQuadrant(jsSI.Get("FactualQuadrant"))
		ret.SummationItems = append(ret.SummationItems, renderables.SummationItem{
			Href:                 jsSI.Get("Href").String(),
			Label:                jsSI.Get("Label").String(),
			ContributingConcepts: contributingConcepts,
			MaxDepth:             maxDepth,
			RelevantContexts:     relevantContexts,
			FactualQuadrant:      factualQuadrant,
		})
	}
	return &ret
}
func setCGrid(v *renderables.CGrid) {
	if v == nil {
		view.Set("cGrid", js.Null())
		return
	}
	jsMap := make(map[string]interface{})
	summationItems := make([]interface{}, len(v.SummationItems))
	for j, goSI := range v.SummationItems {
		jsSI := make(map[string]interface{})
		jsSI["Href"] = goSI.Href
		jsSI["Label"] = goSI.Label
		contributingConcepts := make([]interface{}, len(goSI.ContributingConcepts))
		for j, goCC := range goSI.ContributingConcepts {
			jsCC := make(map[string]interface{})
			jsCC["Href"] = goCC.Href
			jsCC["Label"] = goCC.Label
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
func selectedDGrid() *renderables.DGrid {
	var ret renderables.DGrid
	jsVal := view.Get("dGrid")
	jsRootDomains := jsVal.Get("RootDomains")
	rdLen := jsRootDomains.Length()
	for i := 0; i < rdLen; i++ {
		jsRD := jsRootDomains.Index(i)
		jsPrimaryItems := jsRD.Get("PrimaryItems")
		piLen := jsPrimaryItems.Length()
		primaryItems := make([]renderables.PrimaryItem, piLen)
		for j := 0; j < piLen; j++ {
			jsPI := jsPrimaryItems.Index(j)
			jsHypercubes := jsPI.Get("Hypercubes")
			hypercubeLen := jsHypercubes.Length()
			hypercubes := make([]renderables.Hypercube, hypercubeLen)
			for k := 0; k < hypercubeLen; k++ {
				jsHypercube := jsHypercubes.Index(k)
				hypercubes[k] = renderables.Hypercube{
					Href:           jsHypercube.Get("Href").String(),
					Label:          jsHypercube.Get("Label").String(),
					IsClosed:       jsHypercube.Get("IsClosed").Bool(),
					ContextElement: jsHypercube.Get("ContextElement").String(),
					IsInclusive:    jsHypercube.Get("IsInclusive").Bool(),
				}
			}
			primaryItems[j] = renderables.PrimaryItem{
				Href:       jsPI.Get("Href").String(),
				Label:      jsPI.Get("Label").String(),
				Level:      jsPI.Get("Level").Int(),
				Hypercubes: hypercubes,
			}
		}
		jsEffectiveDimensions := jsRD.Get("EffectiveDimensions")
		edLen := jsEffectiveDimensions.Length()
		effectiveDimensions := make([]renderables.EffectiveDimension, edLen)
		for j := 0; j < edLen; j++ {
			jsED := jsEffectiveDimensions.Index(j)
			effectiveDimensions[j] = renderables.EffectiveDimension{
				Href:  jsED.Get("Href").String(),
				Label: jsED.Get("Label").String(),
			}
		}
		jsEffectiveDomainGrid := jsRD.Get("EffectiveDomainGrid")
		edgLen := jsEffectiveDomainGrid.Length()
		effectiveDomainGrid := make([][]renderables.EffectiveDomain, edgLen)
		for j := 0; j < edgLen; j++ {
			jsEDGRow := jsEffectiveDomainGrid.Index(j)
			jsEDGLen := jsEDGRow.Length()
			effectiveDomains := make([]renderables.EffectiveDomain, jsEDGLen)
			for k := 0; k < jsEDGLen; k++ {
				jsEDom := jsEDGRow.Index(k)
				jsEDomLen := jsEDom.Length()
				effectiveMembers := make([]renderables.EffectiveMember, jsEDomLen)
				for l := 0; l < jsEDomLen; l++ {
					jsEM := jsEDom.Index(l)
					effectiveMembers[l] = renderables.EffectiveMember{
						Href:            jsEM.Get("Href").String(),
						Label:           jsEM.Get("Label").String(),
						IsDefault:       jsEM.Get("IsDefault").Bool(),
						IsStrikethrough: jsEM.Get("IsStrikethrough").Bool(),
					}
				}
				effectiveDomains[k] = effectiveMembers
			}
			effectiveDomainGrid[j] = effectiveDomains
		}
		relevantContexts := convertJSRelevantContexts(jsRD.Get("RelevantContexts"))
		maxDepth := jsRD.Get("MaxDepth").Int()
		maxLevel := jsRD.Get("MaxLevel").Int()
		factualQuadrant := convertJSFactualQuadrant(jsRD.Get("FactualQuadrant"))
		jsHypercubes := jsRD.Get("Hypercubes")
		hypercubeLen := jsHypercubes.Length()
		hypercubes := make([]renderables.Hypercube, hypercubeLen)
		for k := 0; k < hypercubeLen; k++ {
			jsHypercube := jsHypercubes.Index(k)
			hypercubes[k] = renderables.Hypercube{
				Href:           jsHypercube.Get("Href").String(),
				Label:          jsHypercube.Get("Label").String(),
				IsClosed:       jsHypercube.Get("IsClosed").Bool(),
				ContextElement: jsHypercube.Get("ContextElement").String(),
				IsInclusive:    jsHypercube.Get("IsInclusive").Bool(),
			}
		}
		ret.RootDomains = append(ret.RootDomains, renderables.RootDomain{
			Href:                jsRD.Get("Href").String(),
			Label:               jsRD.Get("Label").String(),
			PrimaryItems:        primaryItems,
			MaxDepth:            maxDepth,
			MaxLevel:            maxLevel,
			RelevantContexts:    relevantContexts,
			FactualQuadrant:     factualQuadrant,
			EffectiveDomainGrid: effectiveDomainGrid,
			EffectiveDimensions: effectiveDimensions,
			Hypercubes:          hypercubes,
		})
	}
	return &ret
}
func setDGrid(v *renderables.DGrid) {
	if v == nil {
		view.Set("dGrid", js.Null())
		return
	}
	jsMap := make(map[string]interface{})
	rootDomains := make([]interface{}, len(v.RootDomains))
	for j, goRD := range v.RootDomains {
		jsRD := make(map[string]interface{})
		jsRD["Href"] = goRD.Href
		jsRD["Label"] = goRD.Label
		primaryItems := make([]interface{}, len(goRD.PrimaryItems))
		for j, goPI := range goRD.PrimaryItems {
			jsPI := make(map[string]interface{})
			jsPI["Href"] = goPI.Href
			jsPI["Label"] = goPI.Label
			jsPI["Level"] = goPI.Level
			hypercubes := make([]interface{}, len(goPI.Hypercubes))
			for k, goHC := range goPI.Hypercubes {
				jsHC := make(map[string]interface{})
				jsHC["Href"] = goHC.Href
				jsHC["Label"] = goHC.Label
				jsHC["ContextElement"] = goHC.ContextElement
				jsHC["IsClosed"] = goHC.IsClosed
				jsHC["IsInclusive"] = goHC.IsInclusive
				hypercubes[k] = jsHC
			}
			jsPI["Hypercubes"] = hypercubes
			primaryItems[j] = jsPI
		}
		jsRD["PrimaryItems"] = primaryItems
		jsRD["MaxDepth"] = goRD.MaxDepth
		jsRD["MaxLevel"] = goRD.MaxLevel
		jsRD["RelevantContexts"] = convertRelevantContextsToJS(goRD.RelevantContexts)
		jsRD["FactualQuadrant"] = convertFactualQuadrantToJS(goRD.FactualQuadrant)
		hypercubes := make([]interface{}, len(goRD.Hypercubes))
		for k, goHC := range goRD.Hypercubes {
			jsHC := make(map[string]interface{})
			jsHC["Href"] = goHC.Href
			jsHC["Label"] = goHC.Label
			jsHC["ContextElement"] = goHC.ContextElement
			jsHC["IsClosed"] = goHC.IsClosed
			jsHC["IsInclusive"] = goHC.IsInclusive
			hypercubes[k] = jsHC
		}
		jsRD["Hypercubes"] = hypercubes
		effectiveDimensions := make([]interface{}, len(goRD.EffectiveDimensions))
		for k, goED := range goRD.EffectiveDimensions {
			jsED := make(map[string]interface{})
			jsED["Href"] = goED.Href
			jsED["Label"] = goED.Label
			effectiveDimensions[k] = jsED
		}
		jsRD["EffectiveDimensions"] = effectiveDimensions
		effectiveDomainGrid := make([]interface{}, len(goRD.EffectiveDomainGrid))
		for i, goEDGRow := range goRD.EffectiveDomainGrid {
			jsEDGRow := make([]interface{}, len(goEDGRow))
			for ii, goEDom := range goEDGRow {
				jsEDom := make([]interface{}, len(goEDom))
				for iii, goEM := range goEDom {
					jsEm := make(map[string]interface{})
					jsEm["Href"] = goEM.Href
					jsEm["Label"] = goEM.Label
					jsEm["IsDefault"] = goEM.IsDefault
					jsEm["IsStrikethrough"] = goEM.IsStrikethrough
					jsEDom[iii] = jsEm
				}
				jsEDGRow[ii] = jsEDom
			}
			effectiveDomainGrid[i] = jsEDGRow
		}
		jsRD["EffectiveDomainGrid"] = effectiveDomainGrid
		rootDomains[j] = jsRD
	}
	jsMap["RootDomains"] = rootDomains
	view.Set("dGrid", jsMap)
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

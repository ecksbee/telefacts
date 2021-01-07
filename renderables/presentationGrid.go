package renderables

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"ecks-bee.com/telefacts/xbrl"
)

type IndentedLabel struct {
	Href        string
	Indentation int
}

type PGrid struct {
	IndentedLabels   []IndentedLabel
	MaxIndentation   int
	RelevantContexts []RelevantContext
	MaxDepth         int
	FactualQuadrant  [][]string
}

func MarshalPGrid(entityIndex int, relationshipSetIndex int, schema *xbrl.Schema,
	instance *xbrl.Instance, presentation *xbrl.PresentationLinkbase,
	factFinder FactFinder) ([]byte, error) {
	schemedEntities := sortedEntities(instance)
	if entityIndex > len(schemedEntities)-1 {
		return nil, fmt.Errorf("invalid entity index")
	}
	linkroleURIs := sortedRelationshipSets(schema)
	if relationshipSetIndex > len(linkroleURIs)-1 {
		return nil, fmt.Errorf("invalid relationship set index")
	}
	linkroleURI := linkroleURIs[relationshipSetIndex]
	schemedEntity := schemedEntities[entityIndex]
	indentedLabels, maxIndentation := getIndentedLabels(linkroleURI, schema, presentation)
	relevantContexts, maxDepth := getPresentationContexts(schemedEntity, instance, schema, indentedLabels)
	factualQuadrant := getPFactualQuadrant(indentedLabels, relevantContexts, factFinder)
	return json.Marshal(PGrid{
		IndentedLabels:   indentedLabels,
		MaxIndentation:   maxIndentation,
		RelevantContexts: relevantContexts,
		MaxDepth:         maxDepth,
		FactualQuadrant:  factualQuadrant,
	})
}

func getIndentedLabels(linkroleURI string, schema *xbrl.Schema, presentation *xbrl.PresentationLinkbase) ([]IndentedLabel, int) {
	var presentationLinks []xbrl.PresentationLink
	for _, roleRef := range presentation.RoleRef {
		if linkroleURI == roleRef.RoleURI {
			presentationLinks = presentation.PresentationLinks
			break
		}
	}
	for _, presentationLink := range presentationLinks {
		if presentationLink.Role == linkroleURI {
			var root locatorNode
			arcs := presentationLink.PresentationArc
			root.Children = make([]*locatorNode, 0, len(arcs))
			sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
			for _, arc := range arcs {
				if arc.Arcrole == xbrl.PresentationArcrole {
					from, _ := find(&root, arc.From)
					if from != nil {
						to, toIndex := find(&root, arc.To)
						if to != nil {
							root.Children[toIndex] = root.Children[len(root.Children)-1]
							root.Children = root.Children[:len(root.Children)-1]
							from.Children = append(from.Children, to)
						} else {
							order, _ := strconv.ParseFloat(arc.Order, 64)
							from.Children = append(from.Children, &locatorNode{
								Locator:  arc.To,
								Order:    order,
								Children: make([]*locatorNode, 0, len(arcs)),
							})
						}
					} else {
						from = &locatorNode{
							Locator:  arc.From,
							Children: make([]*locatorNode, 0, len(arcs)),
						}
						root.Children = append(root.Children, from)
						to, toIndex := find(&root, arc.To)
						if to != nil {
							root.Children[toIndex] = root.Children[len(root.Children)-1]
							root.Children = root.Children[:len(root.Children)-1]
							from.Children = append(from.Children, to)
						} else {
							order, _ := strconv.ParseFloat(arc.Order, 64)
							from.Children = append(from.Children, &locatorNode{
								Locator:  arc.To,
								Order:    order,
								Children: make([]*locatorNode, 0, len(arcs)),
							})
						}
					}
				}
			}
			ret := make([]IndentedLabel, 0, len(arcs))
			maxIndent := 1
			var makeIndents func(node *locatorNode, level int)
			makeIndents = func(node *locatorNode, level int) {
				if len(node.Children) <= 0 {
					return
				}
				if level+1 > maxIndent {
					maxIndent = level + 1
				}
				sort.SliceStable(node.Children, func(p, q int) bool {
					return node.Children[p].Order < node.Children[q].Order
				})
				for _, c := range node.Children {
					href := mapPLocatorToHref(linkroleURI, presentation, c.Locator)
					ret = append(ret, IndentedLabel{
						Href:        href,
						Indentation: level,
					})
					makeIndents(c, level+1)
				}
			}
			makeIndents(&root, 0)
			return ret, maxIndent
		}
	}
	return nil, -1
}

func getPresentationContexts(schemedEntity string, instance *xbrl.Instance,
	schema *xbrl.Schema, indentedLabels []IndentedLabel) ([]RelevantContext, int) {
	hrefs := make([]string, len(indentedLabels))
	for i, indentedLabel := range indentedLabels {
		hrefs[i] = indentedLabel.Href
	}
	return getRelevantContexts(schemedEntity, instance, schema, hrefs)
}

func getPFactualQuadrant(indentedLabels []IndentedLabel,
	relevantContexts []RelevantContext,
	factFinder FactFinder) [][]string {
	rowCount := len(indentedLabels)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return [][]string{{}}
	}
	var ret [][]string
	for i := 0; i < rowCount; i++ {
		var row []string
		href := indentedLabels[i].Href
		for j := 0; j < colCount; j++ {
			var fact *xbrl.Fact
			contextRef := relevantContexts[j].ContextRef
			fact = factFinder.FindFact(href, contextRef)
			row = append(row, render(fact))
		}
		ret = append(ret, row)
	}
	return ret
}

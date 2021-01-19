package renderables

import (
	"encoding/json"
	"fmt"
	"sort"

	"ecks-bee.com/telefacts/xbrl"
)

type IndentedLabel struct {
	Href        string
	Label       string
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
			arcs := presentationLink.PresentationArcs
			root := tree(arcs, xbrl.PresentationArcrole)
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
						Label:       href,
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
	hrefs := make([]string, 0, len(indentedLabels))
	for _, indentedLabel := range indentedLabels {
		hrefs = append(hrefs, indentedLabel.Href)
	}
	ret := getFactualQuadrant(hrefs, relevantContexts,
		factFinder)
	return ret
}

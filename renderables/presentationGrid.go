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
	relevantContexts, maxDepth := getPresentationContexts(schemedEntity, instance, schema, indentedLabels, maxIndentation)
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
	schema *xbrl.Schema, indentedLabels []IndentedLabel, maxIndentation int) ([]RelevantContext, int) {
	factuaHrefs := make([]string, 0, len(indentedLabels))
	for _, indentedLabel := range indentedLabels {
		var c *xbrl.Concept
		_, c, _ = xbrl.HashQuery(schema, indentedLabel.Href) //todo catch errors
		if !c.Abstract {
			factuaHrefs = append(factuaHrefs, indentedLabel.Href)
		}
	}
	if len(factuaHrefs) <= 0 {
		return []RelevantContext{}, 0
	}

	maxDepth := 0
	contextRefTaken := make(map[string]bool)
	ret := make([]RelevantContext, 0, len(instance.Context)>>1)
	for _, factualEdgeHref := range factuaHrefs {
		for _, fact := range instance.Facts { //todo parallelize nlogn
			if _, taken := contextRefTaken[fact.ContextRef]; taken {
				continue
			}
			factualHref, _, err := xbrl.NameQuery(schema, fact.XMLName.Space, fact.XMLName.Local)
			if err != nil {
				continue
			}
			if factualEdgeHref == factualHref {
				var context *xbrl.Context
				context = getContext(instance, fact.ContextRef)
				if len(context.Entity.Identitifier) <= 0 {
					continue
				}
				contextualSchemedEntity := context.Entity.Identitifier[0].Scheme + "/" + context.Entity.Identitifier[0].Text
				if contextualSchemedEntity != schemedEntity {
					continue
				}
				dommems := domainMembersString(context)
				if len(dommems) > maxDepth {
					maxDepth = len(dommems)
				}
				ret = append(ret, RelevantContext{
					ContextRef:          context.ID,
					PeriodHeader:        periodString(context),
					DomainMemberHeaders: dommems,
				})
				contextRefTaken[fact.ContextRef] = true
			}
		}
	}
	sortContexts(ret)
	return ret, maxDepth
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

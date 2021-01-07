package renderables

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"ecks-bee.com/telefacts/xbrl"
)

type SummationItem struct {
	Href                 string
	RelevantContexts     []RelevantContext
	MaxDepth             int
	ContributingConcepts []ContributingConcept
	FactualQuadrant      [][]string
}

type ContributingConcept struct {
	Weight          string
	Href            string
	IsSummationItem bool
}

type CGrid struct {
	SummationItems []SummationItem
}

func MarshalCGrid(entityIndex int, relationshipSetIndex int, schema *xbrl.Schema,
	instance *xbrl.Instance, calculation *xbrl.CalculationLinkbase,
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

	summationItems := getSummationItems(schemedEntity, linkroleURI, schema, instance,
		calculation, factFinder)
	return json.Marshal(CGrid{
		SummationItems: summationItems,
	})
}

func getSummationItems(schemedEntity string, linkroleURI string, schema *xbrl.Schema,
	instance *xbrl.Instance, calculation *xbrl.CalculationLinkbase,
	factFinder FactFinder) []SummationItem {
	var calculationLinks []xbrl.CalculationLink
	for _, roleRef := range calculation.RoleRef {
		if linkroleURI == roleRef.RoleURI {
			calculationLinks = calculation.CalculationLinks
			break
		}
	}
	for _, calculationLink := range calculationLinks {
		if calculationLink.Role == linkroleURI {
			arcs := calculationLink.CalculationArc
			type cStruct struct {
				Href   string
				Order  float64
				Weight float64
			}
			cMap := make(map[string][]cStruct)
			for _, arc := range arcs {
				if arc.Arcrole == xbrl.CalculationArcrole {
					order, _ := strconv.ParseFloat(arc.Order, 64)
					weight, _ := strconv.ParseFloat(arc.Weight, 64)
					fromHref := mapCLocatorToHref(linkroleURI, calculation, arc.From)
					cMap[fromHref] = append(cMap[fromHref],
						cStruct{
							Href:   mapCLocatorToHref(linkroleURI, calculation, arc.To),
							Order:  order,
							Weight: weight,
						})
				}
			}
			for key, slice := range cMap {
				sort.SliceStable(cMap[key], func(i, j int) bool {
					return slice[i].Order < slice[j].Order
				})
			}
			ret := make([]SummationItem, 0, len(cMap))
			for from, slice := range cMap {
				contributingConcepts := make([]ContributingConcept, 0, len(slice))
				for _, cstruct := range slice {
					_, isSummationItem := cMap[cstruct.Href]
					contributingConcepts = append(contributingConcepts, ContributingConcept{
						Href:            cstruct.Href,
						Weight:          fmt.Sprintf("%.1f", cstruct.Weight),
						IsSummationItem: isSummationItem,
					})
				}
				maxDepth := 0
				relevantContexts := make([]RelevantContext, 0)
				factualQuadrant := getCFactualQuadrant(from, contributingConcepts, relevantContexts, factFinder)
				ret = append(ret, SummationItem{
					Href:                 from,
					ContributingConcepts: contributingConcepts,
					MaxDepth:             maxDepth,
					RelevantContexts:     relevantContexts,
					FactualQuadrant:      factualQuadrant,
				})
			}
			sort.SliceStable(ret, func(i, j int) bool {
				return ret[i].Href < ret[j].Href
			})
			return ret
		}
	}
	return nil
}

func getCFactualQuadrant(summationItem string, contributingConcepts []ContributingConcept,
	relevantContexts []RelevantContext,
	factFinder FactFinder) [][]string {
	rowCount := len(contributingConcepts)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return [][]string{{}}
	}
	var ret [][]string
	for i := 0; i < rowCount; i++ {
		var row []string
		href := contributingConcepts[i].Href
		for j := 0; j < colCount; j++ {
			var fact *xbrl.Fact
			contextRef := relevantContexts[j].ContextRef
			fact = factFinder.FindFact(href, contextRef)
			row = append(row, render(fact))
		}
		ret = append(ret, row)
	}
	//todo summationItem facts
	return ret
}

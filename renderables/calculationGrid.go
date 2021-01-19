package renderables

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"

	"ecks-bee.com/telefacts/xbrl"
)

type SummationItem struct {
	Href                 string
	Label                string
	RelevantContexts     []RelevantContext
	MaxDepth             int
	ContributingConcepts []ContributingConcept
	FactualQuadrant      [][]string
}

type ContributingConcept struct {
	Sign            string
	Scale           string
	Href            string
	Label           string
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
				Href  string
				Order float64
				Sign  rune
				Scale float64
			}
			cMap := make(map[string][]cStruct)
			for _, arc := range arcs {
				if arc.Arcrole == xbrl.CalculationArcrole {
					order, _ := strconv.ParseFloat(arc.Order, 64)
					weight, _ := strconv.ParseFloat(arc.Weight, 64)
					fromHref := mapCLocatorToHref(linkroleURI, calculation, arc.From)
					sign := '+'
					if weight < 0 {
						sign = '-'
					}
					scale := math.Abs(weight)
					cMap[fromHref] = append(cMap[fromHref],
						cStruct{
							Href:  mapCLocatorToHref(linkroleURI, calculation, arc.To),
							Order: order,
							Sign:  sign,
							Scale: scale,
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
				fqLabels := make([]string, 0, len(slice)+1)
				for _, cstruct := range slice {
					_, isSummationItem := cMap[cstruct.Href]
					sign := fmt.Sprintf("%c", cstruct.Sign)
					scale := fmt.Sprintf("%.1f", cstruct.Scale)
					contributingConcepts = append(contributingConcepts, ContributingConcept{
						Href:            cstruct.Href,
						Label:           cstruct.Href,
						Scale:           scale,
						Sign:            sign,
						IsSummationItem: isSummationItem,
					})
					fqLabels = append(fqLabels, cstruct.Href)
				}
				fqLabels = append(fqLabels, from)
				relevantContexts, maxDepth := getRelevantContexts(schemedEntity, instance, schema, fqLabels)
				factualQuadrant := getFactualQuadrant(fqLabels, relevantContexts, factFinder)
				ret = append(ret, SummationItem{
					Href:                 from,
					Label:                from,
					ContributingConcepts: contributingConcepts,
					MaxDepth:             maxDepth,
					RelevantContexts:     relevantContexts,
					FactualQuadrant:      factualQuadrant,
				})
			}
			sort.SliceStable(ret, func(i, j int) bool {
				return ret[i].Label < ret[j].Label
			})
			return ret
		}
	}
	return nil
}

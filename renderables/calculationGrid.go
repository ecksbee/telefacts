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
	DomainMemberHeaders  []string
	PeriodHeaders        []string
	PeriodLabels         []string
	ContributingConcepts []ContributingConcept
	FactualQuadrant      [][]string
}

type ContributingConcept struct {
	Weight string
	Href   string
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
	// schemedEntity := schemedEntities[entityIndex]

	summationItems := getSummationItems(linkroleURI, schema, calculation)
	// factualQuadrant := getFactualQuadrant(indentedLabels, relevantContexts, factFinder)
	return json.Marshal(CGrid{
		SummationItems: summationItems,
	})
}

func getSummationItems(linkroleURI string, schema *xbrl.Schema, calculation *xbrl.CalculationLinkbase) []SummationItem {
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
					cMap[arc.From] = append(cMap[arc.From],
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
				href := mapCLocatorToHref(linkroleURI, calculation, from)
				contributingConcepts := make([]ContributingConcept, 0, len(slice))
				for _, cstruct := range slice {
					contributingConcepts = append(contributingConcepts, ContributingConcept{
						Href:   cstruct.Href,
						Weight: fmt.Sprintf("%.1f", cstruct.Weight),
					})
				}
				dms := make([]string, 0)
				periodHeaders := make([]string, 0)
				periodLabels := make([]string, 0)
				fq := make([][]string, 0)
				ret = append(ret, SummationItem{
					Href:                 href,
					ContributingConcepts: contributingConcepts,
					DomainMemberHeaders:  dms,
					PeriodHeaders:        periodHeaders,
					PeriodLabels:         periodLabels,
					FactualQuadrant:      fq,
				})
			}
			return ret
		}
	}
	return nil
}

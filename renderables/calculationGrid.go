package renderables

import (
	"fmt"
	"math"
	"sort"

	"ecksbee.com/telefacts/attr"
	"ecksbee.com/telefacts/hydratables"
)

type SummationItem struct {
	Href                 string
	Label                LabelPack
	RelevantContexts     []RelevantContext
	MaxDepth             int
	ContributingConcepts []ContributingConcept
	FactualQuadrant      [][]string
}

type ContributingConcept struct {
	Sign            string
	Scale           string
	Href            string
	Label           LabelPack
	IsSummationItem bool
}

type CGrid struct {
	SummationItems []SummationItem
}

func cGrid(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder) (CGrid, []LabelRole, []Lang, error) {
	summationItems, labelRoles, langs := getSummationItems(schemedEntity, linkroleURI,
		h, factFinder)
	return CGrid{
		SummationItems: summationItems,
	}, labelRoles, langs, nil
}

func getSummationItems(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder) ([]SummationItem, []LabelRole, []Lang) {
	var calculationLinks []hydratables.CalculationLink
	for _, calculation := range h.CalculationLinkbases {
		for _, roleRef := range calculation.RoleRefs {
			if linkroleURI == roleRef.RoleURI {
				calculationLinks = calculation.CalculationLinks
				break
			}
		}
		for _, calculationLink := range calculationLinks {
			if calculationLink.Role == linkroleURI {
				arcs := calculationLink.CalculationArcs
				type cStruct struct {
					Href  string
					Order float64
					Sign  rune
					Scale float64
				}
				cMap := make(map[string][]cStruct)
				for _, arc := range arcs {
					if arc.Arcrole == attr.CalculationArcrole {
						order := arc.Order
						weight := arc.Weight
						fromHref := mapCLocatorToHref(linkroleURI, &calculation, arc.From)
						sign := '+'
						if weight < 0 {
							sign = '-'
						}
						scale := math.Abs(weight)
						cMap[fromHref] = append(cMap[fromHref],
							cStruct{
								Href:  mapCLocatorToHref(linkroleURI, &calculation, arc.To),
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
				labelPacks := make([]LabelPack, 0, len(calculationLinks))
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
							Label:           GetLabel(h, cstruct.Href),
							Scale:           scale,
							Sign:            sign,
							IsSummationItem: isSummationItem,
						})
						fqLabels = append(fqLabels, cstruct.Href)
					}
					fqLabels = append(fqLabels, from)
					relevantContexts, maxDepth, contextualLabelPack := getRelevantContexts(schemedEntity, h, fqLabels)
					factualQuadrant := getFactualQuadrant(fqLabels, relevantContexts, factFinder)
					siLabelPack := GetLabel(h, from)
					labelPacks = append(labelPacks, siLabelPack)
					labelPacks = append(labelPacks, contextualLabelPack...)
					ret = append(ret, SummationItem{
						Href:                 from,
						Label:                siLabelPack,
						ContributingConcepts: contributingConcepts,
						MaxDepth:             maxDepth,
						RelevantContexts:     relevantContexts,
						FactualQuadrant:      factualQuadrant,
					})
				}
				sort.SliceStable(ret, func(i, j int) bool {
					return getPureLabel(ret[i].Label) < getPureLabel(ret[j].Label)
				})
				reduced := reduce(labelPacks)
				labelRoles, langs := destruct(*reduced)
				return ret, labelRoles, langs
			}
		}
	}
	return nil, nil, nil
}

package renderables

import (
	"fmt"
	"math"
	"sort"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
)

type SummationItem struct {
	Href                 string
	Label                LabelPack
	BalanceType          string
	RelevantContexts     []RelevantContext
	MaxDepth             int
	ContributingConcepts []ContributingConcept
	FactualQuadrant      FactualQuadrant
}

type ContributingConcept struct {
	Sign            string
	Scale           string
	BalanceType     string
	Href            string
	Label           LabelPack
	IsSummationItem bool
}

type CGrid struct {
	SummationItems []SummationItem
}

func cGrid(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder) (CGrid, []LabelRole, []Lang, error) {
	summationItems, labelRoles, langs := getSummationItems(schemedEntity, linkroleURI,
		h, factFinder, conceptFinder, measurementFinder)
	return CGrid{
		SummationItems: summationItems,
	}, labelRoles, langs, nil
}

func getSummationItems(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder) ([]SummationItem, []LabelRole, []Lang) {
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
				var (
					labelRoles []LabelRole
					langs      []Lang
				)
				ret := make([]SummationItem, 0, len(cMap))
				for from, slice := range cMap {
					contributingConcepts := make([]ContributingConcept, 0, len(slice))
					fqLabels := make([]string, 0, len(slice)+1)
					for _, cstruct := range slice {
						_, isSummationItem := cMap[cstruct.Href]
						sign := fmt.Sprintf("%c", cstruct.Sign)
						scale := fmt.Sprintf("%.1f", cstruct.Scale)
						cLabelPack := GetLabel(h, cstruct.Href)
						_, concept, err := h.HashQuery(cstruct.Href)
						if err != nil {
							continue
						}
						contributingConcepts = append(contributingConcepts, ContributingConcept{
							Href:            cstruct.Href,
							Label:           cLabelPack,
							BalanceType:     concept.Balance,
							Scale:           scale,
							Sign:            sign,
							IsSummationItem: isSummationItem,
						})
						fqLabels = append(fqLabels, cstruct.Href)
						labelPacks = append(labelPacks, cLabelPack)
					}
					fqLabels = append(fqLabels, from)
					relevantContexts, maxDepth, contextualLabelPack := getRelevantContexts(schemedEntity, h, fqLabels)
					siLabelPack := GetLabel(h, from)
					labelPacks = append(labelPacks, siLabelPack)
					labelPacks = append(labelPacks, contextualLabelPack...)
					reduced := reduce(labelPacks)
					if reduced != nil {
						labelRoles, langs = destruct(*reduced)
					}
					factualQuadrant := getFactualQuadrant(fqLabels, relevantContexts, factFinder, conceptFinder, measurementFinder, labelRoles, langs)
					_, sumConcept, err := h.HashQuery(from)
					if err != nil {
						continue
					}
					ret = append(ret, SummationItem{
						Href:                 from,
						Label:                siLabelPack,
						BalanceType:          sumConcept.Balance,
						ContributingConcepts: contributingConcepts,
						MaxDepth:             maxDepth,
						RelevantContexts:     relevantContexts,
						FactualQuadrant:      factualQuadrant,
					})
				}
				sort.SliceStable(ret, func(i, j int) bool {
					return getPureLabel(ret[i].Label) < getPureLabel(ret[j].Label)
				})
				return ret, labelRoles, langs
			}
		}
	}
	return nil, nil, nil
}

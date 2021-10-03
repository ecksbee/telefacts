package renderables

import (
	"sort"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
)

type IndentedLabel struct {
	Href        string
	Label       LabelPack
	Indentation int
}

type PGrid struct {
	IndentedLabels []IndentedLabel
	PeriodHeaders
	ContextualMemberGrid
	VoidQuadrant
	FactualQuadrant FactualQuadrant
	FootnoteGrid    [][][]int
	Footnotes       []string
}

func pGrid(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder, conceptFinder ConceptFinder,
	measurementFinder MeasurementFinder) (PGrid, []LabelRole, []Lang, error) {
	indentedLabels, labelPacks := getIndentedLabels(linkroleURI, h)
	relevantContexts, segmentTypedDomainTrees, scenarioTypedDomainTrees, contextualLabelPacks :=
		getPresentationContexts(schemedEntity, h, indentedLabels)
	labelPacks = append(labelPacks, contextualLabelPacks...)
	reduced := reduce(labelPacks)
	var (
		labelRoles []LabelRole
		langs      []Lang
	)
	if reduced != nil {
		labelRoles, langs = destruct(*reduced)
	}
	factualQuadrant, footnoteGrid, footnotes := getPFactualQuadrant(indentedLabels,
		relevantContexts, factFinder, conceptFinder, measurementFinder, langs)
	memberGrid, voidQuadrant := getMemberGridAndVoidQuadrant(relevantContexts,
		segmentTypedDomainTrees, scenarioTypedDomainTrees)
	return PGrid{
		IndentedLabels:       indentedLabels,
		PeriodHeaders:        getPeriodHeaders(relevantContexts),
		ContextualMemberGrid: memberGrid,
		VoidQuadrant:         voidQuadrant,
		FactualQuadrant:      factualQuadrant,
		FootnoteGrid:         footnoteGrid,
		Footnotes:            footnotes,
	}, labelRoles, langs, nil
}

func pArcs(pArcs []hydratables.PresentationArc) []arc {
	ret := make([]arc, 0, len(pArcs))
	for _, pArc := range pArcs {
		ret = append(ret, arc{
			Arcrole: pArc.Arcrole,
			Order:   pArc.Order,
			From:    pArc.From,
			To:      pArc.To,
		})
	}
	return ret
}

func getIndentedLabels(linkroleURI string, h *hydratables.Hydratable) ([]IndentedLabel, []LabelPack) {
	labelPacks := make([]LabelPack, 0, 100)
	for _, presentation := range h.PresentationLinkbases {
		var presentationLinks []hydratables.PresentationLink
		for _, roleRef := range presentation.RoleRefs {
			if linkroleURI == roleRef.RoleURI {
				presentationLinks = presentation.PresentationLinks
				break
			}
		}
		for _, presentationLink := range presentationLinks {
			if presentationLink.Role == linkroleURI {
				arcs := presentationLink.PresentationArcs
				pArcs := pArcs(arcs)
				root := tree(pArcs, attr.PresentationArcrole)
				ret := make([]IndentedLabel, 0, len(arcs))
				var makeIndents func(node *locatorNode, level int)
				makeIndents = func(node *locatorNode, level int) {
					if len(node.Children) <= 0 {
						return
					}
					sort.SliceStable(node.Children, func(p, q int) bool {
						return node.Children[p].Order < node.Children[q].Order
					})
					for _, c := range node.Children {
						href := mapPLocatorToHref(linkroleURI, &presentation, c.Locator)
						iLabel := GetLabel(h, href)
						ret = append(ret, IndentedLabel{
							Href:        href,
							Label:       iLabel,
							Indentation: level,
						})
						labelPacks = append(labelPacks, iLabel)
						makeIndents(c, level+1)
					}
				}
				makeIndents(&root, 0)
				return ret, labelPacks
			}
		}
	}
	return nil, labelPacks
}

func getPresentationContexts(schemedEntity string, h *hydratables.Hydratable,
	indentedLabels []IndentedLabel) ([]relevantContext, []locatorNode,
	[]locatorNode, []LabelPack) {
	hrefs := make([]string, len(indentedLabels))
	for i, indentedLabel := range indentedLabels {
		hrefs[i] = indentedLabel.Href
	}
	return getRelevantContexts(schemedEntity, h, hrefs)
}

func getPFactualQuadrant(indentedLabels []IndentedLabel,
	relevantContexts []relevantContext, factFinder FactFinder,
	conceptFinder ConceptFinder, measurementFinder MeasurementFinder,
	langs []Lang) (FactualQuadrant, [][][]int, []string) {
	hrefs := make([]string, 0, len(indentedLabels))
	for _, indentedLabel := range indentedLabels {
		hrefs = append(hrefs, indentedLabel.Href)
	}
	ret, footnoteGrid, footnotes := getFactualQuadrant(hrefs,
		relevantContexts, factFinder, conceptFinder,
		measurementFinder, langs)
	return ret, footnoteGrid, footnotes
}

package renderables

import (
	"sort"

	"ecksbee.com/telefacts/attr"
	"ecksbee.com/telefacts/hydratables"
)

type IndentedLabel struct {
	Href        string
	Label       LabelPack
	Indentation int
}

type PGrid struct {
	IndentedLabels   []IndentedLabel
	MaxIndentation   int
	RelevantContexts []RelevantContext
	MaxDepth         int
	FactualQuadrant  [][]string
}

func pGrid(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder) (PGrid, []LabelRole, []Lang, error) {
	indentedLabels, maxIndentation := getIndentedLabels(linkroleURI, h)
	relevantContexts, maxDepth, labelPacks := getPresentationContexts(schemedEntity, h, indentedLabels)
	reduced := reduce(labelPacks)
	labelRoles, langs := destruct(*reduced)
	factualQuadrant := getPFactualQuadrant(indentedLabels, relevantContexts, factFinder)
	return PGrid{
		IndentedLabels:   indentedLabels,
		MaxIndentation:   maxIndentation,
		RelevantContexts: relevantContexts,
		MaxDepth:         maxDepth,
		FactualQuadrant:  factualQuadrant,
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

func getIndentedLabels(linkroleURI string, h *hydratables.Hydratable) ([]IndentedLabel, int) {
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
				maxIndent := 1
				var makeIndents func(node *locatorNode, level int)
				makeIndents = func(node *locatorNode, level int) {
					if len(node.Children) <= 0 {
						return
					}
					if level+1 > maxIndent {
						maxIndent = level + 1
					}
					for _, c := range node.Children {
						href := mapPLocatorToHref(linkroleURI, &presentation, c.Locator)
						ret = append(ret, IndentedLabel{
							Href:        href,
							Label:       GetLabel(h, href),
							Indentation: level,
						})
						makeIndents(c, level+1)
					}
					sort.SliceStable(node.Children, func(p, q int) bool {
						return node.Children[p].Order < node.Children[q].Order
					})
				}
				makeIndents(&root, 0)
				return ret, maxIndent
			}
		}
	}
	return nil, -1
}

func getPresentationContexts(schemedEntity string, h *hydratables.Hydratable,
	indentedLabels []IndentedLabel) ([]RelevantContext, int, []LabelPack) {
	hrefs := make([]string, len(indentedLabels))
	for i, indentedLabel := range indentedLabels {
		hrefs[i] = indentedLabel.Href
	}
	return getRelevantContexts(schemedEntity, h, hrefs)
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

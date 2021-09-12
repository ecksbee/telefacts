package renderables

import (
	"sort"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
)

type RootDomain struct {
	Href         string
	Label        LabelPack
	PrimaryItems []PrimaryItem
	PeriodHeaders
	ContextualMemberGrid
	VoidQuadrant
	FactualQuadrant     FactualQuadrant
	EffectiveDomainGrid [][]EffectiveDomain
	EffectiveDimensions []EffectiveDimension
	DRSNodes            []DRSNode `json:",omitempty"`
}

type PrimaryItem struct {
	Href  string
	Label LabelPack
	Level int
}

type EffectiveDimension struct {
	Href  string
	Label LabelPack
}

type EffectiveDomain []EffectiveMember

type EffectiveMember struct {
	Href            string
	Label           LabelPack
	IsDefault       bool
	IsStrikethrough bool
}

func dArcs(dArcs []hydratables.DefinitionArc) []arc {
	ret := make([]arc, 0, len(dArcs))
	for _, dArc := range dArcs {
		ret = append(ret, arc{
			Arcrole: dArc.Arcrole,
			Order:   dArc.Order,
			From:    dArc.From,
			To:      dArc.To,
		})
	}
	return ret
}

func getRootDomains(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder) ([]RootDomain, []LabelRole, []Lang) {
	ret := []RootDomain{}
	labelRoles := []LabelRole{}
	langs := []Lang{}
	labelPacks := make([]LabelPack, 0, 100)
	for _, definition := range h.DefinitionLinkbases {
		var definitionLinks []hydratables.DefinitionLink
		for _, roleRef := range definition.RoleRefs {
			if linkroleURI == roleRef.RoleURI {
				definitionLinks = definition.DefinitionLinks
				break
			}
		}
		for _, definitionLink := range definitionLinks {
			if definitionLink.Role == linkroleURI {
				arcs := definitionLink.DefinitionArcs
				indentedItems := make([]PrimaryItem, 0, len(arcs))
				var makeIndents func(node *locatorNode, level int)
				makeIndents = func(node *locatorNode, level int) {
					if len(node.Children) <= 0 {
						return
					}
					sort.SliceStable(node.Children, func(p, q int) bool {
						return node.Children[p].Order < node.Children[q].Order
					})
					for _, c := range node.Children {
						href := mapDLocatorToHref(linkroleURI, &definition, c.Locator)
						piLabel := GetLabel(h, href)
						labelPacks = append(labelPacks, piLabel)
						indentedItems = append(indentedItems, PrimaryItem{
							Href:  href,
							Label: piLabel,
							Level: level,
						})
						makeIndents(c, level+1)
					}
				}
				dArcs := dArcs(arcs)
				domainMemberNetwork := tree(dArcs, attr.DomainMemberArcrole)
				effectiveDimensions, effectiveDimensionHrefs, edLabelRoles, edLangs := getEffectiveDimensions(linkroleURI, arcs, h)
				labelRoles = append(labelRoles, edLabelRoles...)
				langs = append(langs, edLangs...)
				dimensionDomainNetwork := tree(dArcs, attr.DimensionDomainArcrole)
				defaultDimensionsNetwork := tree(dArcs, attr.DimensionDefaultArcrole)
				primaryItemNetwork, explicitDomainNetwork :=
					getPrimaryItemNetworkAndExplicitDomainNetwork(&domainMemberNetwork, &dimensionDomainNetwork,
						&defaultDimensionsNetwork)
				exclusiveHypercubeNetwork := tree(dArcs, attr.HasExclusiveHypercubeArcrole)
				inclusiveHypercubeNetwork := tree(dArcs, attr.HasInclusiveHypercubeArcrole)
				hypercubeDimensionNetwork := tree(dArcs, attr.HypercubeDimensionArcrole)
				for _, root := range primaryItemNetwork.Children {
					makeIndents(root, 0)
					rootHref := mapDLocatorToHref(linkroleURI, &definition, root.Locator)
					primaryItemHrefs := []string{}
					primaryItemHrefs = append(primaryItemHrefs, rootHref)
					for _, indentedItem := range indentedItems {
						primaryItemHrefs = append(primaryItemHrefs, indentedItem.Href)
					}
					locToHref := func(loc string) string {
						return mapDLocatorToHref(linkroleURI, &definition, loc)
					}
					edGrid, edLabels := getEffectiveDomainGrid(primaryItemHrefs, effectiveDimensionHrefs,
						&dimensionDomainNetwork, primaryItemNetwork,
						explicitDomainNetwork, &exclusiveHypercubeNetwork, &inclusiveHypercubeNetwork,
						&hypercubeDimensionNetwork, &defaultDimensionsNetwork, locToHref, h)
					labelPacks = append(labelPacks, edLabels...)
					relevantContexts, segmentTypedDomainArcs, scenarioTypedDomainTrees, contextualLabelPack :=
						getRelevantContexts(schemedEntity, h, primaryItemHrefs)
					labelPacks = append(labelPacks, contextualLabelPack...)
					rdLabelPack := GetLabel(h, rootHref)
					labelPacks = append(labelPacks, rdLabelPack)
					memberGrid, voidQuadrant := getMemberGridAndVoidQuadrant(relevantContexts, segmentTypedDomainArcs, scenarioTypedDomainTrees)
					rootDomain := RootDomain{
						PrimaryItems:         indentedItems,
						Href:                 rootHref,
						Label:                rdLabelPack,
						PeriodHeaders:        getPeriodHeaders(relevantContexts),
						ContextualMemberGrid: memberGrid,
						VoidQuadrant:         voidQuadrant,
						EffectiveDimensions:  effectiveDimensions,
						EffectiveDomainGrid:  edGrid,
					}
					reduced := reduce(labelPacks)
					if reduced != nil {
						dLabelRoles, dLangs := destruct(*reduced)
						labelRoles = append(labelRoles, dLabelRoles...)
						langs = append(langs, dLangs...)
					}
					rootDomain = injectFactualQuadrant(rootDomain, relevantContexts, factFinder, conceptFinder, measurementFinder, langs)
					ret = append(ret, rootDomain)
				}
			}
		}
	}
	return ret, labelRoles, langs
}

func getEffectiveDimensions(linkroleURI string,
	arcs []hydratables.DefinitionArc, h *hydratables.Hydratable) ([]EffectiveDimension, []string, []LabelRole, []Lang) {
	labelPacks := []LabelPack{}
	effectiveDimensionMap := make(map[string]bool)
	effectiveDimensions := make([]EffectiveDimension, 0, len(arcs))
	effectiveDimensionHrefs := make([]string, 0, len(arcs))
	for _, definition := range h.DefinitionLinkbases {
		for _, arc := range arcs {
			if arc.Arcrole == attr.HypercubeDimensionArcrole {
				dim := mapDLocatorToHref(linkroleURI, &definition, arc.To)
				if effectiveDimensionMap[dim] {
					continue
				}
				effectiveDimensionMap[dim] = true
				effectiveDimensionHrefs = append(effectiveDimensionHrefs, dim)
				dimLabelPack := GetLabel(h, dim)
				labelPacks = append(labelPacks, dimLabelPack)
				effectiveDimensions = append(effectiveDimensions, EffectiveDimension{
					Href:  dim,
					Label: dimLabelPack,
				})
			}
		}
	}
	reduced := reduce(labelPacks)
	labelRoles := []LabelRole{}
	langs := []Lang{}
	if reduced != nil {
		labelRoles, langs = destruct(*reduced)
	}
	return effectiveDimensions, effectiveDimensionHrefs, labelRoles, langs
}

func getPrimaryItemNetworkAndExplicitDomainNetwork(domainMemberNetwork *locatorNode,
	dimensionDomainNetwork *locatorNode, defaultDimensionsNetwork *locatorNode) (*locatorNode, *locatorNode) {
	primaryItemNetwork := locatorNode{}
	explicitDomainNetwork := locatorNode{}
	hypercubeDomains := make([]string, 0, len(dimensionDomainNetwork.Children)+len(defaultDimensionsNetwork.Children))
	for _, dimNode := range dimensionDomainNetwork.Children {
		for _, domNode := range dimNode.Children {
			hypercubeDomains = append(hypercubeDomains, domNode.Locator)
		}
	}
	for _, dimNode := range defaultDimensionsNetwork.Children {
		for _, defNode := range dimNode.Children {
			hypercubeDomains = append(hypercubeDomains, defNode.Locator)
		}
	}
	for _, dmNode := range domainMemberNetwork.Children {
		if !stringInSlice(dmNode.Locator, hypercubeDomains) {
			primaryItemNetwork.Children = append(primaryItemNetwork.Children, dmNode)
		} else {
			explicitDomainNetwork.Children = append(explicitDomainNetwork.Children, dmNode)
		}
	}
	return &primaryItemNetwork, &explicitDomainNetwork
}

func injectFactualQuadrant(incompleteRootDomain RootDomain, relevantContexts []relevantContext,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder,
	langs []Lang) RootDomain {
	hrefs := make([]string, 0, len(incompleteRootDomain.PrimaryItems)+1)
	hrefs = append(hrefs, incompleteRootDomain.Href)
	for _, primaryItem := range incompleteRootDomain.PrimaryItems {
		hrefs = append(hrefs, primaryItem.Href)
	}
	factualQuadrant := getFactualQuadrant(hrefs, relevantContexts, factFinder, conceptFinder,
		measurementFinder, langs)
	incompleteRootDomain.FactualQuadrant = factualQuadrant
	return incompleteRootDomain
}

func getEffectiveDomainGrid(primaryItemHrefs []string, effectiveDimensionHrefs []string,
	dimensionDomainNetwork *locatorNode, primaryItemNetwork *locatorNode,
	explicitDomainNetwork *locatorNode, exclusiveHypercubeNetwork *locatorNode, inclusiveHypercubeNetwork *locatorNode,
	hypercubeDimensionNetwork *locatorNode, dimensionDefaultNetwork *locatorNode,
	mapDLocatorToHref func(string) string, h *hydratables.Hydratable) ([][]EffectiveDomain, []LabelPack) {
	ret := make([][]EffectiveDomain, 0, len(primaryItemHrefs))
	labelPacks := make([]LabelPack, 0, len(primaryItemHrefs))
	for _, primaryItem := range primaryItemHrefs {
		row := make([]EffectiveDomain, 0, len(effectiveDimensionHrefs))
		for _, effectiveDimension := range effectiveDimensionHrefs {
			effectiveDomain, edLabels := getEffectiveDomain(primaryItem, effectiveDimension,
				dimensionDomainNetwork, primaryItemNetwork, explicitDomainNetwork, exclusiveHypercubeNetwork,
				inclusiveHypercubeNetwork, hypercubeDimensionNetwork, dimensionDefaultNetwork, mapDLocatorToHref, h)
			row = append(row, effectiveDomain)
			labelPacks = append(labelPacks, edLabels...)
		}
		ret = append(ret, row)
	}
	return ret, labelPacks
}

func getEffectiveDomain(primaryItemHref string, effectiveDimensionHref string,
	dimensionDomainNetwork *locatorNode, primaryItemNetwork *locatorNode,
	explicitDomainNetwork *locatorNode, exclusiveHypercubeNetwork *locatorNode, inclusiveHypercubeNetwork *locatorNode,
	hypercubeDimensionNetwork *locatorNode, dimensionDefaultNetwork *locatorNode,
	mapDLocatorToHref func(string) string, h *hydratables.Hydratable) (EffectiveDomain, []LabelPack) {
	inclusiveHypercubeHrefs := []string{}
	exclusiveHypercubeHrefs := []string{}
	for _, inclusiveHypercubeNode := range inclusiveHypercubeNetwork.Children {
		if primaryItemHref == mapDLocatorToHref(inclusiveHypercubeNode.Locator) {
			for _, inclusiveHypercubeChildNode := range inclusiveHypercubeNode.Children {
				inclusiveHypercubeHref := mapDLocatorToHref(inclusiveHypercubeChildNode.Locator)
				for _, hypercubeDimensionNode := range hypercubeDimensionNetwork.Children {
					if inclusiveHypercubeHref == mapDLocatorToHref(hypercubeDimensionNode.Locator) {
						for _, hypercubeDimensionChildNode := range hypercubeDimensionNode.Children {
							if effectiveDimensionHref == mapDLocatorToHref(hypercubeDimensionChildNode.Locator) {
								inclusiveHypercubeHrefs = append(inclusiveHypercubeHrefs, inclusiveHypercubeHref)
							}
						}
					}
				}
			}
		}
	}
	for _, exclusiveHypercubeNode := range exclusiveHypercubeNetwork.Children {
		if primaryItemHref == mapDLocatorToHref(exclusiveHypercubeNode.Locator) {
			for _, exclusiveHypercubeChildNode := range exclusiveHypercubeNode.Children {
				exclusiveHypercubeHref := mapDLocatorToHref(exclusiveHypercubeChildNode.Locator)
				for _, hypercubeDimensionNode := range hypercubeDimensionNetwork.Children {
					if exclusiveHypercubeHref == mapDLocatorToHref(hypercubeDimensionNode.Locator) {
						for _, hypercubeDimensionChildNode := range hypercubeDimensionNode.Children {
							if effectiveDimensionHref == mapDLocatorToHref(hypercubeDimensionChildNode.Locator) {
								exclusiveHypercubeHrefs = append(exclusiveHypercubeHrefs, exclusiveHypercubeHref)
							}
						}
					}
				}
			}
		}
	}
	for _, inclusiveHypercubeHref := range inclusiveHypercubeHrefs {
		for _, exclusiveHypercubeHref := range exclusiveHypercubeHrefs {
			if inclusiveHypercubeHref == exclusiveHypercubeHref {
				return []EffectiveMember{}, []LabelPack{}
			}
		}
	}
	labelPacks := make([]LabelPack, 0, 100)
	ret := []EffectiveMember{}
	defaultMembersMap := make(map[string]bool)
	for _, hypercubeDimensionNode := range hypercubeDimensionNetwork.Children {
		for _, inclusiveHypercubeHref := range inclusiveHypercubeHrefs {
			if inclusiveHypercubeHref == mapDLocatorToHref(hypercubeDimensionNode.Locator) {
				for _, dimensionDefaultNode := range dimensionDefaultNetwork.Children {
					defaultDimensionHref := mapDLocatorToHref(dimensionDefaultNode.Locator)
					if defaultDimensionHref != effectiveDimensionHref {
						continue
					}
					for _, defaultMemberNode := range dimensionDefaultNode.Children {
						defaultMemberHref := mapDLocatorToHref(defaultMemberNode.Locator)
						if defaultMembersMap[defaultMemberHref] {
							return []EffectiveMember{}, []LabelPack{}
						}
						defaultMembersMap[defaultMemberHref] = true
						defMemLabel := GetLabel(h, defaultMemberHref)
						labelPacks = append(labelPacks, defMemLabel)
						ret = append(ret, EffectiveMember{
							Href:            defaultMemberHref,
							Label:           defMemLabel,
							IsDefault:       true,
							IsStrikethrough: false,
						})
					}
				}
			}
		}
	}
	excludedDefaultMembersMap := make(map[string]bool)
	for _, hypercubeDimensionNode := range hypercubeDimensionNetwork.Children {
		for _, exclusiveHypercubeHref := range exclusiveHypercubeHrefs {
			if exclusiveHypercubeHref == mapDLocatorToHref(hypercubeDimensionNode.Locator) {
				for _, dimensionDefaultNode := range dimensionDefaultNetwork.Children {
					defaultDimensionHref := mapDLocatorToHref(dimensionDefaultNode.Locator)
					if excludedDefaultMembersMap[defaultDimensionHref] {
						continue
					}
					for _, defaultMemberNode := range dimensionDefaultNode.Children {
						defaultMemberHref := mapDLocatorToHref(defaultMemberNode.Locator)
						if excludedDefaultMembersMap[defaultMemberHref] {
							return []EffectiveMember{}, []LabelPack{}
						}
						excludedDefaultMembersMap[defaultMemberHref] = true
						defMemLabel := GetLabel(h, defaultMemberHref)
						labelPacks = append(labelPacks, defMemLabel)
						ret = append(ret, EffectiveMember{
							Href:            defaultMemberHref,
							Label:           defMemLabel,
							IsDefault:       true,
							IsStrikethrough: true,
						})
					}
				}
			}
		}
	}

	membersMap := make(map[string]bool)
	excludedMembersMap := make(map[string]bool)
	var traverseExplicitDomainNetwork func(*locatorNode, bool)
	traverseExplicitDomainNetwork = func(node *locatorNode, isInclusive bool) {
		memberHref := mapDLocatorToHref(node.Locator)
		if isInclusive {
			if membersMap[memberHref] {
				ret = []EffectiveMember{}
				labelPacks = []LabelPack{}
				return
			}
			membersMap[memberHref] = true
			efMemLabel := GetLabel(h, memberHref)
			ret = append(ret, EffectiveMember{
				Href:            memberHref,
				Label:           efMemLabel,
				IsDefault:       false,
				IsStrikethrough: false,
			})
		} else {
			if excludedMembersMap[memberHref] {
				ret = []EffectiveMember{}
				labelPacks = []LabelPack{}
				return
			}
			excludedMembersMap[memberHref] = true
			efMemLabel := GetLabel(h, memberHref)
			ret = append(ret, EffectiveMember{
				Href:            memberHref,
				Label:           efMemLabel,
				IsDefault:       false,
				IsStrikethrough: true,
			})
		}
		if len(node.Children) <= 0 {
			return
		}
		for _, child := range node.Children {
			traverseExplicitDomainNetwork(child, isInclusive)
		}
	}
	for _, hypercubeDimensionNode := range hypercubeDimensionNetwork.Children {
		for _, inclusiveHypercubeHref := range inclusiveHypercubeHrefs {
			if inclusiveHypercubeHref == mapDLocatorToHref(hypercubeDimensionNode.Locator) {
				for _, inclusiveDimensionNode := range hypercubeDimensionNode.Children {
					inclusiveDimensionHref := mapDLocatorToHref(inclusiveDimensionNode.Locator)
					if effectiveDimensionHref != inclusiveDimensionHref {
						continue
					}
					for _, dimensionDomainNode := range dimensionDomainNetwork.Children {
						if inclusiveDimensionHref != mapDLocatorToHref(dimensionDomainNode.Locator) {
							continue
						}
						for _, node := range explicitDomainNetwork.Children {
							traverseExplicitDomainNetwork(node, true)
						}
					}
				}
			}
		}
	}
	for _, hypercubeDimensionNode := range hypercubeDimensionNetwork.Children {
		for _, exclusiveHypercubeHref := range exclusiveHypercubeHrefs {
			if exclusiveHypercubeHref == mapDLocatorToHref(hypercubeDimensionNode.Locator) {
				for _, exclusiveDimensionNode := range hypercubeDimensionNode.Children {
					exclusiveDimensionHref := mapDLocatorToHref(exclusiveDimensionNode.Locator)
					if effectiveDimensionHref != exclusiveDimensionHref {
						continue
					}
					for _, dimensionDomainNode := range dimensionDomainNetwork.Children {
						if exclusiveDimensionHref != mapDLocatorToHref(dimensionDomainNode.Locator) {
							continue
						}
						for _, node := range explicitDomainNetwork.Children {
							traverseExplicitDomainNetwork(node, false)
						}
					}
				}
			}
		}
	}
	sort.SliceStable(ret, func(i int, j int) bool {
		a := ret[i]
		b := ret[j]
		if a.Href != b.Href {
			return a.Href < b.Href
		}
		if a.IsDefault != b.IsDefault {
			return a.IsDefault
		}
		if a.IsStrikethrough != b.IsStrikethrough {
			return b.IsStrikethrough
		}
		return false
	})
	return ret, labelPacks
}

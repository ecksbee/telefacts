package renderables

import (
	"sort"

	"ecksbee.com/telefacts/attr"
	"ecksbee.com/telefacts/hydratables"
)

type RootDomain struct {
	Href                string
	Label               LabelPack
	RelevantContexts    []RelevantContext
	MaxDepth            int
	MaxLevel            int
	PrimaryItems        []PrimaryItem
	FactualQuadrant     [][]string
	EffectiveDomainGrid [][]EffectiveDomain
	EffectiveDimensions []EffectiveDimension
	Hypercubes          []Hypercube
}

type PrimaryItem struct {
	Href       string
	Label      LabelPack
	Level      int
	Hypercubes []Hypercube
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

type Hypercube struct {
	Href           string
	Label          LabelPack
	IsClosed       bool
	ContextElement string
	IsInclusive    bool
	Nodes          []struct {
		Source struct {
			Href  string
			Label string
		}
		Target struct {
			Href  string
			Label string
		}
		Order   string
		Default bool
		Usable  bool
	}
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
	factFinder FactFinder) ([]RootDomain, []LabelRole, []Lang) {
	ret := []RootDomain{}
	labelRoles := []LabelRole{}
	langs := []Lang{}
	for _, definition := range h.DefinitionLinkbases {
		var definitionLinks []hydratables.DefinitionLink
		for _, roleRef := range definition.RoleRefs {
			if linkroleURI == roleRef.RoleURI {
				definitionLinks = definition.DefinitionLinks
				break
			}
		}
		ret := []RootDomain{}
		for _, definitionLink := range definitionLinks {
			if definitionLink.Role == linkroleURI {
				arcs := definitionLink.DefinitionArcs
				indentedItems := make([]PrimaryItem, 0, len(arcs))
				maxIndent := 0
				var makeIndents func(node *locatorNode, level int)
				makeIndents = func(node *locatorNode, level int) {
					if len(node.Children) <= 0 {
						return
					}
					if level+1 > maxIndent {
						maxIndent = level + 1
					}
					for _, c := range node.Children {
						href := mapDLocatorToHref(linkroleURI, &definition, c.Locator)
						indentedItems = append(indentedItems, PrimaryItem{
							Href:       href,
							Label:      GetLabel(h, href),
							Level:      level,
							Hypercubes: getHypercubes(c.Locator, arcs, linkroleURI, &definition, h),
						})
						makeIndents(c, level+1)
					}
					sort.SliceStable(node.Children, func(p, q int) bool {
						return node.Children[p].Order < node.Children[q].Order
					})
				}
				dArcs := dArcs(arcs)
				domainMemberNetwork := tree(dArcs, attr.DomainMemberArcrole)
				effectiveDimensions, effectiveDimensionHrefs := getEffectiveDimensions(linkroleURI, arcs, h)
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
					relevantContexts, maxDepth, labelPacks := getRelevantContexts(schemedEntity, h, primaryItemHrefs)
					rdLabelPack := GetLabel(h, rootHref)
					labelPacks = append(labelPacks, rdLabelPack)
					rootDomain := RootDomain{
						PrimaryItems:        indentedItems,
						Href:                rootHref,
						Label:               GetLabel(h, rootHref),
						MaxLevel:            maxIndent,
						RelevantContexts:    relevantContexts,
						MaxDepth:            maxDepth,
						EffectiveDimensions: effectiveDimensions,
						EffectiveDomainGrid: getEffectiveDomainGrid(primaryItemHrefs, effectiveDimensionHrefs,
							&dimensionDomainNetwork, primaryItemNetwork,
							explicitDomainNetwork, &exclusiveHypercubeNetwork, &inclusiveHypercubeNetwork,
							&hypercubeDimensionNetwork, &defaultDimensionsNetwork, locToHref, h),
						Hypercubes: getHypercubes(root.Locator, arcs, linkroleURI, &definition, h),
					}
					rootDomain = injectFactualQuadrant(rootDomain, relevantContexts, factFinder)
					ret = append(ret, rootDomain)
					reduced := reduce(labelPacks)
					labelRoles, langs = destruct(*reduced)
				}
			}
		}
	}
	return ret, labelRoles, langs
}

func getEffectiveDimensions(linkroleURI string,
	arcs []hydratables.DefinitionArc, h *hydratables.Hydratable) ([]EffectiveDimension, []string) {
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
				effectiveDimensions = append(effectiveDimensions, EffectiveDimension{
					Href:  dim,
					Label: GetLabel(h, dim),
				})
			}
		}
	}
	return effectiveDimensions, effectiveDimensionHrefs
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

func injectFactualQuadrant(incompleteRootDomain RootDomain, relevantContexts []RelevantContext,
	factFinder FactFinder) RootDomain {
	hrefs := make([]string, 0, len(incompleteRootDomain.PrimaryItems)+1)
	hrefs = append(hrefs, incompleteRootDomain.Href)
	for _, primaryItem := range incompleteRootDomain.PrimaryItems {
		hrefs = append(hrefs, primaryItem.Href)
	}
	factualQuadrant := getFactualQuadrant(hrefs, relevantContexts, factFinder)
	incompleteRootDomain.FactualQuadrant = factualQuadrant
	return incompleteRootDomain
}

func getHypercubes(primaryItemLoc string, arcs []hydratables.DefinitionArc, linkroleURI string,
	definition *hydratables.DefinitionLinkbase, h *hydratables.Hydratable) []Hypercube {
	ret := make([]Hypercube, 0, len(arcs))
	for _, arc := range arcs {
		if arc.Arcrole == attr.HasExclusiveHypercubeArcrole || arc.Arcrole == attr.HasInclusiveHypercubeArcrole {
			if arc.From != primaryItemLoc {
				continue
			}
			href := mapDLocatorToHref(linkroleURI, definition, arc.To)
			ret = append(ret, Hypercube{
				Href:           href,
				Label:          GetLabel(h, href),
				IsClosed:       arc.Closed,
				ContextElement: arc.ContextElement,
				IsInclusive:    arc.Arcrole == attr.HasInclusiveHypercubeArcrole,
			})
		}
	}
	return ret
}

func getEffectiveDomainGrid(primaryItemHrefs []string, effectiveDimensionHrefs []string,
	dimensionDomainNetwork *locatorNode, primaryItemNetwork *locatorNode,
	explicitDomainNetwork *locatorNode, exclusiveHypercubeNetwork *locatorNode, inclusiveHypercubeNetwork *locatorNode,
	hypercubeDimensionNetwork *locatorNode, dimensionDefaultNetwork *locatorNode,
	mapDLocatorToHref func(string) string, h *hydratables.Hydratable) [][]EffectiveDomain {
	ret := make([][]EffectiveDomain, 0, len(primaryItemHrefs))
	for _, primaryItem := range primaryItemHrefs {
		row := make([]EffectiveDomain, 0, len(effectiveDimensionHrefs))
		for _, effectiveDimension := range effectiveDimensionHrefs {
			effectiveDomain := getEffectiveDomain(primaryItem, effectiveDimension,
				dimensionDomainNetwork, primaryItemNetwork, explicitDomainNetwork, exclusiveHypercubeNetwork,
				inclusiveHypercubeNetwork, hypercubeDimensionNetwork, dimensionDefaultNetwork, mapDLocatorToHref, h)
			row = append(row, effectiveDomain)
		}
		ret = append(ret, row)
	}
	return ret
}

func getEffectiveDomain(primaryItemHref string, effectiveDimensionHref string,
	dimensionDomainNetwork *locatorNode, primaryItemNetwork *locatorNode,
	explicitDomainNetwork *locatorNode, exclusiveHypercubeNetwork *locatorNode, inclusiveHypercubeNetwork *locatorNode,
	hypercubeDimensionNetwork *locatorNode, dimensionDefaultNetwork *locatorNode,
	mapDLocatorToHref func(string) string, h *hydratables.Hydratable) EffectiveDomain {
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
				return []EffectiveMember{}
			}
		}
	}
	ret := []EffectiveMember{}
	defaultMembersMap := make(map[string]bool)
	excludedDefaultMembersMap := make(map[string]bool)
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
							return []EffectiveMember{}
						}
						defaultMembersMap[defaultMemberHref] = true
						ret = append(ret, EffectiveMember{
							Href:            defaultMemberHref,
							Label:           GetLabel(h, defaultMemberHref),
							IsDefault:       true,
							IsStrikethrough: false,
						})
					}
				}
			}
		}
	}
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
							return []EffectiveMember{}
						}
						excludedDefaultMembersMap[defaultMemberHref] = true
						ret = append(ret, EffectiveMember{
							Href:            defaultMemberHref,
							Label:           GetLabel(h, defaultMemberHref),
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
				return
			}
			membersMap[memberHref] = true
			ret = append(ret, EffectiveMember{
				Href:            memberHref,
				Label:           GetLabel(h, memberHref),
				IsDefault:       false,
				IsStrikethrough: false,
			})
		} else {
			if excludedMembersMap[memberHref] {
				ret = []EffectiveMember{}
				return
			}
			excludedMembersMap[memberHref] = true
			ret = append(ret, EffectiveMember{
				Href:            memberHref,
				Label:           GetLabel(h, memberHref),
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
		if a.IsStrikethrough != a.IsStrikethrough {
			return b.IsStrikethrough
		}
		return false
	})
	return ret
}

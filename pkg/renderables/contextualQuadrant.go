package renderables

import (
	"sort"

	myarcs "github.com/joshuanario/arcs"
)

type PeriodHeaders []LanguagePack

type ContextualMemberCell struct {
	ExplicitMember *ExplicitMember `json:",omitempty"`
	TypedMember    string          `json:",omitempty"`
}

type VoidCell struct {
	IsParenthesized bool
	Indentation     int
	Dimension       Dimension
	TypedDomain     *TypedDomain `json:",omitempty"`
}

type ContextualMemberGrid [][]*ContextualMemberCell

type VoidQuadrant []*VoidCell

func getPeriodHeaders(relevantContexts []relevantContext) PeriodHeaders {
	ret := make(PeriodHeaders, len(relevantContexts))
	for i, ctx := range relevantContexts {
		ret[i] = ctx.PeriodHeader
	}
	return ret
}

func getMemberGridAndVoidQuadrant(relevantContexts []relevantContext,
	segmentTypedDomainTrees []myarcs.RArc, scenarioTypedDomainTrees []myarcs.RArc) (ContextualMemberGrid,
	VoidQuadrant) {
	if len(relevantContexts) <= 0 {
		return ContextualMemberGrid{}, VoidQuadrant{}
	}
	sort.SliceStable(segmentTypedDomainTrees, func(i, j int) bool {
		if segmentTypedDomainTrees[i].Locator == segmentTypedDomainTrees[j].Locator {
			return i < j
		}
		return segmentTypedDomainTrees[i].Locator < segmentTypedDomainTrees[j].Locator
	})
	sort.SliceStable(scenarioTypedDomainTrees, func(i, j int) bool {
		if scenarioTypedDomainTrees[i].Locator == scenarioTypedDomainTrees[j].Locator {
			return i < j
		}
		return scenarioTypedDomainTrees[i].Locator < scenarioTypedDomainTrees[j].Locator
	})
	voidQuadrant := getVoidQuadrant(relevantContexts, segmentTypedDomainTrees, scenarioTypedDomainTrees)
	rowCount := len(voidQuadrant)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return ContextualMemberGrid{}, VoidQuadrant{}
	}
	ret := make(ContextualMemberGrid, rowCount)
	for i, voidCell := range voidQuadrant {
		if voidCell == nil {
			return ContextualMemberGrid{}, VoidQuadrant{}
		}
		row := make([]*ContextualMemberCell, colCount)
		for j, ctx := range relevantContexts {
			cell := &ContextualMemberCell{}
			for _, ctxMember := range ctx.Members {
				if ctxMember.Dimension.Href == voidCell.Dimension.Href {
					if ctxMember.TypedDomain != nil && voidCell.TypedDomain != nil {
						if ctxMember.TypedDomain.Href == voidCell.TypedDomain.Href {
							cell = &ContextualMemberCell{
								TypedMember: ctxMember.TypedMember,
							}
							break
						}
					} else {
						if ctxMember.TypedDomain == nil && voidCell.TypedDomain == nil {
							cell = &ContextualMemberCell{
								ExplicitMember: ctxMember.ExplicitMember,
							}
							break
						}
					}
				}
			}
			row[j] = cell
		}
		ret[i] = row
	}
	return ret, voidQuadrant
}

func getVoidQuadrant(relevantContexts []relevantContext, segmentTypedDomainTrees []myarcs.RArc,
	scenarioTypedDomainTrees []myarcs.RArc) VoidQuadrant {
	segmentExplicitDimensionMap := make(map[string]*RelevantMember)
	scenarioExplicitDimensionMap := make(map[string]*RelevantMember)
	segmentTypedDimensionMap := make(map[string]*myarcs.RArc)
	scenarioTypedDimensionMap := make(map[string]*myarcs.RArc)
	allDimensionMap := make(map[string]*Dimension)
	allTypedDomainMap := make(map[string]*TypedDomain)
	for i := 0; i < len(relevantContexts); i++ {
		members := relevantContexts[i].Members
		for _, member := range members {
			if member.Dimension.Href != "" {
				allDimensionMap[member.Dimension.Href] = &member.Dimension
			}
			if member.TypedDomain != nil {
				allTypedDomainMap[member.TypedDomain.Href] = member.TypedDomain
			}
			if member.ExplicitMember != nil {
				if member.IsSegment {
					segmentExplicitDimensionMap[member.Dimension.Href] = &member
				} else {
					scenarioExplicitDimensionMap[member.Dimension.Href] = &member
				}
			}
		}
	}

	ret := make(VoidQuadrant, 0, len(segmentTypedDimensionMap)+
		len(scenarioTypedDimensionMap)+len(segmentExplicitDimensionMap)+
		len(scenarioExplicitDimensionMap))
	for _, member := range segmentExplicitDimensionMap {
		ret = append(ret, &VoidCell{
			IsParenthesized: false,
			Indentation:     0,
			Dimension:       member.Dimension,
		})
	}
	for _, member := range scenarioExplicitDimensionMap {
		ret = append(ret, &VoidCell{
			IsParenthesized: true,
			Indentation:     0,
			Dimension:       member.Dimension,
		})
	}

	reducedSegmentTypedDomainTree := reduceTrees(segmentTypedDomainTrees)
	reducedScenarioTypedDomainTree := reduceTrees(scenarioTypedDomainTrees)
	var dimension *Dimension
	var makeIndents func(node *myarcs.RArc, level int, isParenthesized bool)
	makeIndents = func(node *myarcs.RArc, level int, isParenthesized bool) {
		if dimension == nil && node.Locator != "" {
			mappedDimension, ok := allDimensionMap[node.Locator]
			if !ok {
				ret = VoidQuadrant{}
				return
			}
			dimension = mappedDimension
		}
		if level == 1 {
			ret = append(ret, &VoidCell{
				IsParenthesized: isParenthesized,
				Indentation:     level,
				Dimension:       *dimension,
				TypedDomain:     nil,
			})
		}
		if level > 1 {
			if node.Locator != "" {
				typedDomain, ok := allTypedDomainMap[node.Locator]
				if !ok {
					ret = VoidQuadrant{}
					return
				}
				ret = append(ret, &VoidCell{
					IsParenthesized: isParenthesized,
					Indentation:     level,
					Dimension:       *dimension,
					TypedDomain:     typedDomain,
				})
			}
		}
		if len(node.Children) <= 0 {
			return
		}
		sort.SliceStable(node.Children, func(p, q int) bool {
			return node.Children[p].Order < node.Children[q].Order
		})
		for _, c := range node.Children {
			makeIndents(c, level+1, isParenthesized)
		}
	}
	makeIndents(&reducedSegmentTypedDomainTree, 0, false)
	makeIndents(&reducedScenarioTypedDomainTree, 0, true)
	return ret
}

func reduceTrees(trees []myarcs.RArc) myarcs.RArc {
	ret := myarcs.RArc{}
	hasNonblankRoots := true
	dimensions := make([]*myarcs.RArc, 0, len(trees))
	for _, root := range trees {
		if root.Locator == "" && len(root.Children) > 0 {
			dimensions = append(dimensions, root.Children...)
		} else {
			hasNonblankRoots = false
		}
	}
	if !hasNonblankRoots {
		return myarcs.RArc{}
	}
	ret.Order = 0
	ret.Locator = ""
	ret.Children = make([]*myarcs.RArc, 0)
	ret = *dedupNodes(dimensions, ret, int(ret.Order))
	return ret
}

func dedupNodes(children []*myarcs.RArc, dst myarcs.RArc, order int) *myarcs.RArc {
	if children == nil || len(children) <= 0 {
		return &dst
	}
	ret := dst
	ret.Order = float64(order)
	dupLocators := make([]string, 0, len(children))
	for _, node := range children {
		dupLocators = append(dupLocators, node.Locator)
	}
	dedupLocators := dedup(dupLocators)
	dedupLocators = sort.StringSlice(dedupLocators)
	dedupChildren := make([]*myarcs.RArc, 0)
	for _, dedupLocator := range dedupLocators {
		order++
		dupGrandchildren := make([]*myarcs.RArc, 0)
		for _, dupChild := range children {
			if dupChild.Locator == dedupLocator {
				copied := make([]*myarcs.RArc, len(dupChild.Children))
				copy(copied, dupChild.Children)
				dupGrandchildren = append(dupGrandchildren, copied...)
			}
		}
		dstChild := myarcs.RArc{
			Locator: dedupLocator,
			Order:   float64(order),
		}
		dstChild = *dedupNodes(dupGrandchildren, dstChild, order)
		dedupChildren = append(dedupChildren, &dstChild)
	}
	ret.Children = dedupChildren
	return &ret
}

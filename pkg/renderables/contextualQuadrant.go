package renderables

import (
	"sort"
)

type PeriodHeaders []*LanguagePack

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
		ret[i] = &ctx.PeriodHeader
	}
	return ret
}

func getMemberGridAndVoidQuadrant(relevantContexts []relevantContext,
	segmentTypedDomainTrees []locatorNode, scenarioTypedDomainTrees []locatorNode) (ContextualMemberGrid,
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
	for i := 0; i < rowCount; i++ {
		row := make([]*ContextualMemberCell, colCount)
		voidCell := voidQuadrant[i]
		if voidCell == nil {
			return ContextualMemberGrid{}, VoidQuadrant{}
		}
		ctx := relevantContexts[i]
		for j := 0; j < colCount; j++ {
			for _, ctxMember := range ctx.Members {
				if ctxMember.Dimension.Href == voidCell.Dimension.Href {
					if ctxMember.TypedDomain.Href == voidCell.TypedDomain.Href {
						row[j] = &ContextualMemberCell{
							TypedMember: ctxMember.TypedMember,
						}
					} else {
						row[j] = &ContextualMemberCell{
							ExplicitMember: ctxMember.ExplicitMember,
						}
					}
					continue
				}
			}
		}
		ret = append(ret, row)
	}
	return ret, VoidQuadrant{}
}

func getVoidQuadrant(relevantContexts []relevantContext, segmentTypedDomainTrees []locatorNode,
	scenarioTypedDomainTrees []locatorNode) VoidQuadrant {
	segmentExplicitDimensionMap := make(map[string]*ContextualMember)
	scenarioExplicitDimensionMap := make(map[string]*ContextualMember)
	segmentTypedDimensionMap := make(map[string]*locatorNode)
	scenarioTypedDimensionMap := make(map[string]*locatorNode)
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

	for i := 0; i < len(segmentTypedDomainTrees); i++ {
		root := segmentTypedDomainTrees[i]
		if root.Locator != "" {
			segmentTypedDimensionMap[root.Locator] = &root
		}
	}

	for i := 0; i < len(segmentTypedDomainTrees); i++ {
		root := segmentTypedDomainTrees[i]
		if root.Locator != "" {
			segmentTypedDimensionMap[root.Locator] = &root
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
	for _, root := range segmentTypedDimensionMap {
		var makeIndents func(node *locatorNode, level int)
		makeIndents = func(node *locatorNode, level int) {
			if len(node.Children) <= 0 {
				ret = VoidQuadrant{}
				return
			}
			dimension, ok := allDimensionMap[node.Locator]
			if !ok {
				ret = VoidQuadrant{}
				return
			}
			sort.SliceStable(node.Children, func(p, q int) bool {
				return node.Children[p].Order < node.Children[q].Order
			})
			for _, c := range node.Children {
				if level <= 0 {
					ret = append(ret, &VoidCell{
						IsParenthesized: false,
						Indentation:     level,
						Dimension:       *dimension,
					})
				} else {
					typedDomain, ok := allTypedDomainMap[node.Locator]
					if !ok {
						ret = VoidQuadrant{}
						return
					}
					ret = append(ret, &VoidCell{
						IsParenthesized: false,
						Indentation:     level,
						Dimension:       *dimension,
						TypedDomain:     typedDomain,
					})
				}
				makeIndents(c, level+1)
			}
		}
		makeIndents(root, 0)
	}
	return ret
}

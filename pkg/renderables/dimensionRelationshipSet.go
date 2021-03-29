package renderables

import (
	"sort"
	"strconv"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
)

type HypercubeConnection struct {
	IsClosed       bool
	ContextElement string
	IsInclusive    bool
	External       string
}
type HypercubeDimensionConnection struct {
	Order    string
	External string
}
type DimensionDomainConnection struct {
	Order    string
	Default  bool
	Usable   bool
	External string
}
type DomainMemberConnection struct {
	Order    string
	Usable   bool
	External string
}

type DRSNode struct {
	Href  string
	Label LabelPack
}

type DRSLink struct {
	SourceHref                   string
	TargetHref                   string
	HypercubeConnection          *HypercubeConnection          `json:",omitempty"`
	HypercubeDimensionConnection *HypercubeDimensionConnection `json:",omitempty"`
	DimensionDomainConnection    *DimensionDomainConnection    `json:",omitempty"`
	DomainMemberConnection       *DomainMemberConnection       `json:",omitempty"`
}

type DRS struct {
	Nodes []DRSNode
	Links []DRSLink
}

func getDRS(schemedEntity string, linkroleURI string, h *hydratables.Hydratable) DRS {
	links := make([]DRSLink, 0, 20)
	hrefs := make(map[string]int)

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
				for _, arc := range arcs {
					var to string
					if arc.TargetRole != "" {
						to = mapDLocatorToHref(arc.TargetRole, &definition, arc.To)
					} else {
						to = mapDLocatorToHref(linkroleURI, &definition, arc.To)
					}
					hrefs[to]++
					from := mapDLocatorToHref(linkroleURI, &definition, arc.From)
					hrefs[from]++
					var (
						hc *HypercubeConnection
						hd *HypercubeDimensionConnection
						dd *DimensionDomainConnection
						dm *DomainMemberConnection
					)
					order := strconv.FormatFloat(arc.Order, 'f', -1, 64)
					switch arc.Arcrole {
					case attr.DomainMemberArcrole:
						dm = &DomainMemberConnection{
							Order:    order,
							Usable:   arc.Usable,
							External: arc.TargetRole,
						}
					case attr.HypercubeDimensionArcrole:
						hd = &HypercubeDimensionConnection{
							Order:    order,
							External: arc.TargetRole,
						}
					case attr.DimensionDefaultArcrole:
						dd = &DimensionDomainConnection{
							Order:    order,
							Default:  true,
							Usable:   true,
							External: arc.TargetRole,
						}
					case attr.DimensionDomainArcrole:
						dd = &DimensionDomainConnection{
							Order:    order,
							Default:  true,
							Usable:   arc.Usable,
							External: arc.TargetRole,
						}
					case attr.HasInclusiveHypercubeArcrole:
						hc = &HypercubeConnection{
							IsClosed:       arc.Closed,
							ContextElement: arc.ContextElement,
							IsInclusive:    true,
							External:       arc.TargetRole,
						}
					case attr.HasExclusiveHypercubeArcrole:
						hc = &HypercubeConnection{
							IsClosed:       arc.Closed,
							ContextElement: arc.ContextElement,
							IsInclusive:    false,
							External:       arc.TargetRole,
						}
					}
					links = append(links, DRSLink{
						SourceHref:                   from,
						TargetHref:                   to,
						DomainMemberConnection:       dm,
						HypercubeConnection:          hc,
						HypercubeDimensionConnection: hd,
						DimensionDomainConnection:    dd,
					})
				}
			}
		}
	}
	nodes := make([]DRSNode, 0, 20)
	for href := range hrefs {
		nodes = append(nodes, DRSNode{
			Href:  href,
			Label: GetLabel(h, href),
		})
	}
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Href < nodes[j].Href })

	return DRS{
		Nodes: nodes,
		Links: links,
	}
}

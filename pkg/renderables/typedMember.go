package renderables

import (
	"ecksbee.com/telefacts/pkg/hydratables"
)

const typedDomainArcRole = "http://ecksbee.com/arc-role/typed-domain"

func tDomArcs(pArcs []hydratables.TypedDomainArc) []arc {
	ret := make([]arc, 0, len(pArcs))
	for _, pArc := range pArcs {
		ret = append(ret, arc{
			Arcrole: typedDomainArcRole,
			Order:   pArc.Order,
			From:    pArc.From,
			To:      pArc.To,
		})
	}
	return ret
}

func getTypedMember(typedMember hydratables.TypedMember, dimension Dimension, isSegment bool, h *hydratables.Hydratable) ([]RelevantMember, []arc, []LabelPack) {
	ret := make([]RelevantMember, 0, len(typedMember.TypedMembersMap))
	arcs := tDomArcs(typedMember.TypedDomainArcs)
	labelPacks := []LabelPack{
		dimension.Label,
	}
	for typedDomain, typedMember := range typedMember.TypedMembersMap {
		typedDomainLabel := GetLabel(h, typedDomain)
		ret = append(ret, RelevantMember{
			Dimension: dimension,
			TypedDomain: &TypedDomain{
				Href:  typedDomain,
				Label: typedDomainLabel,
			},
			TypedMember: typedMember,
		})
	}
	return ret, arcs, labelPacks
}

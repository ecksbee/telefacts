package renderables

import (
	"bufio"
	"encoding/xml"
	"io"
	"strings"

	"ecksbee.com/telefacts/pkg/hydratables"
)

const typedDomainArcRole = "http://ecksbee.com/arc-role/typed-domain"

func getTypedMember(typedMember hydratables.TypedMember, dimension Dimension, isSegment bool, h *hydratables.Hydratable) ([]ContextualMember, []arc, []LabelPack) {
	reader := bufio.NewReader(strings.NewReader(typedMember.Value))
	d := xml.NewDecoder(reader)
	ret := make([]ContextualMember, 0)
	arcs := make([]arc, 0)
	labelPacks := []LabelPack{
		dimension.Label,
	}
	order := 0
	var stack *Stack
	prevHref := dimension.Href
	_, prev, err := h.HashQuery(dimension.Href)
	if err != nil {
		return []ContextualMember{}, []arc{}, []LabelPack{}
	}
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		switch token.(type) {
		case xml.StartElement:
			startElem, ok := token.(xml.StartElement)
			if ok {
				currHref, curr, err := h.NameQuery(startElem.Name.Space, startElem.Name.Local)
				if err != nil {
					return []ContextualMember{}, []arc{}, []LabelPack{}
				}
				order++
				stack.Push(curr)
				arcs = append(arcs, arc{
					Arcrole: typedDomainArcRole,
					From:    prevHref,
					To:      currHref,
				})
				typedDomainLabel := GetLabel(h, currHref)
				labelPacks = append(labelPacks, typedDomainLabel)
				ret = append(ret, ContextualMember{
					Dimension: dimension,
					IsSegment: isSegment,
					TypedDomain: &TypedDomain{
						Href:  currHref,
						Label: typedDomainLabel,
					},
				})
				prev = curr
				prevHref = currHref
			}
		case xml.EndElement:
			endElem, ok := token.(xml.EndElement)
			if ok {
				popped, ok := stack.Pop()
				if !ok {
					return []ContextualMember{}, []arc{}, []LabelPack{}
				}
				if (popped.XMLName.Local != endElem.Name.Local) && (popped.XMLName.Space != endElem.Name.Space) {
					return []ContextualMember{}, []arc{}, []LabelPack{}
				}
				prev, ok := stack.Pop()
				if !ok {
					return []ContextualMember{}, []arc{}, []LabelPack{}
				}
				stack.Push(prev)
			}
		case xml.CharData:
			charData, ok := token.(xml.CharData)
			if ok {
				str := strings.TrimSpace(string(charData))
				typedDomainHref, _, err := h.NameQuery(prev.XMLName.Space, prev.XMLName.Local)
				if err != nil {
					return []ContextualMember{}, []arc{}, []LabelPack{}
				}
				typedDomainLabel := GetLabel(h, typedDomainHref)
				labelPacks = append(labelPacks, typedDomainLabel)
				ret = append(ret, ContextualMember{
					Dimension:   dimension,
					IsSegment:   isSegment,
					TypedMember: str,
					TypedDomain: &TypedDomain{
						Href:  typedDomainHref,
						Label: typedDomainLabel,
					},
				})
			}
		}
	}
	return ret, arcs, labelPacks
}

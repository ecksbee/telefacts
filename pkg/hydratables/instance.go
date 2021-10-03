package hydratables

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
)

type Instance struct {
	FileName      string
	Contexts      []Context
	Units         []Unit
	FootnoteLinks []FootnoteLink
	Facts         []Fact
}

type ExplicitMember struct {
	Dimension struct {
		Href  string
		Value string
	}
	Member struct {
		Href     string
		CharData string
	}
}

type TypedMember struct {
	Dimension struct {
		Href            string
		Value           string
		TypedDomainHref string
	}
	TypedMembersMap map[string]string
	TypedDomainArcs []TypedDomainArc
}

type Context struct {
	ID     string
	Entity Entity
	Period struct {
		Instant  Instant
		Duration Duration
	}
	Scenario DimensionContext
}

type Entity struct {
	Identifier struct {
		Scheme   string
		CharData string
	}
	Segment DimensionContext
}

type DimensionContext struct {
	ExplicitMembers []ExplicitMember
	TypedMembers    []TypedMember
}

type TypedDomainArc struct {
	Order float64
	From  string
	To    string
}

type Instant struct {
	CharData string
}

type Duration struct {
	StartDate string
	EndDate   string
}

type FootnoteLink struct {
	Title        string
	Footnotes    []Footnote
	Locs         []FootnoteLinkLoc
	FootnoteArcs []FootnoteArc
}

type Footnote struct {
	ID       string
	CharData string
	Lang     string
	Label    string
}

type FootnoteLinkLoc struct {
	Href  string
	Label string
}

type FootnoteArc struct {
	From    string
	To      string
	Arcrole string
}

func HydrateInstance(file *serializables.InstanceFile, fileName string, h *Hydratable) (*Instance, error) {
	if len(fileName) <= 0 {
		return nil, fmt.Errorf("empty file name")
	}
	if file == nil {
		return nil, fmt.Errorf("empty file")
	}
	ret := Instance{}
	ret.FileName = fileName
	ret.Contexts = hydrateContexts(file, h)
	ret.Units = hydrateUnits(file)
	ret.Facts = hydrateFacts(file, h)
	ret.FootnoteLinks = hydrateFootnoteLinks(file)
	return &ret, nil
}

func hydrateContexts(instanceFile *serializables.InstanceFile, h *Hydratable) []Context {
	ret := make([]Context, 0, len(instanceFile.Context))
	for _, context := range instanceFile.Context {
		if context.XMLName.Space != attr.XBRLI {
			continue
		}
		item := Context{}
		idAttr := attr.FindAttr(context.XMLAttrs, "id")
		if idAttr == nil || idAttr.Value == "" {
			continue
		}
		item.ID = idAttr.Value
		newEntity := Entity{}
		entity := context.Entity[0]
		if len(entity.Identifier) <= 0 {
			continue
		}
		if entity.Identifier[0].CharData == "" {
			continue
		}
		schemeAttr := attr.FindAttr(entity.Identifier[0].XMLAttrs, "scheme")
		if schemeAttr == nil {
			continue
		}
		newEntity.Identifier.CharData = entity.Identifier[0].CharData
		newEntity.Identifier.Scheme = schemeAttr.Value
		if len(entity.Segment) > 0 {
			segment := DimensionContext{}
			if len(entity.Segment[0].ExplicitMember) > 0 {
				segment.ExplicitMembers = make([]ExplicitMember, 0, len(entity.Segment[0].ExplicitMember))
				for _, explicitMember := range entity.Segment[0].ExplicitMember {
					dimAttr := attr.FindAttr(explicitMember.XMLAttrs, "dimension")
					if dimAttr == nil {
						continue
					}
					dimName := attr.Xmlns(instanceFile.XMLAttrs, dimAttr.Value)
					if dimName.Space == "" {
						continue
					}
					dimRef, dimConcept, err := h.NameQuery(dimName.Space, dimName.Local)
					if err != nil || dimRef == "" || dimConcept == nil {
						continue
					}
					memName := attr.Xmlns(instanceFile.XMLAttrs, explicitMember.CharData)
					if memName.Space == "" {
						continue
					}
					memRef, memConcept, err := h.NameQuery(memName.Space, memName.Local)
					if err != nil || memRef == "" || memConcept == nil {
						continue
					}
					newExplicitMember := ExplicitMember{
						Dimension: struct {
							Href  string
							Value string
						}{
							Href:  dimRef,
							Value: dimAttr.Value,
						},
						Member: struct {
							Href     string
							CharData string
						}{
							Href:     memRef,
							CharData: explicitMember.CharData,
						},
					}
					segment.ExplicitMembers = append(segment.ExplicitMembers, newExplicitMember)
				}
			}
			if len(entity.Segment[0].TypedMember) > 0 {
				segment.TypedMembers = make([]TypedMember, 0, len(entity.Segment[0].TypedMember))
				for _, typedMember := range entity.Segment[0].TypedMember {
					dimAttr := attr.FindAttr(typedMember.XMLAttrs, "dimension")
					if dimAttr == nil {
						continue
					}
					dimName := attr.Xmlns(instanceFile.XMLAttrs, dimAttr.Value)
					if dimName.Space == "" {
						continue
					}
					dimRef, dimConcept, err := h.NameQuery(dimName.Space, dimName.Local)
					if err != nil || dimRef == "" || dimConcept == nil {
						continue
					}
					typedMembersMap, typedDomainArcs := hydrateTypedMembers(dimRef, typedMember.XMLInner,
						instanceFile, h)
					newTypedMember := TypedMember{
						Dimension: struct {
							Href            string
							Value           string
							TypedDomainHref string
						}{
							Href:            dimRef,
							TypedDomainHref: dimConcept.TypedDomainHref,
							Value:           dimAttr.Value,
						},
						TypedMembersMap: typedMembersMap,
						TypedDomainArcs: typedDomainArcs,
					}
					segment.TypedMembers = append(segment.TypedMembers, newTypedMember)
				}
			}
			newEntity.Segment = segment
		}
		item.Entity = newEntity
		if len(context.Period) > 0 {
			if len(context.Period[0].Instant) > 0 {
				instant := Instant{
					CharData: context.Period[0].Instant[0].CharData,
				}
				item.Period.Instant = instant
			}
			if len(context.Period[0].StartDate) > 0 && len(context.Period[0].EndDate) > 0 {
				duration := Duration{
					StartDate: context.Period[0].StartDate[0].CharData,
					EndDate:   context.Period[0].EndDate[0].CharData,
				}
				item.Period.Duration = duration
			}
		}
		if len(context.Scenario) > 0 {
			scenario := DimensionContext{}
			if len(context.Scenario[0].ExplicitMember) > 0 {
				scenario.ExplicitMembers = make([]ExplicitMember, 0, len(context.Scenario[0].ExplicitMember))
				for _, explicitMember := range context.Scenario[0].ExplicitMember {
					dimAttr := attr.FindAttr(explicitMember.XMLAttrs, "dimension")
					if dimAttr == nil || dimAttr.Value == "" {
						continue
					}
					dimName := attr.Xmlns(instanceFile.XMLAttrs, dimAttr.Value)
					if dimName.Space == "" {
						continue
					}
					dimRef, dimConcept, err := h.NameQuery(dimName.Space, dimName.Local)
					if err != nil || dimRef == "" || dimConcept == nil {
						continue
					}
					memName := attr.Xmlns(instanceFile.XMLAttrs, explicitMember.CharData)
					if memName.Space == "" {
						continue
					}
					memRef, memConcept, err := h.NameQuery(memName.Space, memName.Local)
					if err != nil || memRef == "" || memConcept == nil {
						continue
					}
					newExplicitMember := ExplicitMember{
						Dimension: struct {
							Href  string
							Value string
						}{
							Href:  dimRef,
							Value: dimAttr.Value,
						},
						Member: struct {
							Href     string
							CharData string
						}{
							Href:     memRef,
							CharData: explicitMember.CharData,
						},
					}
					scenario.ExplicitMembers = append(scenario.ExplicitMembers, newExplicitMember)
				}
			}
			if len(context.Scenario[0].TypedMember) > 0 {
				scenario.TypedMembers = make([]TypedMember, 0, len(context.Scenario[0].TypedMember))
				for _, typedMember := range context.Scenario[0].TypedMember {
					dimAttr := attr.FindAttr(typedMember.XMLAttrs, "dimension")
					if dimAttr == nil || dimAttr.Value == "" {
						continue
					}
					dimName := attr.Xmlns(instanceFile.XMLAttrs, dimAttr.Value)
					if dimName.Space == "" {
						continue
					}
					dimRef, dimConcept, err := h.NameQuery(dimName.Space, dimName.Local)
					if err != nil || dimRef == "" || dimConcept == nil {
						continue
					}
					typedMembersMap, typedDomainArcs := hydrateTypedMembers(dimRef, typedMember.XMLInner,
						instanceFile, h)
					newTypedMember := TypedMember{
						Dimension: struct {
							Href            string
							Value           string
							TypedDomainHref string
						}{
							Href:            dimRef,
							Value:           dimAttr.Value,
							TypedDomainHref: dimConcept.TypedDomainHref,
						},
						TypedMembersMap: typedMembersMap,
						TypedDomainArcs: typedDomainArcs,
					}
					scenario.TypedMembers = append(scenario.TypedMembers, newTypedMember)
				}
			}
			item.Scenario = scenario
		}
		ret = append(ret, item)
	}
	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].ID < ret[j].ID
	})
	return ret
}

func hydrateTypedMembers(dimensionHref string, typedMember string, instanceFile *serializables.InstanceFile, h *Hydratable) (map[string]string, []TypedDomainArc) {
	reader := bufio.NewReader(strings.NewReader(typedMember))
	d := xml.NewDecoder(reader)
	retMap := make(map[string]string)
	arcs := make([]TypedDomainArc, 0)
	order := 0
	var stack Stack
	stack = make([]*Concept, 0)
	prevHref := dimensionHref
	_, prev, err := h.HashQuery(dimensionHref)
	dimension := prev
	if err != nil {
		return nil, nil
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
				startElemName := attr.Xmlns(instanceFile.XMLAttrs, startElem.Name.Space+":"+startElem.Name.Local)
				if startElemName.Space == "" {
					continue
				}
				currHref, curr, err := h.NameQuery(startElemName.Space, startElemName.Local)
				if err != nil {
					return nil, nil
				}
				order++
				stack.Push(curr)
				arcs = append(arcs, TypedDomainArc{
					From:  prevHref,
					To:    currHref,
					Order: float64(order),
				})
				prev = curr
				prevHref = currHref
			}
		case xml.EndElement:
			endElem, ok := token.(xml.EndElement)
			if ok {
				popped, ok := stack.Pop()
				if !ok {
					return nil, nil
				}
				endElemName := attr.Xmlns(instanceFile.XMLAttrs, endElem.Name.Space+":"+endElem.Name.Local)
				if (popped.XMLName.Local != endElemName.Local) && (popped.XMLName.Space != endElemName.Space) {
					return nil, nil
				}
				prev, ok = stack.Pop()
				if ok {
					stack.Push(prev)
				} else {
					prev = dimension
				}
			}
		case xml.CharData:
			charData, ok := token.(xml.CharData)
			if ok {
				str := strings.TrimSpace(string(charData))
				if str == "" {
					continue
				}
				typedDomainHref, _, err := h.NameQuery(prev.XMLName.Space, prev.XMLName.Local)
				if err != nil {
					return nil, nil
				}
				retMap[typedDomainHref] = str
			}
		}
	}
	return retMap, arcs
}

func hydrateUnits(instanceFile *serializables.InstanceFile) []Unit {
	ret := make([]Unit, 0, len(instanceFile.Unit))
	for _, unit := range instanceFile.Unit {
		if unit.XMLName.Space != attr.XBRLI {
			continue
		}
		idAttr := attr.FindAttr(unit.XMLAttrs, "id")
		if idAttr == nil || idAttr.Value == "" {
			continue
		}
		item := Unit{}
		if len(unit.Measure) <= 0 {
			if len(unit.Divide) <= 0 {
				continue
			}
			if len(unit.Divide[0].UnitDenominator) <= 0 || len(unit.Divide[0].UnitNumerator) <= 0 {
				continue
			}
			if len(unit.Divide[0].UnitDenominator[0].Measure) <= 0 || len(unit.Divide[0].UnitNumerator[0].Measure) <= 0 {
				continue
			}
			numeratorPrefixedName := unit.Divide[0].UnitNumerator[0].Measure[0].CharData
			numeratorMeasure := UnitMeasure{
				XMLName:  attr.Xmlns(instanceFile.XMLAttrs, numeratorPrefixedName),
				CharData: numeratorPrefixedName,
			}
			denominatorPrefixedName := unit.Divide[0].UnitDenominator[0].Measure[0].CharData
			denominatorMeasure := UnitMeasure{
				XMLName:  attr.Xmlns(instanceFile.XMLAttrs, denominatorPrefixedName),
				CharData: denominatorPrefixedName,
			}
			divide := UnitDivide{
				UnitNumerator: struct{ Measure UnitMeasure }{
					Measure: numeratorMeasure,
				},
				UnitDenominator: struct{ Measure UnitMeasure }{
					Measure: denominatorMeasure,
				},
			}
			item.Divide = divide
		} else {
			itemPrefixedName := unit.Measure[0].CharData
			item.Measure = UnitMeasure{
				XMLName:  attr.Xmlns(instanceFile.XMLAttrs, itemPrefixedName),
				CharData: itemPrefixedName,
			}
		}
		item.ID = idAttr.Value
		ret = append(ret, item)
	}

	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].ID < ret[j].ID
	})
	return ret
}

func hydrateFacts(instanceFile *serializables.InstanceFile, h *Hydratable) []Fact {
	ret := make([]Fact, 0, len(instanceFile.Facts))
	for _, fact := range instanceFile.Facts {
		idAttr := attr.FindAttr(fact.XMLAttrs, "id")
		if idAttr == nil || idAttr.Value == "" {
			continue
		}
		if fact.XMLName.Local == "" || fact.XMLName.Space == "" {
			continue
		}
		factRef, factConcept, err := h.NameQuery(fact.XMLName.Space, fact.XMLName.Local)
		if err != nil || factRef == "" || factConcept == nil {
			continue
		}
		contextRefAttr := attr.FindAttr(fact.XMLAttrs, "contextRef")
		if contextRefAttr == nil || contextRefAttr.Value == "" {
			continue
		}
		unitRefAttr := attr.FindAttr(fact.XMLAttrs, "unitRef")
		unitVal := ""
		if unitRefAttr != nil {
			unitVal = unitRefAttr.Value
		}
		// precisionAttr := attr.FindAttr(fact.XMLAttrs, "precision")	//todo
		// if precisionAttr != nil {
		// 	if precisionAttr.Value == "INF" {
		// 		precisionVal = Exact
		// 	} else {
		// 		precision, err := strconv.Atoi(precisionAttr.Value)
		// 		if err == nil {
		// 			precisionVal = Precision(precision)
		// 		}
		// 	}
		// }
		precisionVal := Precisionless
		decimalsAttr := attr.FindAttr(fact.XMLAttrs, "decimals")
		if decimalsAttr != nil {
			if decimalsAttr.Value == "INF" {
				precisionVal = Exact
			} else {
				decimal, err := strconv.Atoi(decimalsAttr.Value)
				if err == nil {
					precisionVal = Precision(decimal)
				}
			}
		}
		nilAttr := attr.FindAttr(fact.XMLAttrs, "nil")
		nilVal := false
		if nilAttr != nil {
			nilVal, _ = strconv.ParseBool(nilAttr.Value)
		}
		newFact := Fact{
			ID:         idAttr.Value,
			Href:       factRef,
			ContextRef: contextRefAttr.Value,
			UnitRef:    unitVal,
			Precision:  precisionVal,
			IsNil:      nilVal,
			XMLInner:   fact.XMLInner,
		}
		ret = append(ret, newFact)
	}
	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].ID < ret[j].ID
	})
	return ret
}

func hydrateFootnoteLinks(instanceFile *serializables.InstanceFile) []FootnoteLink {
	ret := make([]FootnoteLink, 0, len(instanceFile.FootnoteLink))
	for _, footnoteLink := range instanceFile.FootnoteLink {
		item := FootnoteLink{}
		nsAttr := attr.FindAttr(footnoteLink.XMLAttrs, "xmlns")
		if nsAttr == nil || nsAttr.Value != attr.XLINK {
			continue
		}
		typeAttr := attr.FindAttr(footnoteLink.XMLAttrs, "type")
		if typeAttr == nil || typeAttr.Value != "extended" {
			continue
		}
		roleAttr := attr.FindAttr(footnoteLink.XMLAttrs, "xmlns")
		if roleAttr == nil || roleAttr.Value != attr.ROLELINK {
			continue
		}
		titleAttr := attr.FindAttr(footnoteLink.XMLAttrs, "title")
		if titleAttr != nil {
			item.Title = titleAttr.Value
		}
		item.Locs = make([]FootnoteLinkLoc, 0, len(footnoteLink.Loc))
		for _, loc := range footnoteLink.Loc {
			loctypeAttr := attr.FindAttr(loc.XMLAttrs, "type")
			if loctypeAttr == nil || loctypeAttr.Value != "locator" {
				continue
			}
			locnsAttr := attr.FindAttr(loc.XMLAttrs, "xmlns")
			if locnsAttr == nil || locnsAttr.Value != attr.XLINK {
				continue
			}
			hrefAttr := attr.FindAttr(loc.XMLAttrs, "href")
			if hrefAttr == nil || hrefAttr.Value == "" {
				continue
			}
			labelAttr := attr.FindAttr(loc.XMLAttrs, "label")
			if labelAttr == nil || labelAttr.Value == "" {
				continue
			}
			newLoc := FootnoteLinkLoc{
				Href:  hrefAttr.Value,
				Label: labelAttr.Value,
			}
			item.Locs = append(item.Locs, newLoc)
		}
		item.Footnotes = make([]Footnote, 0, len(footnoteLink.Footnote))
		for _, footnote := range footnoteLink.Footnote {
			footnotetypeAttr := attr.FindAttr(footnote.XMLAttrs, "type")
			if footnotetypeAttr == nil || footnotetypeAttr.Value != "resource" {
				continue
			}
			footnoteidAttr := attr.FindAttr(footnote.XMLAttrs, "id")
			if footnoteidAttr == nil || footnoteidAttr.Value == "" {
				continue
			}
			footnotensAttr := attr.FindAttr(footnote.XMLAttrs, "xmlns")
			if footnotensAttr == nil || footnotensAttr.Value != attr.LINK {
				continue
			}
			footnotelabelAttr := attr.FindAttr(footnote.XMLAttrs, "label")
			if footnotelabelAttr == nil || footnotelabelAttr.Value == "" {
				continue
			}
			footnotelangAttr := attr.FindAttr(footnote.XMLAttrs, "lang")
			if footnotelangAttr == nil || footnotelangAttr.Value == "" {
				continue
			}
			newFootnote := Footnote{
				ID:       footnoteidAttr.Value,
				Label:    footnotelabelAttr.Value,
				Lang:     footnotelangAttr.Value,
				CharData: footnote.CharData,
			}
			item.Footnotes = append(item.Footnotes, newFootnote)
		}
		item.FootnoteArcs = make([]FootnoteArc, 0, len(footnoteLink.FootnoteArc))
		for _, footnoteArc := range footnoteLink.FootnoteArc {
			footnoteArcnsAttr := attr.FindAttr(footnoteArc.XMLAttrs, "xmlns")
			if footnoteArcnsAttr == nil || footnoteArcnsAttr.Value != attr.LINK {
				continue
			}
			footnoteArcarcroleAttr := attr.FindAttr(footnoteArc.XMLAttrs, "arcrole")
			if footnoteArcarcroleAttr == nil || footnoteArcarcroleAttr.Value != attr.FactFootnoteArcrole {
				continue
			}
			footnoteArctypeAttr := attr.FindAttr(footnoteArc.XMLAttrs, "type")
			if footnoteArctypeAttr == nil || footnoteArctypeAttr.Value != "arc" {
				continue
			}
			footnoteArcfromAttr := attr.FindAttr(footnoteArc.XMLAttrs, "from")
			if footnoteArcfromAttr == nil || footnoteArcfromAttr.Value == "" {
				continue
			}
			footnoteArctoAttr := attr.FindAttr(footnoteArc.XMLAttrs, "to")
			if footnoteArctoAttr == nil || footnoteArctoAttr.Value == "" {
				continue
			}
			newFootnoteArc := FootnoteArc{
				Arcrole: footnoteArcarcroleAttr.Value,
				From:    footnoteArcfromAttr.Value,
				To:      footnoteArctoAttr.Value,
			}
			item.FootnoteArcs = append(item.FootnoteArcs, newFootnoteArc)
		}
		ret = append(ret, item)
	}
	sort.SliceStable(ret, func(i int, j int) bool {
		return ret[i].Title < ret[j].Title
	})
	return ret
}

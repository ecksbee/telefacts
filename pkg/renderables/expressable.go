package renderables

import (
	"encoding/json"

	"ecksbee.com/telefacts/pkg/hydratables"
)

type Expressable struct {
	Href    string
	Labels  LabelPack
	Context struct {
		Period LanguagePack
		VoidQuadrant
		ContextualMemberGrid
	}
	Expression *MultilingualFact
	Footnotes  []string
}

func MarshalExpressable(name string, contextref string, h *hydratables.Hydratable) ([]byte, error) {
	xmlname, found := h.Document.NamespaceMap[name]
	if !found {
		return nil, nil
	}
	_, found = h.Document.ContextRefMap[contextref]
	if !found {
		return nil, nil
	}
	var extracted *hydratables.Instance
	for _, instance := range h.Instances {
		extracted = &instance
		break
	}
	context := getContext(extracted, contextref)
	period := periodString(context)
	entity := stringify(&Entity{
		Scheme:   context.Entity.Identifier.Scheme,
		CharData: context.Entity.Identifier.CharData,
	})
	href, _, err := h.NameQuery(xmlname.Space, xmlname.Local)
	if err != nil {
		return nil, err
	}
	hydratedFact := h.FindFact(href, contextref)
	if hydratedFact == nil {
		return nil, nil
	}
	relevantContexts, segment, scenario, _ := getRelevantContexts(entity, h, []string{
		href,
	})
	expressedContexts := make([]relevantContext, 0)
	for _, relevantContext := range relevantContexts {
		if relevantContext.ContextRef == contextref {
			expressedContexts = append(expressedContexts, relevantContext)
			break
		}
	}
	memberGrid, voidQuadrant := getMemberGridAndVoidQuadrant(expressedContexts, segment, scenario)
	footnotes := h.GetFootnotes(hydratedFact)
	footnoteTexts := make([]string, 0, len(footnotes))
	for _, footnote := range footnotes {
		footnoteTexts = append(footnoteTexts, footnote.InnerHtml)
	}
	return json.Marshal(Expressable{
		Href:   href,
		Labels: GetLabel(h, href),
		Context: struct {
			Period LanguagePack
			VoidQuadrant
			ContextualMemberGrid
		}{
			Period:               period,
			VoidQuadrant:         voidQuadrant,
			ContextualMemberGrid: memberGrid,
		},
		Expression: render(hydratedFact, h, h, []Lang{PureLabel}),
		Footnotes:  footnoteTexts,
	})
}

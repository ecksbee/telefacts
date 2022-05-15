package renderables

import "ecksbee.com/telefacts/pkg/hydratables"

type Expressable struct {
	Href    string
	Labels  LabelPack
	Context struct {
		Period LanguagePack
		VoidQuadrant
		ContextualMemberGrid
	}
	Measurement string
	Precision   string
	Footnotes   []string
	Value       string
}

func getExpressions(h *hydratables.Hydratable, conceptFinder ConceptFinder) (map[string]Expressable, error) {
	ret := make(map[string]Expressable)
	source := h.Document
	if source != nil {
		for _, item := range source.Expressions {
			id := item.ID
			href, _, err := conceptFinder.NameQuery(item.Name.Space, item.Name.Local)
			if err != nil {
				return nil, err
			}
			if id == "" {
				continue
			}
			var extracted *hydratables.Instance
			context := getContext(extracted, item.ContextRef)
			period := periodString(context)
			entity := stringify(&Entity{
				Scheme:   context.Entity.Identifier.Scheme,
				CharData: context.Entity.Identifier.CharData,
			})
			relevantContexts, segment, scenario, _ := getRelevantContexts(entity, h, []string{
				href,
			})
			expressedContexts := make([]relevantContext, 0)
			for _, relevantContext := range relevantContexts {
				if relevantContext.ContextRef == item.ContextRef {
					expressedContexts = append(expressedContexts, relevantContext)
					break
				}
			}
			memberGrid, voidQuadrant := getMemberGridAndVoidQuadrant(expressedContexts, segment, scenario)
			ret[id] = Expressable{
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
				Measurement: "//todo", //todo
				Precision:   "",
				Footnotes:   make([]string, 0),
				Value:       "",
			}
		}
	}
	return ret, nil
}

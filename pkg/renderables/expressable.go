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

func getExpressions(h *hydratables.Hydratable) (map[string]Expressable, error) {
	ret := make(map[string]Expressable)
	source := h.Document
	if source != nil {
		for _, item := range source.Expressions {
			id := item.ID
			if id == "" {
				continue
			}
			ret[id] = Expressable{
				//todo
			}
		}
	}
	return ret, nil
}

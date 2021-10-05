package hydratables

import (
	"strings"

	"ecksbee.com/telefacts/internal/graph"
	"ecksbee.com/telefacts/pkg/attr"
)

type Fact struct {
	Href       string
	ID         string
	ContextRef string
	UnitRef    string
	Precision  Precision
	IsNil      bool
	XMLInner   string
}

func (h *Hydratable) FindFact(href string, contextRef string) *Fact {
	for _, ins := range h.Instances {
		for _, fact := range ins.Facts {
			if fact.Href == href && fact.ContextRef == contextRef {
				return &fact
			}
		}
	}
	return nil
}

func fnoteArcs(fnoteArcs []FootnoteArc) []graph.Arc {
	ret := make([]graph.Arc, 0, len(fnoteArcs))
	for i, fnoteArc := range fnoteArcs {
		ret = append(ret, graph.Arc{
			Arcrole: fnoteArc.Arcrole,
			Order:   float64(i),
			From:    fnoteArc.From,
			To:      fnoteArc.To,
		})
	}
	return ret
}

func (h *Hydratable) GetFootnotes(fact *Fact) []*Footnote {
	if fact == nil {
		return make([]*Footnote, 0)
	}
	for _, instance := range h.Instances {
		for _, footnoteLink := range instance.FootnoteLinks {
			arcs := footnoteLink.FootnoteArcs
			fnoteArcs := fnoteArcs(arcs)
			root := graph.Tree(fnoteArcs, attr.FactFootnoteArcrole)
			ret := make([]*Footnote, 0, len(arcs))
			if len(root.Children) <= 0 {
				return nil
			}
			for _, child := range root.Children {
				href := mapFnoteLocatorToHref(&instance, child.Locator)
				locatedID := href
				pos := strings.LastIndex(href, "#")
				if pos > -1 {
					adjustedPos := pos + 1
					if adjustedPos < len(href) {
						locatedID = href[adjustedPos:]
					}
				}
				if locatedID == fact.ID {
					for _, grandchild := range child.Children {
						footnoteHref := mapFnoteLocatorToHref(&instance, grandchild.Locator)
						locatedFnoteID := footnoteHref
						pos := strings.LastIndex(footnoteHref, "#")
						if pos > -1 {
							adjustedPos := pos + 1
							if adjustedPos < len(footnoteHref) {
								locatedFnoteID = footnoteHref[adjustedPos:]
							}
						}
						locatedFootnote := h.FindFootnote(locatedFnoteID)
						if locatedFootnote != nil {
							ret = append(ret, locatedFootnote)
						}
					}
					return ret
				}
			}
		}
	}
	return nil
}

func mapFnoteLocatorToHref(instance *Instance, locator string) string {
	footnoteLinks := instance.FootnoteLinks
	for _, footnoteLink := range footnoteLinks {
		for _, loc := range footnoteLink.Locs {
			if loc.Label == locator {
				i := strings.Index(loc.Href, "#")
				if i >= 0 {
					return loc.Href
				}
			}
		}
	}
	return "#" + locator
}

func (h *Hydratable) FindFootnote(id string) *Footnote {
	for _, ins := range h.Instances {
		for _, footnoteLink := range ins.FootnoteLinks {
			for _, footnote := range footnoteLink.Footnotes {
				if footnote.ID == id {
					return &footnote
				}
			}
		}
	}
	return nil
}

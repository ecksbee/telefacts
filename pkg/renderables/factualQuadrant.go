package renderables

import (
	"ecksbee.com/telefacts/pkg/hydratables"
)

type FactExpression struct {
	Head      string
	Core      string
	Tail      string
	CharData  string
	TextBlock string
}

type MultilingualFact map[Lang]FactExpression

type FactualQuadrant [][]*MultilingualFact

func getFactualQuadrant(hrefs []string, relevantContexts []relevantContext,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder,
	langs []Lang) FactualQuadrant {
	rowCount := len(hrefs)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return FactualQuadrant{}
	}
	ret := make([][]*MultilingualFact, rowCount)
	for i := 0; i < rowCount; i++ {
		row := make([]*MultilingualFact, colCount)
		href := hrefs[i]
		for j := 0; j < colCount; j++ {
			var fact *hydratables.Fact
			contextRef := relevantContexts[j].ContextRef
			fact = factFinder.FindFact(href, contextRef)
			row[j] = render(fact, conceptFinder, measurementFinder, langs)
		}
		ret[i] = row
	}
	return ret
}

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
	langs []Lang) (FactualQuadrant, [][][]int, []string) {
	rowCount := len(hrefs)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return FactualQuadrant{}, nil, nil
	}
	ret := make([][]*MultilingualFact, rowCount)
	footnoteGrid := make([][][]*hydratables.Footnote, rowCount)
	for i := 0; i < rowCount; i++ {
		row := make([]*MultilingualFact, colCount)
		footnoteRow := make([][]*hydratables.Footnote, colCount)
		href := hrefs[i]
		for j := 0; j < colCount; j++ {
			var fact *hydratables.Fact
			var footnotes []*hydratables.Footnote
			contextRef := relevantContexts[j].ContextRef
			fact = factFinder.FindFact(href, contextRef)
			footnotes = factFinder.GetFootnotes(fact)
			row[j] = render(fact, conceptFinder, measurementFinder, langs)
			footnoteRow[j] = footnotes
		}
		ret[i] = row
		footnoteGrid[i] = footnoteRow
	}
	grid := make([][][]int, rowCount)
	arr := make([]string, 0)
	//todo fill grid and arr
	return ret, grid, arr
}

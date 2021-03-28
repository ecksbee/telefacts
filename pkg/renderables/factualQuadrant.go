package renderables

import (
	"ecksbee.com/telefacts/pkg/hydratables"
)

type FactExpression struct {
	Head        string
	Core        string
	Tail        string
	TextPreview string
	TextBlock   string
}

type MultilingualFact map[LabelRole]map[Lang]FactExpression

type FactualQuadrant [][]MultilingualFact

func getFactualQuadrant(hrefs []string, relevantContexts []RelevantContext,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder,
	labelRoles []LabelRole, langs []Lang) FactualQuadrant {
	rowCount := len(hrefs)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return FactualQuadrant{}
	}
	var ret [][]MultilingualFact
	for i := 0; i < rowCount; i++ {
		var row []MultilingualFact
		href := hrefs[i]
		for j := 0; j < colCount; j++ {
			var fact *hydratables.Fact
			contextRef := relevantContexts[j].ContextRef
			fact = factFinder.FindFact(href, contextRef)
			row = append(row, render(fact, conceptFinder, measurementFinder, labelRoles, langs))
		}
		ret = append(ret, row)
	}
	return ret
}

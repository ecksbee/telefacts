package renderables

import (
	"ecksbee.com/telefacts/hydratables"
)

func getFactualQuadrant(hrefs []string, relevantContexts []RelevantContext,
	factFinder FactFinder, measurementFinder MeasurementFinder,
	labelRoles []LabelRole, langs []Lang) [][]LabelPack {
	rowCount := len(hrefs)
	colCount := len(relevantContexts)
	if rowCount <= 0 || colCount <= 0 {
		return [][]LabelPack{{}}
	}
	var ret [][]LabelPack
	for i := 0; i < rowCount; i++ {
		var row []LabelPack
		href := hrefs[i]
		for j := 0; j < colCount; j++ {
			var fact *hydratables.Fact
			contextRef := relevantContexts[j].ContextRef
			fact = factFinder.FindFact(href, contextRef)
			row = append(row, render(fact, measurementFinder, labelRoles, langs))
		}
		ret = append(ret, row)
	}
	return ret
}

package renderables

import "ecksbee.com/telefacts/hydratables"

type FactFinder interface {
	FindFact(href string, contextRef string) *hydratables.Fact
}

func render(fact *hydratables.Fact) string {
	if fact == nil {
		return "null"
	}
	var precision string
	if fact.Decimals != "" {
		precision = fact.Decimals
	} else {
		precision = fact.Precision
	}
	return precision + " " + fact.XMLInner + " " + fact.UnitRef
}

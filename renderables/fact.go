package renderables

import "ecksbee.com/telefacts/hydratables"

type FactFinder interface {
	FindFact(href string, contextRef string) *hydratables.Fact
}

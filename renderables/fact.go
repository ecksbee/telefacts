package renderables

import "ecks-bee.com/telefacts/xbrl"

type FactFinder interface {
	FindFact(href string, contextRef string) *xbrl.Fact
}

package renderables

import (
	"sort"

	"ecks-bee.com/telefacts/xbrl"
)

func sortedRelationshipSets(schema *xbrl.Schema) []string {
	linkroleURIs := dedupRelationshipSets(schema)
	sort.SliceStable(linkroleURIs, func(i, j int) bool {
		return linkroleURIs[i] < linkroleURIs[j]
	})
	return linkroleURIs
}

func dedupRelationshipSets(schema *xbrl.Schema) []string {
	if len(schema.Annotation) <= 0 {
		return []string{}
	}
	if len(schema.Annotation[0].Appinfo.RoleType) <= 0 {
		return []string{}
	}
	rsets := func(s *xbrl.Schema) []string {
		ret := []string{}
		slice := schema.Annotation[0].Appinfo.RoleType
		for _, e := range slice {
			if len(e.RoleURI) <= 0 {
				continue
			}
			ret = append(ret, e.RoleURI)
		}
		return ret
	}(schema)
	uniques := dedup(rsets)
	return uniques
}

func dedup(arr []string) []string {
	occured := map[string]bool{}
	ret := []string{}
	for e := range arr {
		if occured[arr[e]] != true {
			occured[arr[e]] = true
			ret = append(ret, arr[e])
		}
	}
	return ret
}

type locatorNode struct {
	Locator  string
	Order    float64
	Children []*locatorNode
}

func find(node *locatorNode, loc string) (*locatorNode, int) {
	if node.Locator == loc {
		return node, -1
	}
	for i, c := range node.Children {
		ret, _ := find(c, loc)
		if ret != nil {
			return ret, i
		}
	}
	return nil, -1
}

func render(fact *xbrl.Fact) string {
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

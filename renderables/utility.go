package renderables

import (
	"sort"
	"strings"

	"ecksbee.com/telefacts/hydratables"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sortedRelationshipSets(h *hydratables.Hydratable) []RelationshipSet {
	rsets := dedupRelationshipSets(h)
	sort.SliceStable(rsets, func(i, j int) bool {
		if rsets[i].Title == rsets[j].Title {
			return rsets[i].RoleURI < rsets[j].RoleURI
		}
		return rsets[i].Title < rsets[j].Title
	})
	return rsets
}

func dedupRelationshipSets(h *hydratables.Hydratable) []RelationshipSet {
	rsets := []RelationshipSet{}
	for _, schema := range h.Schemas {
		if len(schema.Annotation.Appinfo.RoleTypes) <= 0 {
			continue
		}
		for _, e := range schema.Annotation.Appinfo.RoleTypes {
			if len(e.RoleURI) <= 0 {
				continue
			}
			rsets = append(rsets, RelationshipSet{
				RoleURI: e.RoleURI,
				Title:   e.Definition,
			})
		}
	}
	uniques := func(arr []RelationshipSet) []RelationshipSet {
		occured := map[RelationshipSet]bool{}
		u := []RelationshipSet{}
		for e := range arr {
			if occured[arr[e]] != true {
				occured[arr[e]] = true
				u = append(u, arr[e])
			}
		}
		return u
	}(rsets)
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

type arc struct {
	Arcrole string
	Order   float64
	From    string
	To      string
}

func tree(arcs []arc, arcrole string) locatorNode {
	var root locatorNode
	root.Children = make([]*locatorNode, 0, len(arcs))
	sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
	for _, arc := range arcs {
		if arc.Arcrole == arcrole {
			from, _ := find(&root, arc.From)
			if from != nil {
				to, toIndex := find(&root, arc.To)
				if to != nil {
					root.Children[toIndex] = root.Children[len(root.Children)-1]
					root.Children = root.Children[:len(root.Children)-1]
					from.Children = append(from.Children, to)
				} else {
					order := arc.Order
					from.Children = append(from.Children, &locatorNode{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*locatorNode, 0, len(arcs)),
					})
				}
			} else {
				from = &locatorNode{
					Locator:  arc.From,
					Children: make([]*locatorNode, 0, len(arcs)),
				}
				root.Children = append(root.Children, from)
				to, toIndex := find(&root, arc.To)
				if to != nil {
					root.Children[toIndex] = root.Children[len(root.Children)-1]
					root.Children = root.Children[:len(root.Children)-1]
					from.Children = append(from.Children, to)
				} else {
					order := arc.Order
					from.Children = append(from.Children, &locatorNode{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*locatorNode, 0, len(arcs)),
					})
				}
			}
		}
	}
	return root
}

func comparePath(a path, b path) int {
	max := len(a)
	if len(a) > len(b) {
		max = len(b)
	}
	for i := 0; i < max; i++ {
		aa := a[i]
		bb := b[i]
		c := strings.Compare(aa, bb)
		if c != 0 {
			return c
		}
	}
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	return 0
}

type path []string

func paths(node *locatorNode, prior path) []path {
	if node == nil {
		return []path{}
	}
	newPath := append(prior, node.Locator)
	if len(node.Children) <= 0 {
		return []path{
			newPath,
		}
	}
	var ret []path
	for _, child := range node.Children {
		for _, childPath := range paths(child, newPath) {
			ret = append(ret, childPath)
		}
	}
	return ret
}

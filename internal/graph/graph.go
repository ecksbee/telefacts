package graph

import (
	myarcs "github.com/joshuanario/arcs"
)

// type LocatorNode struct {
// 	Locator  string
// 	Order    float64
// 	Children []*LocatorNode
// }

// func Find(node *LocatorNode, loc string) (*LocatorNode, int) {
// 	if node.Locator == loc {
// 		return node, -1
// 	}
// 	for i, c := range node.Children {
// 		ret, _ := Find(c, loc)
// 		if ret != nil {
// 			return ret, i
// 		}
// 	}
// 	return nil, -1
// }

// type Arc struct {
// 	Arcrole string
// 	Order   float64
// 	From    string
// 	To      string
// }

func Tree(arcs []myarcs.Arc, arcrole string) *myarcs.RArc {
	return myarcs.NewRArc(arcs, arcrole)
}

// type Path []string

// func Paths(node *LocatorNode, prior Path) []Path {
// 	if node == nil {
// 		return []Path{}
// 	}
// 	newPath := append(prior, node.Locator)
// 	if len(node.Children) <= 0 {
// 		return []Path{
// 			newPath,
// 		}
// 	}
// 	var ret []Path
// 	for _, child := range node.Children {
// 		ret = append(ret, Paths(child, newPath)...)
// 	}
// 	return ret
// }

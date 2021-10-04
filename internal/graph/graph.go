package graph

import "sort"

type LocatorNode struct {
	Locator  string
	Order    float64
	Children []*LocatorNode
}

func Find(node *LocatorNode, loc string) (*LocatorNode, int) {
	if node.Locator == loc {
		return node, -1
	}
	for i, c := range node.Children {
		ret, _ := Find(c, loc)
		if ret != nil {
			return ret, i
		}
	}
	return nil, -1
}

type Arc struct {
	Arcrole string
	Order   float64
	From    string
	To      string
}

func Tree(arcs []Arc, arcrole string) LocatorNode {
	var root LocatorNode
	root.Children = make([]*LocatorNode, 0, len(arcs))
	sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
	for _, arc := range arcs {
		if arc.Arcrole == arcrole {
			from, _ := Find(&root, arc.From)
			if from != nil {
				to, toIndex := Find(&root, arc.To)
				if to != nil {
					root.Children[toIndex] = root.Children[len(root.Children)-1]
					root.Children = root.Children[:len(root.Children)-1]
					from.Children = append(from.Children, to)
				} else {
					order := arc.Order
					from.Children = append(from.Children, &LocatorNode{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*LocatorNode, 0, len(arcs)),
					})
				}
			} else {
				from = &LocatorNode{
					Locator:  arc.From,
					Children: make([]*LocatorNode, 0, len(arcs)),
				}
				root.Children = append(root.Children, from)
				to, toIndex := Find(&root, arc.To)
				if to != nil {
					root.Children[toIndex] = root.Children[len(root.Children)-1]
					root.Children = root.Children[:len(root.Children)-1]
					from.Children = append(from.Children, to)
				} else {
					order := arc.Order
					from.Children = append(from.Children, &LocatorNode{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*LocatorNode, 0, len(arcs)),
					})
				}
			}
		}
	}
	return root
}

type Path []string

func Paths(node *LocatorNode, prior Path) []Path {
	if node == nil {
		return []Path{}
	}
	newPath := append(prior, node.Locator)
	if len(node.Children) <= 0 {
		return []Path{
			newPath,
		}
	}
	var ret []Path
	for _, child := range node.Children {
		ret = append(ret, Paths(child, newPath)...)
	}
	return ret
}

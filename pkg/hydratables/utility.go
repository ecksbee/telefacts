package hydratables

import "sort"

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

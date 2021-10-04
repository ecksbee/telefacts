package telefacts_test

import (
	"testing"

	"ecksbee.com/telefacts/internal/graph"
)

func TestPaths(t *testing.T) {
	b := &graph.LocatorNode{
		Locator: "B",
		Order:   1,
	}
	c := &graph.LocatorNode{
		Locator: "C",
		Order:   1,
		Children: []*graph.LocatorNode{
			b,
		},
	}
	stimulus := &graph.LocatorNode{
		Locator: "A",
		Order:   1,
		Children: []*graph.LocatorNode{
			b,
			c,
		},
	}
	start := graph.Path{}
	output := graph.Paths(stimulus, start)
	if len(output) != 2 {
		t.Fatalf("expected 2 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestManyPaths(t *testing.T) {
	f := &graph.LocatorNode{
		Locator: "F",
		Order:   1,
	}
	d := &graph.LocatorNode{
		Locator: "D",
		Order:   1,
	}
	b := &graph.LocatorNode{
		Locator: "B",
		Order:   1,
		Children: []*graph.LocatorNode{
			d,
		},
	}
	c := &graph.LocatorNode{
		Locator: "C",
		Order:   2,
		Children: []*graph.LocatorNode{
			b,
			f,
		},
	}
	a := &graph.LocatorNode{
		Locator: "A",
		Order:   1,
		Children: []*graph.LocatorNode{
			b,
			d,
		},
	}
	e := &graph.LocatorNode{
		Locator: "F",
		Order:   3,
		Children: []*graph.LocatorNode{
			d,
		},
	}
	stimulus := &graph.LocatorNode{
		Locator: "root",
		Order:   1,
		Children: []*graph.LocatorNode{
			a,
			c,
			e,
		},
	}
	start := graph.Path{}
	output := graph.Paths(stimulus, start)
	if len(output) != 5 {
		t.Fatalf("expected 5 paths; outcome %d;\n%v\n", len(output), output)
	}
}

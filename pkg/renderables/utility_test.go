package renderables

import (
	"testing"
)

func TestPaths(t *testing.T) {
	b := &locatorNode{
		Locator: "B",
		Order:   1,
	}
	c := &locatorNode{
		Locator: "C",
		Order:   1,
		Children: []*locatorNode{
			b,
		},
	}
	stimulus := &locatorNode{
		Locator: "A",
		Order:   1,
		Children: []*locatorNode{
			b,
			c,
		},
	}
	start := path{}
	output := paths(stimulus, start)
	if len(output) != 2 {
		t.Fatalf("expected 2 paths; outcome %d;\n%v\n", len(output), output)
	}
}

func TestManyPaths(t *testing.T) {
	f := &locatorNode{
		Locator: "F",
		Order:   1,
	}
	d := &locatorNode{
		Locator: "D",
		Order:   1,
	}
	b := &locatorNode{
		Locator: "B",
		Order:   1,
		Children: []*locatorNode{
			d,
		},
	}
	c := &locatorNode{
		Locator: "C",
		Order:   2,
		Children: []*locatorNode{
			b,
			f,
		},
	}
	a := &locatorNode{
		Locator: "A",
		Order:   1,
		Children: []*locatorNode{
			b,
			d,
		},
	}
	e := &locatorNode{
		Locator: "F",
		Order:   3,
		Children: []*locatorNode{
			d,
		},
	}
	stimulus := &locatorNode{
		Locator: "root",
		Order:   1,
		Children: []*locatorNode{
			a,
			c,
			e,
		},
	}
	start := path{}
	output := paths(stimulus, start)
	if len(output) != 5 {
		t.Fatalf("expected 5 paths; outcome %d;\n%v\n", len(output), output)
	}
}

package telefacts_test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestMarshalRenderable_Erroneous_Footnote(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = path.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(serializables.WorkingDirectoryPath, "folders", "test_erroneous")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("test_erroneous")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	slug := "06f7b4f73370e6a982e34df6afcff503"
	data, err := renderables.MarshalRenderable(slug, h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	r := renderables.Renderable{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if r.RelationshipSet.RoleURI != "http://nationalbeverage.com/20200502/role/statement-note-10-stockbased-compensation-summary-of-stock-option-activity-details" {
		t.Fatalf("expected http://nationalbeverage.com/20200502/role/statement-note-10-stockbased-compensation-summary-of-stock-option-activity-details; outcome %s;\n", r.RelationshipSet.RoleURI)
	}
	if (*r.PGrid.FactualQuadrant[6][4])["Unlabelled"].Core != "11.14" {
		t.Fatalf("expected 11.14; outcome %s;\n", (*r.PGrid.FactualQuadrant[6][4])["Unlabelled"].Core)
	}
	if len(r.PGrid.FootnoteGrid[6][4]) != 1 {
		t.Fatalf("expected 1 footnote link; outcome %d;\n", len(r.PGrid.FootnoteGrid[6][4]))
	}
	if len(r.PGrid.Footnotes) != 1 {
		t.Fatalf("expected 1 footnote; outcome %d;\n", len(r.PGrid.Footnotes))
	}
	if r.PGrid.Footnotes[0] != "Weighted average exercise price." {
		t.Fatalf("expected Weighted average exercise price.; outcome %s;\n", r.PGrid.Footnotes[0])
	}
}

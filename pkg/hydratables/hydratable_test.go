package hydratables_test

import (
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestHydrate_Gold(t *testing.T) {
	scache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.SetGlobalDir(path.Join("..", "..", "taxonomies"))
	serializables.InjectCache(scache)
	hydratables.InjectCache(hcache)
	workingDir := path.Join("..", "test", "test_gold")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	entryFilePath := "wk-20200930_htm.xml"
	f, err := serializables.Discover(workingDir, entryFilePath)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if len(h.Instances) != 1 {
		t.Fatalf("expected 1 Instance; outcome %d;\n", len(h.Instances))
	}
	ins, found := h.Instances[entryFilePath]
	if !found {
		t.Fatalf("expected %s Instance to be found;\n", entryFilePath)
	}
	if len(ins.Facts) != 874 {
		t.Fatalf("expected 874 Facts; outcome %d;\n", len(ins.Facts))
	}
	if len(ins.Units) != 4 {
		t.Fatalf("expected 4 Units; outcome %d;\n", len(ins.Units))
	}
	if len(ins.Contexts) != 248 {
		t.Fatalf("expected 248 Context; outcome %d;\n", len(ins.Contexts))
	}
	if len(ins.FootnoteLinks) != 0 {
		t.Fatalf("expected 0 FootnoteLink; outcome %d;\n", len(ins.FootnoteLinks))
	}

	if len(h.Schemas) != 1 {
		t.Fatalf("expected 1 Schema; outcome %d;\n", len(h.Schemas))
	}
	schemaFilePath := "wk-20200930.xsd"
	schema, found := h.Schemas[schemaFilePath]
	if !found {
		t.Fatalf("expected %s Schema to be found;\n", schemaFilePath)
	}
	if len(schema.Annotation.Appinfo.RoleTypes) != 42 {
		t.Fatalf("expected 42 Relationship Set; outcome %d;\n", len(schema.Annotation.Appinfo.RoleTypes))
	}
	if len(schema.Element) != 26 {
		t.Fatalf("expected 26 extension Concept; outcome %d;\n", len(schema.Element))
	}

	preFilePath := "wk-20200930_pre.xml"
	pre, found := h.PresentationLinkbases[preFilePath]
	if !found {
		t.Fatalf("expected %s PresentationLinkbase to be found;\n", preFilePath)
	}
	if len(pre.RoleRefs) != 42 {
		t.Fatalf("expected 42 Relationship Set reference; outcome %d;\n", len(pre.RoleRefs))
	}
	if len(pre.PresentationLinks) != 42 {
		t.Fatalf("expected 42 presentation Network; outcome %d;\n", len(pre.PresentationLinks))
	}

	defFilePath := "wk-20200930_def.xml"
	def, found := h.DefinitionLinkbases[defFilePath]
	if !found {
		t.Fatalf("expected %s DefinitionLinkbase to be found;\n", defFilePath)
	}
	if len(def.RoleRefs) != 41 {
		t.Fatalf("expected 41 Relationship Set reference; outcome %d;\n", len(def.RoleRefs))
	}
	if len(def.DefinitionLinks) != 41 {
		t.Fatalf("expected 41 definition Network; outcome %d;\n", len(def.DefinitionLinks))
	}

	calFilePath := "wk-20200930_cal.xml"
	cal, found := h.CalculationLinkbases[calFilePath]
	if !found {
		t.Fatalf("expected %s CalculationLinkbase to be found;\n", calFilePath)
	}
	if len(cal.RoleRefs) != 42 {
		t.Fatalf("expected 42 Relationship Set reference; outcome %d;\n", len(cal.RoleRefs))
	}
	if len(cal.CalculationLinks) != 42 {
		t.Fatalf("expected 42 calculation Network; outcome %d;\n", len(cal.CalculationLinks))
	}

	labFilePath := "wk-20200930_lab.xml"
	lab, found := h.LabelLinkbases[labFilePath]
	if !found {
		t.Fatalf("expected %s LabelLinkbase to be found;\n", labFilePath)
	}
	if len(lab.RoleRefs) != 7 {
		t.Fatalf("expected 7 Relationship Set reference; outcome %d;\n", len(lab.RoleRefs))
	}
	if len(lab.LabelLink) != 1 {
		t.Fatalf("expected 1 label Network; outcome %d;\n", len(lab.LabelLink))
	}
}

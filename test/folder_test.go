package telefacts_test

import (
	"os"
	"path"
	"testing"
	"time"

	"ecksbee.com/telefacts/pkg/sec"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestDiscover_Gold(t *testing.T) {
	scache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	seccache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.SetGlobalDir(path.Join(".", "data", "taxonomies"))
	serializables.InjectCache(scache)
	sec.InjectCache(seccache)
	workingDir := path.Join(".", "data", "test_gold")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	entryFilePath := "wk-20200930_htm.xml"
	f, err := serializables.Discover(workingDir, entryFilePath)
	time.Sleep(time.Second * 15)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if f.Dir != workingDir {
		t.Fatalf("expected %s Dir; outcome %s;\n", workingDir, f.Dir)
	}
	if len(f.Instances) != 1 {
		t.Fatalf("expected 1 Instance; outcome %d;\n", len(f.Instances))
	}
	ins, found := f.Instances[entryFilePath]
	if !found {
		t.Fatalf("expected %s Instance to be found;\n", entryFilePath)
	}
	if len(ins.SchemaRef) != 1 {
		t.Fatalf("expected 1 SchemaRef; outcome %d;\n", len(ins.SchemaRef))
	}
	if len(ins.Context) != 248 {
		t.Fatalf("expected 248 Context; outcome %d;\n", len(ins.Context))
	}
	if len(ins.Facts) != 874 {
		t.Fatalf("expected 874 Fact; outcome %d;\n", len(ins.Facts))
	}

	if len(f.Schemas) != 1 {
		t.Fatalf("expected 1 Schema; outcome %d;\n", len(f.Schemas))
	}

	if len(f.PresentationLinkbases) != 1 {
		t.Fatalf("expected 1 PresentationLinkbase; outcome %d;\n", len(f.PresentationLinkbases))
	}

	if len(f.DefinitionLinkbases) != 1 {
		t.Fatalf("expected 1 DefinitionLinkbase; outcome %d;\n", len(f.DefinitionLinkbases))
	}

	if len(f.CalculationLinkbases) != 1 {
		t.Fatalf("expected 1 CalculationLinkbase; outcome %d;\n", len(f.CalculationLinkbases))
	}

	if len(f.LabelLinkbases) != 1 {
		t.Fatalf("expected 1 LabelLinkbase; outcome %d;\n", len(f.LabelLinkbases))
	}
}

package telefacts_test

import (
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/serializables"
)

func TestDiscover_Gold(t *testing.T) {
	serializables.VolumePath = path.Join(".", "data")
	workingDir := path.Join(serializables.VolumePath, "folders", "test_gold")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	entryFilePath := "wk-20200930_htm.xml"
	f, err := serializables.Discover("test_gold")
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

func TestDiscover_Ixbrl(t *testing.T) {
	serializables.VolumePath = path.Join(".", "data")
	workingDir := path.Join(serializables.VolumePath, "folders", "test_ixbrl")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	entryFilePath := "wk-20200930.htm"
	f, err := serializables.Discover("test_ixbrl")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	ixbrl, found := f.IxbrlFiles[entryFilePath]
	if !found {
		t.Fatalf("expected %s Ixbrl file to be found;\n", entryFilePath)
	}
	if len(ixbrl.Header.References.SchemaRef) != 1 {
		t.Fatalf("expected 1 schemaRef; outcome %d", len(ixbrl.Header.References.SchemaRef))
	}
	if ixbrl.Header.References.SchemaRef[0] != "wk-20200930.xsd" {
		t.Fatalf("expected wk-20200930.xsd schemaRef; outcome %s", ixbrl.Header.References.SchemaRef[0])
	}
	if len(ixbrl.Header.Resources.Contexts) != 248 {
		t.Fatalf("expected 248 inline context; outcome %d", len(ixbrl.Header.Resources.Contexts))
	}
	if len(ixbrl.Header.Resources.Units) != 4 {
		t.Fatalf("expected 4 inline unit; outcome %d", len(ixbrl.Header.Resources.Units))
	}
	if len(ixbrl.Header.Hidden.Nonfractions) != 1 {
		t.Fatalf("expected 1 inline hidden nonfraction; outcome %d", len(ixbrl.Header.Hidden.Nonfractions))
	}
	if len(ixbrl.Header.Hidden.Nonnumerics) != 7 {
		t.Fatalf("expected 7 inline hidden nonnumeric; outcome %d", len(ixbrl.Header.Hidden.Nonnumerics))
	}
	if len(ixbrl.RenderedFacts.Nonfractions) != 798 {
		t.Fatalf("expected 798 inline rendered nonfraction; outcome %d", len(ixbrl.RenderedFacts.Nonfractions))
	}
	if len(ixbrl.RenderedFacts.Nonnumerics) != 68 {
		t.Fatalf("expected 68 inline rendered nonfraction; outcome %d", len(ixbrl.RenderedFacts.Nonnumerics))
	}
	if len(ixbrl.Header.Hidden.Footnotes) > 0 {
		t.Fatalf("expected 0 inline hidden footnote; outcome %d", len(ixbrl.Header.Hidden.Footnotes))
	}
	if len(ixbrl.RenderedFacts.Footnotes) > 0 {
		t.Fatalf("expected 0 inline rendered footnote; outcome %d", len(ixbrl.RenderedFacts.Footnotes))
	}
	if len(ixbrl.RenderedFacts.Continuations) != 28 {
		t.Fatalf("expected 27 continuations; outcome %d", len(ixbrl.RenderedFacts.Continuations))
	}

}

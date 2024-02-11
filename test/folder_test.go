package telefacts_test

import (
	"os"
	"path/filepath"
	"testing"

	"ecksbee.com/telefacts/pkg/serializables"
)

func TestDiscover_Gold(t *testing.T) {
	serializables.WorkingDirectoryPath = filepath.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = filepath.Join(".", "gts")
	workingDir := filepath.Join(serializables.WorkingDirectoryPath, "folders", "test_gold")
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

func TestDiscover_Erroneous_Images(t *testing.T) {
	serializables.WorkingDirectoryPath = filepath.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = filepath.Join(".", "gts")
	workingDir := filepath.Join(serializables.WorkingDirectoryPath, "folders", "test_erroneous")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	entryFilePath := "fizz20200502_10k_htm.xml"
	f, err := serializables.Discover("test_erroneous")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if f.Dir != workingDir {
		t.Fatalf("expected %s Dir; outcome %s;\n", workingDir, f.Dir)
	}
	_, found := f.Instances[entryFilePath]
	if !found {
		t.Fatalf("expected %s Instance to be found;\n", entryFilePath)
	}
	if len(f.Images) != 18 {
		t.Fatalf("expected 18 images; outcome %d;\n", len(f.Images))
	}
}

func TestDiscover_Image(t *testing.T) {
	serializables.WorkingDirectoryPath = filepath.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = filepath.Join(".", "gts")
	workingDir := filepath.Join(serializables.WorkingDirectoryPath, "folders", "test_image")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("test_image")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if f.Dir != workingDir {
		t.Fatalf("expected %s Dir; outcome %s;\n", workingDir, f.Dir)
	}
	if len(f.Images) != 1 {
		t.Fatalf("expected 1 images; outcome %d;\n", len(f.Images))
	}
}

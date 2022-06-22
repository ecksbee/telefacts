package telefacts_test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestCatalag_Expressables(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.VolumePath = path.Join(".", "data")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(".", "data", "folders", "test_ix")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		os.MkdirAll(workingDir, fs.FileMode(0700))
	}
	defer func() {
		os.RemoveAll(workingDir)
	}()
	zipFile := path.Join(".", "data", "test_ix.zip")
	err = unZipTestData(workingDir, zipFile)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("test_ix")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	data, err := renderables.MarshalCatalog(h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	c := renderables.Catalog{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if c.DocumentName != "cmg-20200331x10q.htm" {
		t.Fatalf("expected cmg-20200331x10q.htm; outcome %s;\n", c.DocumentName)
	}
}

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

func TestCatalog_Expressables(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = path.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(".", "wd", "folders", "test_ix")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		os.MkdirAll(workingDir, fs.FileMode(0700))
	}
	defer func() {
		os.RemoveAll(workingDir)
	}()
	zipFile := path.Join(".", "wd", "test_ix.zip")
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
	data, err = renderables.MarshalExpressable("us-gaap:EffectOfExchangeRateOnCashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents", "Duration_1_1_2020_To_3_31_2020", h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	e := renderables.Expressable{}
	err = json.Unmarshal(data, &e)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if e.Href != "http://xbrl.fasb.org/us-gaap/2019/elts/us-gaap-2019-01-31.xsd#us-gaap_EffectOfExchangeRateOnCashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents" {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2019/elts/us-gaap-2019-01-31.xsd#us-gaap_EffectOfExchangeRateOnCashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents; outcome %s;\n", e.Href)
	}
}

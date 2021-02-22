package renderables_test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/hydratables"
	"ecksbee.com/telefacts/renderables"
	"ecksbee.com/telefacts/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func names() map[string]map[string]string {
	ret := make(map[string]map[string]string)
	cik := "http://www.sec.gov/CIK"
	ret[cik] = make(map[string]string)
	ret[cik]["0001445305"] = "WORKIVA INC"
	ret[cik] = make(map[string]string)
	ret[cik]["0000069891"] = "FILER DIRECT CORP"
	ret[cik] = make(map[string]string)
	ret[cik]["0000843006"] = "NATIONAL BEVERAGE CORP"
	return ret
}

func TestMarshalRenderable_Gold_BalanceSheet(t *testing.T) {
	scache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.SetGlobalDir(path.Join("..", "taxonomies"))
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
	slug := "883459b49fae34a739704b6db51d6b1d"
	data, err := renderables.MarshalRenderable(slug, names(), h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	r := renderables.Renderable{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}

	if len(r.LabelRoles) != 2 {
		t.Fatalf("expected 2 LabelRole; outcome %d;\n", len(r.LabelRoles))
	}

	if len(r.Lang) != 2 {
		t.Fatalf("expected 2 Lang; outcome %d;\n", len(r.Lang))
	}

}

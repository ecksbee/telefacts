package telefacts_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestMarshalRenderable_High_Precision(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = filepath.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = filepath.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := filepath.Join(serializables.WorkingDirectoryPath, "folders", "high_precision")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("high_precision")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	slug := "b8d6d55ca3d34f8be977a0cab0bae837"
	data, err := renderables.MarshalRenderable(slug, h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	r := renderables.Renderable{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if r.RelationshipSet.RoleURI != "http://www.workiva.com/role/StockBasedCompensationEmployeeStockPurchasePlanDetails" {
		t.Fatalf("expected http://www.workiva.com/role/StockBasedCompensationEmployeeStockPurchasePlanDetails; outcome %s;\n", r.RelationshipSet.RoleURI)
	}
}

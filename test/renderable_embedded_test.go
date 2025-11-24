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

func TestMarshalRenderable_Embedded_BalanceSheet(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = filepath.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = filepath.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := filepath.Join(serializables.WorkingDirectoryPath, "folders", "embedded_linkbases")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("embedded_linkbases")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	// catalog, err := renderables.MarshalCatalog(h)
	// fmt.Printf("%v", catalog)
	slug := "b7a5b2a49f3891add85ccea2a65dcf6b"
	data, err := renderables.MarshalRenderable(slug, h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	r := renderables.Renderable{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}

	if len(r.LabelRoles) != 3 {
		t.Fatalf("expected 3 LabelRole; outcome %d;\n", len(r.LabelRoles))
	}

	if len(r.Lang) != 3 {
		t.Fatalf("expected 3 Lang; outcome %d;\n", len(r.Lang))
	}

	if r.RelationshipSet.Title != "100010 - Statement - CONSOLIDATED BALANCE SHEETS" {
		t.Fatalf("expected 100010 - Statement - CONSOLIDATED BALANCE SHEETS; outcome %s;\n", r.RelationshipSet.Title)
	}

	if len(r.PGrid.IndentedLabels) != 56 {
		t.Fatalf("expected 56 Indented Labels; outcome %d;\n", len(r.PGrid.IndentedLabels))
	}

	iLabel42 := r.PGrid.IndentedLabels[42]
	if langPack, found := iLabel42.Label[renderables.Default]; iLabel42.Href != `https://xbrl.fasb.org/us-gaap/2023/elts/us-gaap-2023.xsd#us-gaap_AccruedInsuranceNoncurrent` ||
		!found || len(langPack) != 3 ||
		langPack[renderables.English] != `Accrued Insurance, Noncurrent` {
		t.Fatalf("expected Accrued Insurance, Noncurrent; outcome %v;\n", r.PGrid.IndentedLabels[42])
	}

	sItems := r.CGrid.SummationItems
	if len(sItems) != 10 {
		t.Fatalf("expected 10 Summation Items; outcome %d;\n", len(sItems))
	}
	sItem0 := sItems[0]
	if sItem0.Href != `dpz-20231231.xsd#dpz_AssetsNoncurrentExcludesPropertyPlantAndEquipmentNet` {
		t.Fatalf("expected dpz-20231231.xsd#dpz_AssetsNoncurrentExcludesPropertyPlantAndEquipmentNet; outcome %v;\n", sItem0)
	}
	if len(sItem0.ContributingConcepts) != 7 {
		t.Fatalf("expected 7 Contributing Concepts; outcome %d;\n", len(sItem0.ContributingConcepts))
	}
	if sItem0.ContributingConcepts[5].Href != `https://xbrl.fasb.org/us-gaap/2023/elts/us-gaap-2023.xsd#us-gaap_Investments` ||
		sItem0.ContributingConcepts[5].Sign != `+` || sItem0.ContributingConcepts[0].Scale != `1.0` {
		t.Fatalf("expected postive contribution from https://xbrl.fasb.org/us-gaap/2023/elts/us-gaap-2023.xsd#us-gaap_Investments; outcome %v;\n", sItem0.ContributingConcepts[5])
	}

}

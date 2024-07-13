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

func TestMarshalRenderable_Multiple_Hypercube(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = filepath.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = filepath.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := filepath.Join(serializables.WorkingDirectoryPath, "folders", "multiple_hypercube")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("multiple_hypercube")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	slug := "67117dcea6dfce4b295dfdab4bf10adb"
	data, err := renderables.MarshalRenderable(slug, h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	r := renderables.Renderable{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if r.RelationshipSet.RoleURI != "http://www.workiva.com/role/IncomeTaxesNarrativeDetails" {
		t.Fatalf("expected http://www.workiva.com/role/IncomeTaxesNarrativeDetails; outcome %s;\n", r.RelationshipSet.RoleURI)
	}
	if len(r.DGrid.RootDomains) != 2 {
		t.Fatalf("expected 2 root domains; outcome %d;\n", len(r.DGrid.RootDomains))
	}
	firstRootDomain := r.DGrid.RootDomains[0]
	if firstRootDomain.Href != "https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_OperatingLossCarryforwardsLineItems" {
		t.Fatalf("expected https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_OperatingLossCarryforwardsLineItems; outcome %s;\n", firstRootDomain.Href)
	}
	if len(firstRootDomain.PrimaryItems) != 1 {
		t.Fatalf("expected 1 child primary items; outcome %d;\n", len(firstRootDomain.PrimaryItems))
	}
	if firstRootDomain.PrimaryItems[0].Href != "https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_OperatingLossCarryforwards" {
		t.Fatalf("expected https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_OperatingLossCarryforwards; outcome %s;\n", firstRootDomain.PrimaryItems[0].Href)
	}
	secondRootDomain := r.DGrid.RootDomains[1]
	if secondRootDomain.Href != "https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_TaxCreditCarryforwardLineItems" {
		t.Fatalf("expected https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_TaxCreditCarryforwardLineItems; outcome %s;\n", secondRootDomain.Href)
	}
	if len(secondRootDomain.PrimaryItems) != 1 {
		t.Fatalf("expected 1 child primary items; outcome %d;\n", len(secondRootDomain.PrimaryItems))
	}
	if secondRootDomain.PrimaryItems[0].Href != "https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_TaxCreditCarryforwardAmount" {
		t.Fatalf("expected https://xbrl.fasb.org/us-gaap/2022/elts/us-gaap-2022.xsd#us-gaap_TaxCreditCarryforwardAmount; outcome %s;\n", secondRootDomain.PrimaryItems[0].Href)
	}
}

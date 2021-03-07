package renderables_test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func names() map[string]map[string]string {
	ret := make(map[string]map[string]string)
	cik := "http://www.sec.gov/CIK"
	ret[cik] = make(map[string]string)
	ret[cik]["0001445305"] = "WORKIVA INC"
	ret[cik]["0000069891"] = "FILER DIRECT CORP"
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

	if len(r.LabelRoles) != 3 {
		t.Fatalf("expected 3 LabelRole; outcome %d;\n", len(r.LabelRoles))
	}

	if len(r.Lang) != 2 {
		t.Fatalf("expected 2 Lang; outcome %d;\n", len(r.Lang))
	}

	if r.RelationshipSet.Title != "1001002 - Statement - CONDENSED CONSOLIDATED BALANCE SHEETS" {
		t.Fatalf("expected 1001002 - Statement - CONDENSED CONSOLIDATED BALANCE SHEETS; outcome %s;\n", r.RelationshipSet.Title)
	}

	if r.Subject.Name != "WORKIVA INC" {
		t.Fatalf("expected WORKIVA INC; outcome %s;\n", r.Subject.Name)
	}

	if len(r.PGrid.IndentedLabels) != 43 {
		t.Fatalf("expected 43 Indented Labels; outcome %d;\n", len(r.PGrid.IndentedLabels))
	}

	if r.PGrid.MaxIndentation != 6 {
		t.Fatalf("expected 6 maximum indentation; outcome %d;\n", r.PGrid.MaxIndentation)
	}

	if r.PGrid.MaxDepth != 1 {
		t.Fatalf("expected 1 maximum depth; outcome %d;\n", r.PGrid.MaxDepth)
	}

	iLabel42 := r.PGrid.IndentedLabels[42]
	if langPack, found := iLabel42.Label[renderables.Default]; iLabel42.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_LiabilitiesAndStockholdersEquity` ||
		!found || len(langPack) != 2 ||
		langPack[renderables.English] != `Liabilities and Equity` {
		t.Fatalf("expected Liabilities and Equity; outcome %v;\n", r.PGrid.IndentedLabels[42])
	}

	if len(r.DGrid.RootDomains) != 1 {
		t.Fatalf("expected 1 Root Domain; outcome %d;\n", len(r.DGrid.RootDomains))
	}

	rd := r.DGrid.RootDomains[0]
	if rd.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementLineItems` {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementLineItems; outcome %s;\n", rd.Href)
	}

	if rd.MaxDepth != 1 {
		t.Fatalf("expected 1 maximum depth; outcome %d;\n", rd.MaxDepth)
	}

	if rd.MaxLevel != 3 {
		t.Fatalf("expected 3 maximum level; outcome %d;\n", rd.MaxLevel)
	}

	if len(rd.PrimaryItems) != 36 {
		t.Fatalf("expected 36 non-root primary items; outcome %d;\n", len(rd.PrimaryItems))
	}

	if len(rd.EffectiveDomainGrid) != 37 {
		t.Fatalf("expected 37 effective domain grid item; outcome %d;\n", len(rd.EffectiveDomainGrid))
	}

	edom := rd.EffectiveDomainGrid[0][0]
	if len(edom) != 4 {
		t.Fatalf("expected 4 effective domain members; outcome %d;\n", len(edom))
	}
	edom0 := edom[0]
	if edom0.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ClassOfStockDomain` ||
		!edom0.IsDefault {
		t.Fatalf("expected default member, http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ClassOfStockDomain; outcome %v;\n", edom0)
	}
	edom1 := edom[1]
	if edom1.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ClassOfStockDomain` ||
		edom1.IsDefault {
		t.Fatalf("expected nondefault member, http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ClassOfStockDomain; outcome %v;\n", edom1)
	}
	edom2 := edom[2]
	if edom2.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CommonClassAMember` ||
		edom2.IsDefault {
		t.Fatalf("expected nondefault member, http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CommonClassAMember; outcome %v;\n", edom2)
	}
	edom3 := edom[3]
	if edom3.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CommonClassBMember` ||
		edom3.IsDefault {
		t.Fatalf("expected nondefault member, http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CommonClassBMember; outcome %v;\n", edom3)
	}

	if len(rd.EffectiveDimensions) != 1 {
		t.Fatalf("expected 1 effective dimension; outcome %d;\n", len(rd.EffectiveDimensions))
	}

	ed := rd.EffectiveDimensions[0]
	if ed.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementClassOfStockAxis` {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementClassOfStockAxis; outcome %s;\n", ed.Href)
	}

	sItems := r.CGrid.SummationItems
	if len(sItems) != 6 {
		t.Fatalf("expected 6 Summation Items; outcome %d;\n", len(sItems))
	}
	sItem0 := sItems[0]
	if sItem0.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_Assets` {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_Assets; outcome %v;\n", sItem0)
	}
	if sItem0.MaxDepth != 0 {
		t.Fatalf("expected 0 max depth; outcome %d;\n", sItem0.MaxDepth)
	}
	if len(sItem0.ContributingConcepts) != 6 {
		t.Fatalf("expected 6 Contributing Concepts; outcome %d;\n", len(sItem0.ContributingConcepts))
	}
	if sItem0.ContributingConcepts[5].Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OperatingLeaseRightOfUseAsset` ||
		sItem0.ContributingConcepts[5].Sign != `+` || sItem0.ContributingConcepts[0].Scale != `1.0` {
		t.Fatalf("expected postive contribution from http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OperatingLeaseRightOfUseAsset; outcome %v;\n", sItem0.ContributingConcepts[5])
	}

}

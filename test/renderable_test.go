package telefacts_test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestMarshalRenderable_Ix_Narrative(t *testing.T) {
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
	slug := "bde8e195e86119e0ef56096707591d82"
	data, err := renderables.MarshalRenderable(slug, h)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	r := renderables.Renderable{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}

func TestMarshalRenderable_Gold_BalanceSheet(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = path.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(serializables.WorkingDirectoryPath, "folders", "test_gold")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("test_gold")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	slug := "18cacc74495a0098202b251879753ab2"
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

	if r.RelationshipSet.Title != "1001002 - Statement - CONDENSED CONSOLIDATED BALANCE SHEETS" {
		t.Fatalf("expected 1001002 - Statement - CONDENSED CONSOLIDATED BALANCE SHEETS; outcome %s;\n", r.RelationshipSet.Title)
	}

	if r.Subject.Name != "WORKIVA INC" {
		t.Fatalf("expected WORKIVA INC; outcome %s;\n", r.Subject.Name)
	}

	if len(r.PGrid.IndentedLabels) != 43 {
		t.Fatalf("expected 43 Indented Labels; outcome %d;\n", len(r.PGrid.IndentedLabels))
	}

	iLabel42 := r.PGrid.IndentedLabels[42]
	if langPack, found := iLabel42.Label[renderables.Default]; iLabel42.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CommonClassBMember` ||
		!found || len(langPack) != 3 ||
		langPack[renderables.English] != `Common Class B [Member]` {
		t.Fatalf("expected Common Class B [Member]; outcome %v;\n", r.PGrid.IndentedLabels[42])
	}

	if len(r.DGrid.RootDomains) != 1 {
		t.Fatalf("expected 1 Root Domain; outcome %d;\n", len(r.DGrid.RootDomains))
	}

	rd := r.DGrid.RootDomains[0]
	if rd.Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementLineItems` {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementLineItems; outcome %s;\n", rd.Href)
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
	if len(sItem0.ContributingConcepts) != 6 {
		t.Fatalf("expected 6 Contributing Concepts; outcome %d;\n", len(sItem0.ContributingConcepts))
	}
	if sItem0.ContributingConcepts[5].Href != `http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OperatingLeaseRightOfUseAsset` ||
		sItem0.ContributingConcepts[5].Sign != `+` || sItem0.ContributingConcepts[0].Scale != `1.0` {
		t.Fatalf("expected postive contribution from http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OperatingLeaseRightOfUseAsset; outcome %v;\n", sItem0.ContributingConcepts[5])
	}

}

func TestMarshalRenderable_Gold_TypedMember(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = path.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(serializables.WorkingDirectoryPath, "folders", "test_gold")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		t.Fatalf("Error: " + err.Error())
		return
	}
	f, err := serializables.Discover("test_gold")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	slug := "a57d920a142d3cef3867d22167db52f6"
	data, err := renderables.MarshalRenderable(slug, h)
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

	if len(r.Lang) != 3 {
		t.Fatalf("expected 3 Lang; outcome %d;\n", len(r.Lang))
	}

	if r.RelationshipSet.Title != "2429414 - Disclosure - Revenue Recognition - Deferred Revenue and Transaction Price Allocated to the Remaining Performance Obligations (Details)" {
		t.Fatalf("expected 2429414 - Disclosure - Revenue Recognition - Deferred Revenue and Transaction Price Allocated to the Remaining Performance Obligations (Details); outcome %s;\n", r.RelationshipSet.Title)
	}

	if r.Subject.Name != "WORKIVA INC" {
		t.Fatalf("expected WORKIVA INC; outcome %s;\n", r.Subject.Name)
	}

	if len(r.PGrid.PeriodHeaders) != 6 {
		t.Fatalf("expected 6 PeriodHeaders; outcome %d;\n", len(r.PGrid.PeriodHeaders))
	}

	if len(r.PGrid.PeriodHeaders) != 6 {
		t.Fatalf("expected 6 PeriodHeaders; outcome %d;\n", len(r.PGrid.PeriodHeaders))
	}

	if r.PGrid.PeriodHeaders[5][renderables.PureLabel] != "2020-09-30" {
		t.Fatalf("expected 2020-09-30; outcome %s;\n", r.PGrid.PeriodHeaders[5][renderables.PureLabel])
	}

	if len(r.PGrid.VoidQuadrant) != 2 {
		t.Fatalf("expected 2 VoidQuadrant; outcome %d;\n", len(r.PGrid.VoidQuadrant))
	}

	voidQ0 := r.PGrid.VoidQuadrant[0]
	if voidQ0 == nil || voidQ0.IsParenthesized || voidQ0.Indentation != 1 ||
		voidQ0.TypedDomain != nil ||
		voidQ0.Dimension.Href != "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis" {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis; outcome %s;\n", voidQ0.Dimension.Href)
	}

	voidQ1 := r.PGrid.VoidQuadrant[1]
	if voidQ1 == nil || voidQ1.IsParenthesized || voidQ1.Indentation != 2 ||
		voidQ1.TypedDomain == nil ||
		voidQ1.TypedDomain.Href != "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis.domain" &&
			voidQ1.Dimension.Href != "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis" {
		t.Fatalf("expected http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis.domain; outcome %v;\n", voidQ1.TypedDomain)
	}

	if len(r.PGrid.VoidQuadrant) != 2 {
		t.Fatalf("expected 2 VoidQuadrant; outcome %d;\n", len(r.PGrid.VoidQuadrant))
	}

	if len(r.PGrid.ContextualMemberGrid) != 2 {
		t.Fatalf("expected 2 ContextualMemberGrid rows; outcome %d;\n", len(r.PGrid.ContextualMemberGrid))
	}

	if len(r.PGrid.ContextualMemberGrid[0]) != 6 {
		t.Fatalf("expected 6 ContextualMemberGrid columns; outcome %d;\n", len(r.PGrid.ContextualMemberGrid[0]))
	}

	if r.PGrid.ContextualMemberGrid[1][5].TypedMember != "2020-10-01" &&
		r.PGrid.ContextualMemberGrid[1][5].ExplicitMember == nil {
		t.Fatalf("expected 2020-10-01; outcome %d;\n", len(r.PGrid.ContextualMemberGrid[1][5].TypedMember))
	}

	if len(r.PGrid.FactualQuadrant) != 7 {
		t.Fatalf("expected 7 FactualQuadrant rows; outcome %d;\n", len(r.PGrid.FactualQuadrant))
	}

	if len(r.PGrid.FactualQuadrant[0]) != 6 {
		t.Fatalf("expected 6 FactualQuadrant columns; outcome %d;\n", len(r.PGrid.FactualQuadrant[0]))
	}
}

func hydratableFactory(id string) (*hydratables.Hydratable, error) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(serializables.WorkingDirectoryPath, "folders", id)
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf(err.Error())
	}
	f, err := serializables.Discover(id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return h, nil
}

func bencmarkMarshalRenderable(slug string, h *hydratables.Hydratable, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := renderables.MarshalRenderable(slug, h)
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

func BenchmarkMarshalRenderable_Gold_BalanceSheet(b *testing.B) {
	h, err := hydratableFactory("test_gold")
	if err != nil {
		b.Fatalf(err.Error())
	}
	slug := "883459b49fae34a739704b6db51d6b1d"
	bencmarkMarshalRenderable(slug, h, b)
}

func BenchmarkMarshalRenderable_Gold_EquityTable(b *testing.B) {
	h, err := hydratableFactory("test_gold")
	if err != nil {
		b.Fatalf(err.Error())
	}
	slug := "4d034c1e44b980a9940e857682b81991"
	bencmarkMarshalRenderable(slug, h, b)
}

func ExampleMarshalCatalog_Hello() {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = path.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(serializables.WorkingDirectoryPath, "folders", "hello")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		panic("Error: " + err.Error())
	}
	f, err := serializables.Discover("hello")
	if err != nil {
		panic("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		panic("Error: " + err.Error())
	}
	slug := "f5ed7171b09c4487172b60563de820dd"
	_, err = renderables.MarshalRenderable(slug, h)
	if err != nil {
		panic("Error: " + err.Error())
	}
	fmt.Println(slug)
	// Output: f5ed7171b09c4487172b60563de820dd
}

func ExampleMarshalCatalog_485() {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
	serializables.GlobalTaxonomySetPath = path.Join(".", "gts")
	hydratables.InjectCache(hcache)
	workingDir := path.Join(serializables.WorkingDirectoryPath, "folders", "485")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		panic("Error: " + err.Error())
	}
	f, err := serializables.Discover("485")
	if err != nil {
		panic("Error: " + err.Error())
	}
	h, err := hydratables.Hydrate(f)
	if err != nil {
		panic("Error: " + err.Error())
	}
	slug := "79133e50935a2a5e7c3fcc915137ed9c"
	_, err = renderables.MarshalRenderable(slug, h)
	if err != nil {
		panic("Error: " + err.Error())
	}
	fmt.Println(slug)
	// Output: 79133e50935a2a5e7c3fcc915137ed9c
}

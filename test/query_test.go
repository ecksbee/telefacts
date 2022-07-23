package telefacts_test

import (
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func TestHashQuery_Gold(t *testing.T) {
	hcache := gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	serializables.WorkingDirectoryPath = path.Join(".", "wd")
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

	_, cccc, err := h.NameQuery("http://www.workiva.com/20200930", "IncreaseDecreaseInOperatingLeaseLiability")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if cccc.ID != "wk_IncreaseDecreaseInOperatingLeaseLiability" || cccc.Abstract ||
		cccc.PeriodType != "duration" || cccc.SubstitutionGroup.Local != "item" ||
		cccc.XMLName.Space != "http://www.workiva.com/20200930" ||
		cccc.XMLName.Local != "IncreaseDecreaseInOperatingLeaseLiability" {
		t.Fatalf("expected IncreaseDecreaseInOperatingLeaseLiability; outcome %v;\n", cccc)
	}
	_, ccc, err := h.HashQuery("wk-20200930.xsd#wk_IncreaseDecreaseInOperatingLeaseLiability")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if ccc.ID != "wk_IncreaseDecreaseInOperatingLeaseLiability" || ccc.Abstract ||
		ccc.PeriodType != "duration" || ccc.SubstitutionGroup.Local != "item" ||
		ccc.XMLName.Space != "http://www.workiva.com/20200930" ||
		ccc.XMLName.Local != "IncreaseDecreaseInOperatingLeaseLiability" {
		t.Fatalf("expected IncreaseDecreaseInOperatingLeaseLiability; outcome %v;\n", ccc)
	}

	_, c, err := h.HashQuery("http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementOfStockholdersEquityAbstract")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if c.ID != "us-gaap_StatementOfStockholdersEquityAbstract" || !c.Abstract ||
		c.PeriodType != "duration" || c.SubstitutionGroup.Local != "item" ||
		c.XMLName.Space != "http://fasb.org/us-gaap/2020-01-31" ||
		c.XMLName.Local != "StatementOfStockholdersEquityAbstract" {
		t.Fatalf("expected StatementOfStockholdersEquityAbstract; outcome %v;\n", c)
	}
	_, cc, err := h.NameQuery("http://fasb.org/us-gaap/2020-01-31", "StatementOfStockholdersEquityAbstract")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if cc.ID != "us-gaap_StatementOfStockholdersEquityAbstract" || !cc.Abstract ||
		cc.PeriodType != "duration" || cc.SubstitutionGroup.Local != "item" ||
		cc.XMLName.Space != "http://fasb.org/us-gaap/2020-01-31" ||
		cc.XMLName.Local != "StatementOfStockholdersEquityAbstract" {
		t.Fatalf("expected StatementOfStockholdersEquityAbstract; outcome %v;\n", cc)
	}
	typedDomainHref, typedDomain, err := h.NameQuery("http://fasb.org/us-gaap/2020-01-31", "RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis.domain")
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	if typedDomainHref == "" || typedDomain == nil {
		t.Fatalf("exptected RevenueRemainingPerformanceObligationExpectedTimingOfSatisfactionStartDateAxis.domain; outcome: nil or blank typed domain")
	}
}

package xbrl_test

import (
	"testing"

	"ecks-bee.com/telefacts/xbrl"
)

const testInstance = `
<?xml version="1.0" encoding="US-ASCII" ?>
    <!-- TEST COMMENTS -->
<xbrli:xbrl xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:link="http://www.xbrl.org/2003/linkbase" xmlns:xbrli="http://www.xbrl.org/2003/instance" xmlns:xbrldt="http://xbrl.org/2005/xbrldt" xmlns:xbrldi="http://xbrl.org/2006/xbrldi" xmlns:dei="http://xbrl.sec.gov/dei/2019-01-31" xmlns:ref="http://www.xbrl.org/2006/ref" xmlns:iso4217="http://www.xbrl.org/2003/iso4217" xmlns:us-gaap="http://fasb.org/us-gaap/2020-01-31" xmlns:us-roles="http://fasb.org/us-roles/2020-01-31" xmlns:nonnum="http://www.xbrl.org/dtr/type/non-numeric" xmlns:num="http://www.xbrl.org/dtr/type/numeric" xmlns:us-types="http://fasb.org/us-types/2020-01-31" xmlns:fr="http://www.xbrl-fr.org/xbrl/2003-02-29" xmlns:country="http://xbrl.sec.gov/country/2020-01-31" xmlns:srt="http://fasb.org/srt/2020-01-31" xmlns:EG="http://example.com/20200630">
    <link:schemaRef xlink:href="isdr-20200630.xsd" xlink:type="simple" />
	<xbrli:context id="c1">
		<xbrli:entity>
			<xbrli:identifier scheme="http://www.un.org/">Escargot</xbrli:identifier>
		</xbrli:entity>
		<xbrli:period>
			<xbrli:instant>2001-08-16</xbrli:instant>
		</xbrli:period>
		<xbrli:scenario>
			<fr:scenarioType>actual</fr:scenarioType>
		</xbrli:scenario>
	</xbrli:context>
    <xbrli:context id="From2020-04-01to2020-06-30_us-gaap_CostOfSalesMember">
      <xbrli:entity>
        <xbrli:identifier scheme="http://www.sec.gov/CIK">1234567890</xbrli:identifier>
        <xbrli:segment>
          <xbrldi:explicitMember dimension="us-gaap:IncomeStatementLocationAxis">us-gaap:CostOfSalesMember</xbrldi:explicitMember>
        </xbrli:segment>
      </xbrli:entity>
      <xbrli:period>
        <xbrli:startDate>2020-04-01</xbrli:startDate>
        <xbrli:endDate>2020-06-30</xbrli:endDate>
      </xbrli:period>
    </xbrli:context>
    <xbrli:context id="From2020-01-01to2020-06-30">
      <xbrli:entity>
        <xbrli:identifier scheme="http://www.sec.gov/CIK">1234567890</xbrli:identifier>
      </xbrli:entity>
      <xbrli:period>
        <xbrli:startDate>2020-01-01</xbrli:startDate>
        <xbrli:endDate>2020-06-30</xbrli:endDate>
      </xbrli:period>
    </xbrli:context>
    <xbrli:context id="AsOf2020-06-30_TypedDimensions">
      <xbrli:entity>
		<xbrli:identifier scheme="http://www.sec.gov/CIK">1234567890</xbrli:identifier>
		<segment>
			<xbrldi:typedMember dimension="EG:TeamDim">
				<EG:Team>Lakers</EG:Team>
			</xbrldi:typedMember>
			<xbrldi:typedMember dimension="EG:DrinkDim">
				<EG:Drink>Coors</EG:Drink>
			</xbrldi:typedMember>
		</segment>
      </xbrli:entity>
      <xbrli:period>
        <xbrli:instant>2020-06-30</xbrli:instant>
      </xbrli:period>
    </xbrli:context>
    <xbrli:context id="AsOf2019-12-31">
      <xbrli:entity>
        <xbrli:identifier scheme="http://www.sec.gov/CIK">1234567890</xbrli:identifier>
      </xbrli:entity>
      <xbrli:period>
        <xbrli:instant>2019-12-31</xbrli:instant>
      </xbrli:period>
    </xbrli:context>
    <xbrli:context id="From2019-01-01to2019-06-30">
      <xbrli:entity>
        <xbrli:identifier scheme="http://www.sec.gov/CIK">1234567890</xbrli:identifier>
      </xbrli:entity>
      <xbrli:period>
        <xbrli:startDate>2019-01-01</xbrli:startDate>
        <xbrli:endDate>2019-06-30</xbrli:endDate>
      </xbrli:period>
    </xbrli:context>
    <xbrli:unit id="USD">
      <xbrli:measure>iso4217:USD</xbrli:measure>
    </xbrli:unit>
    <xbrli:unit id="Shares">
      <xbrli:measure>xbrli:shares</xbrli:measure>
    </xbrli:unit>
    <xbrli:unit id="USDPShares">
      <xbrli:divide>
        <xbrli:unitNumerator>
          <xbrli:measure>iso4217:USD</xbrli:measure>
        </xbrli:unitNumerator>
        <xbrli:unitDenominator>
          <xbrli:measure>xbrli:shares</xbrli:measure>
        </xbrli:unitDenominator>
      </xbrli:divide>
    </xbrli:unit>
    <xbrli:unit id="Percent">
      <xbrli:measure>xbrli:pure</xbrli:measure>
    </xbrli:unit>
	<xbrli:unit id="EUR">
		<xbrli:measure>ISO4217:EUR</xbrli:measure>
	</xbrli:unit>
    <us-gaap:LineOfCredit contextRef="AsOf2020-06-30" unitRef="USD" decimals="-3">0</us-gaap:LineOfCredit>
	<fr:assetsTotal id="f1" precision="4" unitRef="EUR" contextRef="c1">2600</fr:assetsTotal>
	<fr:liabilitiesTotal id="f2" precision="4" unitRef="EUR" contextRef="c1">2600</fr:liabilitiesTotal>
	<fr:equityTotal id="f3" precision="4" unitRef="EUR" contextRef="c1">1100</fr:equityTotal>
	<link:footnoteLink xlink:type="extended" xlink:title="Document 1 Footnotes" xlink:role="http://www.xbrl.org/2003/role/link">
		<link:footnote xlink:type="resource" xlink:label="footnote1" xlink:role="http://www.xbrl.org/2003/role/footnote" xml:lang="en">Including the effects of the merger.</link:footnote>
		<link:footnote xlink:type="resource" xlink:label="footnote1" xlink:role="http://www.xbrl.org/2003/role/footnote" xml:lang="fr">Y compris les effets de la fusion.</link:footnote>
		<link:loc xlink:type="locator" xlink:label="fact1" xlink:href="#f1"/>
		<link:loc xlink:type="locator" xlink:label="fact1" xlink:href="#f2"/>
		<link:loc xlink:type="locator" xlink:label="fact1" xlink:href="#f3"/>
		<link:footnoteArc xlink:type="arc" xlink:from="fact1" xlink:to="footnote1" xlink:title="view explanatory footnote" xlink:arcrole="http://www.xbrl.org/2003/arcrole/fact-footnote"/>
	</link:footnoteLink>
</xbrli:xbrl>
`

func TestDecodeInstance(t *testing.T) {
	stimulus := []byte(testInstance)
	decoded, err := xbrl.DecodeInstance(stimulus)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	contexts := decoded.Context
	if len(contexts) != 6 {
		t.Fatalf("expected 6 contexts; outcome %d;\n%v\n", len(contexts), contexts)
	}
	units := decoded.Unit
	if len(units) != 5 {
		t.Fatalf("expected 5 units; outcome %d;\n%v\n", len(units), units)
	}
	schemaRefs := decoded.SchemaRef
	if len(schemaRefs) != 1 {
		t.Fatalf("expected 1 schemaRef; outcome %d;\n%v\n", len(schemaRefs), schemaRefs)
	}
	facts := decoded.Facts
	if len(facts) != 4 {
		t.Fatalf("expected 4 fact; outcome %d;\n%v\n", len(facts), facts)
	}
	footnoteLinks := decoded.FootnoteLink
	if len(footnoteLinks) != 1 {
		t.Fatalf("expected 1 footnoteLink; outcome %d;\n%v\n", len(footnoteLinks), footnoteLinks)
	}
}

package serializables_test

import (
	"testing"

	"ecksbee.com/telefacts/serializables"
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

func TestDecodeInstanceFile(t *testing.T) {
	stimulus := []byte(testInstance)
	decoded, err := serializables.DecodeInstanceFile(stimulus)
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

const testInstancePrefixless = `<xbrl xmlns="http://www.xbrl.org/2003/instance" xmlns:dei="http://xbrl.sec.gov/dei/2020-01-31" xmlns:iso4217="http://www.xbrl.org/2003/iso4217" xmlns:link="http://www.xbrl.org/2003/linkbase" xmlns:srt="http://fasb.org/srt/2020-01-31" xmlns:us-gaap="http://fasb.org/us-gaap/2020-01-31" xmlns:wk="http://www.workiva.com/20200930" xmlns:xbrldi="http://xbrl.org/2006/xbrldi" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xml:lang="en-US">
<link:schemaRef xlink:href="wk-20200930.xsd" xlink:type="simple"/>

<context id="i379ca915ae9345069ac9acf544faa7f5_D20200101-20200930">
<entity>
<identifier scheme="http://www.sec.gov/CIK">0001445305</identifier>
</entity>
<period>
<startDate>2020-01-01</startDate>
<endDate>2020-09-30</endDate>
</period>
</context>
<context id="ib63b333842984c9992f7bffd44303f82_I20201031">
<entity>
<identifier scheme="http://www.sec.gov/CIK">0001445305</identifier>
<segment>
<xbrldi:explicitMember dimension="us-gaap:StatementClassOfStockAxis">us-gaap:CommonClassAMember</xbrldi:explicitMember>
</segment>
</entity>
<period>
<instant>2020-10-31</instant>
</period>
</context>
<context id="ieadfe3bba8134188a48e44a78f0494b0_I20191231">
<entity>
<identifier scheme="http://www.sec.gov/CIK">0001445305</identifier>
<segment>
<xbrldi:explicitMember dimension="us-gaap:AwardTypeAxis">us-gaap:RestrictedStockUnitsRSUMember</xbrldi:explicitMember>
</segment>
</entity>
<period>
<instant>2019-12-31</instant>
</period>
</context>
<context id="i59b532e5f05d45f48fca68755586d87d_D20190101-20190930">
<entity>
<identifier scheme="http://www.sec.gov/CIK">0001445305</identifier>
<segment>
<xbrldi:explicitMember dimension="us-gaap:AwardTypeAxis">us-gaap:RestrictedStockUnitsRSUMember</xbrldi:explicitMember>
</segment>
</entity>
<period>
<startDate>2019-01-01</startDate>
<endDate>2019-09-30</endDate>
</period>
</context>

<unit id="shares">
<measure>shares</measure>
</unit>
<unit id="usd">
<measure>iso4217:USD</measure>
</unit>
<unit id="usdPerShare">
<divide>
<unitNumerator>
<measure>iso4217:USD</measure>
</unitNumerator>
<unitDenominator>
<measure>shares</measure>
</unitDenominator>
</divide>
</unit>
<unit id="number">
<measure>pure</measure>
</unit>

<wk:IncreaseDecreaseInOperatingLeaseLiability contextRef="i13d9cdcd7ad54a0ea865c4883a7aa166_D20190701-20190930" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8zMS9mcmFnOmY1YWMzNTU1NGQ5NDQyMTg4ZDQ2N2E5NjA0YThiOTViL3RhYmxlOmE3ODhhZDk2YTk3ODQyMDE5MjNhNTliOWViYmU5ZThmL3RhYmxlcmFuZ2U6YTc4OGFkOTZhOTc4NDIwMTkyM2E1OWI5ZWJiZTllOGZfMjItMy0xLTEtMA_6c864169-4fa2-438a-bd16-751839809bb1" unitRef="usd">-758000</wk:IncreaseDecreaseInOperatingLeaseLiability>
<wk:IncreaseDecreaseInOperatingLeaseLiability contextRef="i379ca915ae9345069ac9acf544faa7f5_D20200101-20200930" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8zMS9mcmFnOmY1YWMzNTU1NGQ5NDQyMTg4ZDQ2N2E5NjA0YThiOTViL3RhYmxlOmE3ODhhZDk2YTk3ODQyMDE5MjNhNTliOWViYmU5ZThmL3RhYmxlcmFuZ2U6YTc4OGFkOTZhOTc4NDIwMTkyM2E1OWI5ZWJiZTllOGZfMjItNS0xLTEtMA_697b0b4a-f994-46cf-9adb-16de53aa3466" unitRef="usd">-3438000</wk:IncreaseDecreaseInOperatingLeaseLiability>
<wk:IncreaseDecreaseInOperatingLeaseLiability contextRef="id7ba101c63b94137966e8d5cdb18d84e_D20190101-20190930" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8zMS9mcmFnOmY1YWMzNTU1NGQ5NDQyMTg4ZDQ2N2E5NjA0YThiOTViL3RhYmxlOmE3ODhhZDk2YTk3ODQyMDE5MjNhNTliOWViYmU5ZThmL3RhYmxlcmFuZ2U6YTc4OGFkOTZhOTc4NDIwMTkyM2E1OWI5ZWJiZTllOGZfMjItNy0xLTEtMA_0dda71b3-200e-4f10-a050-253d64f6f30f" unitRef="usd">-2226000</wk:IncreaseDecreaseInOperatingLeaseLiability>

<dei:EntityShellCompany contextRef="i379ca915ae9345069ac9acf544faa7f5_D20200101-20200930" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xL2ZyYWc6M2E1OGVkNmQ4NDY3NGI0Yzk1MGQzMGQwOGZlMzBlNjIvdGV4dHJlZ2lvbjozYTU4ZWQ2ZDg0Njc0YjRjOTUwZDMwZDA4ZmUzMGU2Ml8xODY4_6602876e-a10d-4190-9688-8f6be9d41455">false</dei:EntityShellCompany>
<dei:EntityCommonStockSharesOutstanding contextRef="ib63b333842984c9992f7bffd44303f82_I20201031" decimals="INF" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xL2ZyYWc6M2E1OGVkNmQ4NDY3NGI0Yzk1MGQzMGQwOGZlMzBlNjIvdGV4dHJlZ2lvbjozYTU4ZWQ2ZDg0Njc0YjRjOTUwZDMwZDA4ZmUzMGU2Ml8xNzQ3_61dce85e-1042-4e40-b3af-9cf56cb78966" unitRef="shares">39896661</dei:EntityCommonStockSharesOutstanding>
<dei:EntityCommonStockSharesOutstanding contextRef="i3184d5cccd43432da232da57d7002ef2_I20201031" decimals="INF" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xL2ZyYWc6M2E1OGVkNmQ4NDY3NGI0Yzk1MGQzMGQwOGZlMzBlNjIvdGV4dHJlZ2lvbjozYTU4ZWQ2ZDg0Njc0YjRjOTUwZDMwZDA4ZmUzMGU2Ml8xODAy_cd1def32-33a9-484e-bed0-f79b73f31759" unitRef="shares">8299610</dei:EntityCommonStockSharesOutstanding>
<us-gaap:CashAndCashEquivalentsAtCarryingValue contextRef="ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xOS9mcmFnOjU5Y2QwNmVlYzgwNTQ2NWE5ZjMyMzQ4ODg5MTY4ZmY2L3RhYmxlOjFiMzM0ZmU1MDJiNTQxYTQ4NTk3N2RiNjJhNGEyZmQ2L3RhYmxlcmFuZ2U6MWIzMzRmZTUwMmI1NDFhNDg1OTc3ZGI2MmE0YTJmZDZfNi0xLTEtMS0w_be035ad3-0c89-4702-b376-c8956e84a18f" unitRef="usd">426132000</us-gaap:CashAndCashEquivalentsAtCarryingValue>
<us-gaap:CashAndCashEquivalentsAtCarryingValue contextRef="ifc7bef0caea14d3ba6592c2ac3e05544_I20191231" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xOS9mcmFnOjU5Y2QwNmVlYzgwNTQ2NWE5ZjMyMzQ4ODg5MTY4ZmY2L3RhYmxlOjFiMzM0ZmU1MDJiNTQxYTQ4NTk3N2RiNjJhNGEyZmQ2L3RhYmxlcmFuZ2U6MWIzMzRmZTUwMmI1NDFhNDg1OTc3ZGI2MmE0YTJmZDZfNi0zLTEtMS0w_0e595110-c6ee-4a53-afb2-5aa2c1972ab2" unitRef="usd">381742000</us-gaap:CashAndCashEquivalentsAtCarryingValue>
<us-gaap:AvailableForSaleSecuritiesDebtSecuritiesCurrent contextRef="ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xOS9mcmFnOjU5Y2QwNmVlYzgwNTQ2NWE5ZjMyMzQ4ODg5MTY4ZmY2L3RhYmxlOjFiMzM0ZmU1MDJiNTQxYTQ4NTk3N2RiNjJhNGEyZmQ2L3RhYmxlcmFuZ2U6MWIzMzRmZTUwMmI1NDFhNDg1OTc3ZGI2MmE0YTJmZDZfNy0xLTEtMS0w_514f2bf3-385c-472f-97fe-8b834ab2a509" unitRef="usd">97750000</us-gaap:AvailableForSaleSecuritiesDebtSecuritiesCurrent>
<us-gaap:AvailableForSaleSecuritiesDebtSecuritiesCurrent contextRef="ifc7bef0caea14d3ba6592c2ac3e05544_I20191231" decimals="-3" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV8xOS9mcmFnOjU5Y2QwNmVlYzgwNTQ2NWE5ZjMyMzQ4ODg5MTY4ZmY2L3RhYmxlOjFiMzM0ZmU1MDJiNTQxYTQ4NTk3N2RiNjJhNGEyZmQ2L3RhYmxlcmFuZ2U6MWIzMzRmZTUwMmI1NDFhNDg1OTc3ZGI2MmE0YTJmZDZfNy0zLTEtMS0w_271d2160-e2de-44ce-9e96-16078354ab2e" unitRef="usd">106214000</us-gaap:AvailableForSaleSecuritiesDebtSecuritiesCurrent>
<us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount contextRef="i35820fd697f4451691f86a6641327ef6_D20190101-20190930" decimals="0" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RhYmxlOjRjMTUwOWI2ZTBmZTQ2NGFiMGJiZWU4ODI3Yjk3ZDk5L3RhYmxlcmFuZ2U6NGMxNTA5YjZlMGZlNDY0YWIwYmJlZTg4MjdiOTdkOTlfMi0zLTEtMS0w_0ec9f592-e626-4dcd-8c3e-0464f39a9f1d" unitRef="shares">4542729</us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount>
<us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount contextRef="ie3dab6944c334d2ab2e5d3635e761d5e_D20200101-20200930" decimals="0" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RhYmxlOjRjMTUwOWI2ZTBmZTQ2NGFiMGJiZWU4ODI3Yjk3ZDk5L3RhYmxlcmFuZ2U6NGMxNTA5YjZlMGZlNDY0YWIwYmJlZTg4MjdiOTdkOTlfMy0xLTEtMS0w_fc5ea5d9-389f-4b9f-92f9-8f0a86e8aa5f" unitRef="shares">2974719</us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount>
<us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount contextRef="i999f4c7c21eb4606b7a9a591eccbaebf_D20190101-20190930" decimals="0" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RhYmxlOjRjMTUwOWI2ZTBmZTQ2NGFiMGJiZWU4ODI3Yjk3ZDk5L3RhYmxlcmFuZ2U6NGMxNTA5YjZlMGZlNDY0YWIwYmJlZTg4MjdiOTdkOTlfMy0zLTEtMS0w_f5f9f321-03d2-48d1-8305-69f1e6d709de" unitRef="shares">2844156</us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount>
<us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount contextRef="ic99e6e28e6a640f481551dc0d5e8c038_D20200101-20200930" decimals="0" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RhYmxlOjRjMTUwOWI2ZTBmZTQ2NGFiMGJiZWU4ODI3Yjk3ZDk5L3RhYmxlcmFuZ2U6NGMxNTA5YjZlMGZlNDY0YWIwYmJlZTg4MjdiOTdkOTlfNC0xLTEtMS0w_0dbbfca0-f9e0-47fe-b96c-e48d2b728163" unitRef="shares">95084</us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount>
<us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount contextRef="i51ad35890d054a49916bd8d50844064b_D20190101-20190930" decimals="0" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RhYmxlOjRjMTUwOWI2ZTBmZTQ2NGFiMGJiZWU4ODI3Yjk3ZDk5L3RhYmxlcmFuZ2U6NGMxNTA5YjZlMGZlNDY0YWIwYmJlZTg4MjdiOTdkOTlfNC0zLTEtMS0w_ec9d1d1e-6b76-4473-b4c6-6987ea67d4e3" unitRef="shares">77788</us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount>
<us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount contextRef="i329e1bdcd3474f879177d386b4fcc28b_D20200101-20200930" decimals="-5" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RleHRyZWdpb246ZmFiNDZkOTk2NmFmNDg4MDkxNDlkNDgyZTAyMDc5YTRfMTIzMw_9a4e1eb9-a5da-422c-807b-c02d97d3d328" unitRef="shares">4300000</us-gaap:AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount>
<us-gaap:DebtInstrumentConvertibleConversionPrice1 contextRef="i17db5d7860a840fbbb273f6501008f0a_I20190831" decimals="2" id="id3VybDovL2RvY3MudjEvZG9jOmRiY2FjMzAyYmJiODQxZjc5YmM3Zjk4NGEwYWExNTMxL3NlYzpkYmNhYzMwMmJiYjg0MWY3OWJjN2Y5ODRhMGFhMTUzMV82NC9mcmFnOmZhYjQ2ZDk5NjZhZjQ4ODA5MTQ5ZDQ4MmUwMjA3OWE0L3RleHRyZWdpb246ZmFiNDZkOTk2NmFmNDg4MDkxNDlkNDgyZTAyMDc5YTRfMTcxMw_b9358364-736c-4082-aaea-6aa6b5da29a2" unitRef="usdPerShare">80.16</us-gaap:DebtInstrumentConvertibleConversionPrice1>
</xbrl>`

func TestDecodeInstanceFilePrefixless(t *testing.T) {
	stimulus := []byte(testInstancePrefixless)
	decoded, err := serializables.DecodeInstanceFile(stimulus)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	contexts := decoded.Context
	if len(contexts) != 4 {
		t.Fatalf("expected 4 contexts; outcome %d;\n%v\n", len(contexts), contexts)
	}
	units := decoded.Unit
	if len(units) != 4 {
		t.Fatalf("expected 4 units; outcome %d;\n%v\n", len(units), units)
	}
	schemaRefs := decoded.SchemaRef
	if len(schemaRefs) != 1 {
		t.Fatalf("expected 1 schemaRef; outcome %d;\n%v\n", len(schemaRefs), schemaRefs)
	}
	facts := decoded.Facts
	if len(facts) != 17 {
		t.Fatalf("expected 17 fact; outcome %d;\n%v\n", len(facts), facts)
	}
}

package telefacts_test

import (
	"testing"

	"ecksbee.com/telefacts/pkg/serializables"
)

const testSchema = `
<?xml version="1.0" encoding="US-ASCII" ?>
    <!-- Field: Doc-Info; Name: Generator; Value: QXInteractive; Version: 4.26a -->
    <!-- Field: Doc-Info; Name: Source; Value: 197797 12312020 10K.xfr; Date: 2020%2D02%2D27T14:07:19Z -->
    <!-- Field: Doc-Info; Name: Status; Value: 0x80100019 -->
    <!-- Field: Doc-Info; Name: Misc; Value: 9aQ5w7xRiXgen8uOanNMSaiBOcPtwZB6aYlBvmlzfdkqDygvYrtXcSmTimTWMZkK -->
<schema xmlns="http://www.w3.org/2001/XMLSchema" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:link="http://www.xbrl.org/2003/linkbase" xmlns:xbrli="http://www.xbrl.org/2003/instance" xmlns:xbrldt="http://xbrl.org/2005/xbrldt" xmlns:xbrldi="http://xbrl.org/2006/xbrldi" xmlns:dei="http://xbrl.sec.gov/dei/2018-01-31" xmlns:us-gaap="http://fasb.org/us-gaap/2019-01-31" xmlns:us-roles="http://fasb.org/us-roles/2019-01-31" xmlns:nonnum="http://www.xbrl.org/dtr/type/non-numeric" xmlns:num="http://www.xbrl.org/dtr/type/numeric" xmlns:us-types="http://fasb.org/us-types/2019-01-31" xmlns:ISDR="http://issuerdirect.com/20191231" elementFormDefault="qualified" targetNamespace="http://issuerdirect.com/20191231">
    <annotation>
      <appinfo>
	<link:roleType roleURI="http://issuerdirect.com/role/DocumentAndEntityInformation" id="DocumentAndEntityInformation">
	  <link:definition>00000001 - Document - Document and Entity Information</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/BalanceSheets" id="BalanceSheets">
	  <link:definition>00000002 - Statement - Consolidated Balance Sheets</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/BalanceSheetsParenthetical" id="BalanceSheetsParenthetical">
	  <link:definition>00000003 - Statement - Consolidated Balance Sheets (Parenthetical)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/StatementsOfIncome" id="StatementsOfIncome">
	  <link:definition>00000004 - Statement - Consolidated Statements of Income</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/StatementsOfComprehensiveIncome" id="StatementsOfComprehensiveIncome">
	  <link:definition>00000005 - Statement - Consolidated Statements of Comprehensive Income</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/StatementsOfStockholdersEquity" id="StatementsOfStockholdersEquity">
	  <link:definition>00000006 - Statement - Consolidated Statements of Stockholders' Equity</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/StatementsOfCashFlows" id="StatementsOfCashFlows">
	  <link:definition>00000007 - Statement - Consolidated Statements of Cash Flows</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note1.DescriptionBackgroundAndBasisOfOperations" id="Note1.DescriptionBackgroundAndBasisOfOperations">
	  <link:definition>00000008 - Disclosure - Note 1. Description, Background and Basis of Operations</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPolicies" id="Note2.SummaryOfSignificantAccountingPolicies">
	  <link:definition>00000009 - Disclosure - Note 2. Summary of Significant Accounting Policies</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note3.FixedAssets" id="Note3.FixedAssets">
	  <link:definition>00000010 - Disclosure - Note 3. Fixed Assets</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note4.RecentAcquisitions" id="Note4.RecentAcquisitions">
	  <link:definition>00000011 - Disclosure - Note 4. Recent Acquisitions</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note5.GoodwillAndOtherIntangibleAssets" id="Note5.GoodwillAndOtherIntangibleAssets">
	  <link:definition>00000012 - Disclosure - Note 5. Goodwill and Other Intangible Assets</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note6.LineOfCredit" id="Note6.LineOfCredit">
	  <link:definition>00000013 - Disclosure - Note 6. Line of Credit</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note7.Equity" id="Note7.Equity">
	  <link:definition>00000014 - Disclosure - Note 7. Equity</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnits" id="Note8.StockOptionsAndRestrictedStockUnits">
	  <link:definition>00000015 - Disclosure - Note 8. Stock Options and Restricted Stock Units</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note9.Leases" id="Note9.Leases">
	  <link:definition>00000016 - Disclosure - Note 9. Leases</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note10.CommitmentsAndContingencies" id="Note10.CommitmentsAndContingencies">
	  <link:definition>00000017 - Disclosure - Note 10. Commitments and Contingencies</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note11.Revenues" id="Note11.Revenues">
	  <link:definition>00000018 - Disclosure - Note 11. Revenues</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note12.IncomeTaxes" id="Note12.IncomeTaxes">
	  <link:definition>00000019 - Disclosure - Note 12. Income Taxes</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note13.EmployeeBenefitPlan" id="Note13.EmployeeBenefitPlan">
	  <link:definition>00000020 - Disclosure - Note 13. Employee Benefit Plan</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPoliciesPolicies" id="Note2.SummaryOfSignificantAccountingPoliciesPolicies">
	  <link:definition>00000021 - Disclosure - Note 2. Summary of Significant Accounting Policies (Policies)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPoliciesTables" id="Note2.SummaryOfSignificantAccountingPoliciesTables">
	  <link:definition>00000022 - Disclosure - Note 2. Summary of Significant Accounting Policies (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note3.FixedAssetsTables" id="Note3.FixedAssetsTables">
	  <link:definition>00000023 - Disclosure - Note 3. Fixed Assets (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note4.RecentAcquisitionsTables" id="Note4.RecentAcquisitionsTables">
	  <link:definition>00000024 - Disclosure - Note 4. Recent Acquisitions (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note5.GoodwillAndOtherIntangibleAssetsTables" id="Note5.GoodwillAndOtherIntangibleAssetsTables">
	  <link:definition>00000025 - Disclosure - Note 5. Goodwill and Other Intangible Assets (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnitsTables" id="Note8.StockOptionsAndRestrictedStockUnitsTables">
	  <link:definition>00000026 - Disclosure - Note 8. Stock Options and Restricted Stock Units (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note9.LeasesTables" id="Note9.LeasesTables">
	  <link:definition>00000027 - Disclosure - Note 9. Leases (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note11.RevenuesTables" id="Note11.RevenuesTables">
	  <link:definition>00000028 - Disclosure - Note 11. Revenues (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note12.IncomeTaxesTables" id="Note12.IncomeTaxesTables">
	  <link:definition>00000029 - Disclosure - Note 12. Income Taxes (Tables)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPoliciesDetails" id="Note2.SummaryOfSignificantAccountingPoliciesDetails">
	  <link:definition>00000030 - Disclosure - Note 2. Summary of Significant Accounting Policies (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPoliciesDetails1" id="Note2.SummaryOfSignificantAccountingPoliciesDetails1">
	  <link:definition>00000031 - Disclosure - Note 2. Summary of Significant Accounting Policies (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPoliciesDetails2" id="Note2.SummaryOfSignificantAccountingPoliciesDetails2">
	  <link:definition>00000032 - Disclosure - Note 2. Summary of Significant Accounting Policies (Details 2)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note2.SummaryOfSignificantAccountingPoliciesDetailsNarrative" id="Note2.SummaryOfSignificantAccountingPoliciesDetailsNarrative">
	  <link:definition>00000033 - Disclosure - Note 2. Summary of Significant Accounting Policies (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note3.FixedAssetsDetails" id="Note3.FixedAssetsDetails">
	  <link:definition>00000034 - Disclosure - Note 3. Fixed Assets (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note3.FixedAssetsDetailsNarrative" id="Note3.FixedAssetsDetailsNarrative">
	  <link:definition>00000035 - Disclosure - Note 3. Fixed Assets (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note4.RecentAcquisitionsDetails" id="Note4.RecentAcquisitionsDetails">
	  <link:definition>00000036 - Disclosure - Note 4. Recent Acquisitions (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note4.RecentAcquisitionsDetails1" id="Note4.RecentAcquisitionsDetails1">
	  <link:definition>00000037 - Disclosure - Note 4. Recent Acquisitions (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note4.RecentAcquisitionsDetails2" id="Note4.RecentAcquisitionsDetails2">
	  <link:definition>00000038 - Disclosure - Note 4. Recent Acquisitions (Details 2)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note4.RecentAcquisitionsDetails3" id="Note4.RecentAcquisitionsDetails3">
	  <link:definition>00000039 - Disclosure - Note 4. Recent Acquisitions (Details 3)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note5.GoodwillAndOtherIntangibleAssetsDetails" id="Note5.GoodwillAndOtherIntangibleAssetsDetails">
	  <link:definition>00000040 - Disclosure - Note 5. Goodwill and Other Intangible Assets (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note5.GoodwillAndOtherIntangibleAssetsDetails1" id="Note5.GoodwillAndOtherIntangibleAssetsDetails1">
	  <link:definition>00000041 - Disclosure - Note 5. Goodwill and Other Intangible Assets (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note5.GoodwillAndOtherIntangibleAssetsDetailsNarrative" id="Note5.GoodwillAndOtherIntangibleAssetsDetailsNarrative">
	  <link:definition>00000042 - Disclosure - Note 5. Goodwill and Other Intangible Assets (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note6.LineOfCreditDetailsNarrative" id="Note6.LineOfCreditDetailsNarrative">
	  <link:definition>00000043 - Disclosure - Note 6. Line of Credit (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note7.EquityDetails" id="Note7.EquityDetails">
	  <link:definition>00000044 - Disclosure - Note 7. Equity (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note7.EquityDetailsNarrative" id="Note7.EquityDetailsNarrative">
	  <link:definition>00000045 - Disclosure - Note 7. Equity (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnitsDetails" id="Note8.StockOptionsAndRestrictedStockUnitsDetails">
	  <link:definition>00000046 - Disclosure - Note 8. Stock Options and Restricted Stock Units (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnitsDetails1" id="Note8.StockOptionsAndRestrictedStockUnitsDetails1">
	  <link:definition>00000047 - Disclosure - Note 8. Stock Options and Restricted Stock Units (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnitsDetails2" id="Note8.StockOptionsAndRestrictedStockUnitsDetails2">
	  <link:definition>00000048 - Disclosure - Note 8. Stock Options and Restricted Stock Units (Details 2)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnitsDetails3" id="Note8.StockOptionsAndRestrictedStockUnitsDetails3">
	  <link:definition>00000049 - Disclosure - Note 8. Stock Options and Restricted Stock Units (Details 3)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note8.StockOptionsAndRestrictedStockUnitsDetailsNarrative" id="Note8.StockOptionsAndRestrictedStockUnitsDetailsNarrative">
	  <link:definition>00000050 - Disclosure - Note 8. Stock Options and Restricted Stock Units (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note9.LeasesDetails" id="Note9.LeasesDetails">
	  <link:definition>00000051 - Disclosure - Note 9. Leases (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note9.LeasesDetails1" id="Note9.LeasesDetails1">
	  <link:definition>00000052 - Disclosure - Note 9. Leases (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note9.LeasesDetailsNarrative" id="Note9.LeasesDetailsNarrative">
	  <link:definition>00000053 - Disclosure - Note 9. Leases (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note10.RevenuesDetails" id="Note10.RevenuesDetails">
	  <link:definition>00000054 - Disclosure - Note 10. Revenues (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note10.RevenuesDetails1" id="Note10.RevenuesDetails1">
	  <link:definition>00000055 - Disclosure - Note 10. Revenues (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note11.IncomeTaxesDetails" id="Note11.IncomeTaxesDetails">
	  <link:definition>00000056 - Disclosure - Note 11. Income Taxes (Details)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note11.IncomeTaxesDetails1" id="Note11.IncomeTaxesDetails1">
	  <link:definition>00000057 - Disclosure - Note 11. Income Taxes (Details 1)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note11.IncomeTaxesDetails2" id="Note11.IncomeTaxesDetails2">
	  <link:definition>00000058 - Disclosure - Note 11. Income Taxes (Details 2)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:roleType roleURI="http://issuerdirect.com/role/Note12.EmployeeBenefitPlanDetailsNarrative" id="Note12.EmployeeBenefitPlanDetailsNarrative">
	  <link:definition>00000059 - Disclosure - Note 12. Employee Benefit Plan (Details Narrative)</link:definition>
	  <link:usedOn>link:presentationLink</link:usedOn>
	  <link:usedOn>link:calculationLink</link:usedOn>
	  <link:usedOn>link:definitionLink</link:usedOn>
	</link:roleType>
	<link:linkbaseRef xlink:type="simple" xlink:href="isdr-20191231_pre.xml" xlink:role="http://www.xbrl.org/2003/role/presentationLinkbaseRef" xlink:arcrole="http://www.w3.org/1999/xlink/properties/linkbase" xlink:title="Presentation Links" />
	<link:linkbaseRef xlink:type="simple" xlink:href="isdr-20191231_lab.xml" xlink:role="http://www.xbrl.org/2003/role/labelLinkbaseRef" xlink:arcrole="http://www.w3.org/1999/xlink/properties/linkbase" xlink:title="Label Links" />
	<link:linkbaseRef xlink:type="simple" xlink:href="isdr-20191231_cal.xml" xlink:role="http://www.xbrl.org/2003/role/calculationLinkbaseRef" xlink:arcrole="http://www.w3.org/1999/xlink/properties/linkbase" xlink:title="Calculation Links" />
	<link:linkbaseRef xlink:type="simple" xlink:href="isdr-20191231_def.xml" xlink:role="http://www.xbrl.org/2003/role/definitionLinkbaseRef" xlink:arcrole="http://www.w3.org/1999/xlink/properties/linkbase" xlink:title="Definition Links" />
      </appinfo>
    </annotation>
    <import namespace="http://www.xbrl.org/2003/instance" schemaLocation="http://www.xbrl.org/2003/xbrl-instance-2003-12-31.xsd" />
    <import namespace="http://www.xbrl.org/2003/linkbase" schemaLocation="http://www.xbrl.org/2003/xbrl-linkbase-2003-12-31.xsd" />
    <import namespace="http://xbrl.sec.gov/dei/2018-01-31" schemaLocation="https://xbrl.sec.gov/dei/2018/dei-2018-01-31.xsd" />
    <import namespace="http://fasb.org/us-gaap/2019-01-31" schemaLocation="http://xbrl.fasb.org/us-gaap/2019/elts/us-gaap-2019-01-31.xsd" />
    <import namespace="http://fasb.org/us-types/2019-01-31" schemaLocation="http://xbrl.fasb.org/us-gaap/2019/elts/us-types-2019-01-31.xsd" />
    <import namespace="http://www.xbrl.org/dtr/type/non-numeric" schemaLocation="http://www.xbrl.org/dtr/type/nonNumeric-2009-12-16.xsd" />
    <import namespace="http://www.xbrl.org/dtr/type/numeric" schemaLocation="http://www.xbrl.org/dtr/type/numeric-2009-12-16.xsd" />
    <import namespace="http://xbrl.sec.gov/country/2017-01-31" schemaLocation="https://xbrl.sec.gov/country/2017/country-2017-01-31.xsd" />
    <import namespace="http://fasb.org/srt/2019-01-31" schemaLocation="http://xbrl.fasb.org/srt/2019/elts/srt-2019-01-31.xsd" />
    <element id="ISDR_CapitalizedComputerSoftwareAmortizationCostOfRevenues" name="CapitalizedComputerSoftwareAmortizationCostOfRevenues" nillable="true" xbrli:periodType="duration" xbrli:balance="credit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_NumberOfOptionsExercised" name="NumberOfOptionsExercised" nillable="true" xbrli:periodType="duration" type="xbrli:sharesItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_RangeOfExercisePriceOptionsOutstandingBeginning" name="RangeOfExercisePriceOptionsOutstandingBeginning" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_RangeOfExercisePriceOptionsGranted" name="RangeOfExercisePriceOptionsGranted" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_RangeOfExercisePriceOptionsExercised" name="RangeOfExercisePriceOptionsExercised" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_RangeOfExercisePriceOptionsForfeited" name="RangeOfExercisePriceOptionsForfeited" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_RangeOfExercisePriceOptionsOutstandingEnding" name="RangeOfExercisePriceOptionsOutstandingEnding" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ShareBasedCompensationArrangementByShareBasedPaymentAwardNonOptionEquityInstrumentsOutstandingWeightedAverageExercisePrice" name="ShareBasedCompensationArrangementByShareBasedPaymentAwardNonOptionEquityInstrumentsOutstandingWeightedAverageExercisePrice" nillable="true" xbrli:periodType="instant" type="num:perShareItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ShareBasedCompensationArrangementsByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsGrantsInPeriodWeightedAverageExercisePrice" name="ShareBasedCompensationArrangementsByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsGrantsInPeriodWeightedAverageExercisePrice" nillable="true" xbrli:periodType="duration" type="num:perShareItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ShareBasedCompensationArrangementsByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriodWeightedAverageExercisePrice" name="ShareBasedCompensationArrangementsByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriodWeightedAverageExercisePrice" nillable="true" xbrli:periodType="duration" type="num:perShareItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ShareBasedCompensationArrangementsByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeituresInPeriodWeightedAverageExercisePrice" name="ShareBasedCompensationArrangementsByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeituresInPeriodWeightedAverageExercisePrice" nillable="true" xbrli:periodType="duration" type="num:perShareItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_SharebasedCompensationArrangementBySharebasedPaymentAwardEquityInstrumentsOtherThanOptionsAggregateIntrinsicValueGranted" name="SharebasedCompensationArrangementBySharebasedPaymentAwardEquityInstrumentsOtherThanOptionsAggregateIntrinsicValueGranted" nillable="true" xbrli:periodType="duration" xbrli:balance="debit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_SharebasedCompensationArrangementBySharebasedPaymentAwardEquityInstrumentsOtherThanOptionsAggregateIntrinsicValueForfeited" name="SharebasedCompensationArrangementBySharebasedPaymentAwardEquityInstrumentsOtherThanOptionsAggregateIntrinsicValueForfeited" nillable="true" xbrli:periodType="duration" xbrli:balance="debit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_PermanentDifferenceStockbasedCompensation" name="PermanentDifferenceStockbasedCompensation" nillable="true" xbrli:periodType="duration" type="num:percentItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_EffectiveIncomeTaxRateReconciliationForeignEarningsTaxReform" name="EffectiveIncomeTaxRateReconciliationForeignEarningsTaxReform" nillable="true" xbrli:periodType="duration" type="num:percentItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_IncomeTaxReconciliationForeignEarningsTaxReform" name="IncomeTaxReconciliationForeignEarningsTaxReform" nillable="true" xbrli:periodType="duration" xbrli:balance="debit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ScheduleOfEstimatedUsefulLivesForPropertyAndEquipment" name="ScheduleOfEstimatedUsefulLivesForPropertyAndEquipment" nillable="true" xbrli:periodType="duration" type="nonnum:textBlockItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_CustomerRelationships" name="CustomerRelationships" nillable="true" xbrli:periodType="instant" xbrli:balance="debit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_DistributionPartnerRelationships" name="DistributionPartnerRelationships" nillable="true" xbrli:periodType="instant" xbrli:balance="debit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ExercisePriceRange" name="ExercisePriceRange" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_OperatingLeasesFutureMinimumPaymentsPresentValueOfNetMinimumPayments" name="OperatingLeasesFutureMinimumPaymentsPresentValueOfNetMinimumPayments" nillable="true" xbrli:periodType="instant" xbrli:balance="credit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TransactionCosts" name="TransactionCosts" nillable="true" xbrli:periodType="instant" xbrli:balance="debit" type="xbrli:monetaryItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_SECComplianceServicesMember" name="SECComplianceServicesMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption1Member" name="StockOption1Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption2Member" name="StockOption2Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption3Member" name="StockOption3Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption4Member" name="StockOption4Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption5Member" name="StockOption5Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption6Member" name="StockOption6Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TotalMember" name="TotalMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption7Member" name="StockOption7Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_StockOption8Member" name="StockOption8Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_DisclosureManagementMember" name="DisclosureManagementMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ShareholderCommunicationsMember" name="ShareholderCommunicationsMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_FulfillmentAndDistributionMember" name="FulfillmentAndDistributionMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_SoftwareLicensingMember" name="SoftwareLicensingMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TransferAgentServicesMember" name="TransferAgentServicesMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_PrecisionIRMember" name="PrecisionIRMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_AccesswireMember" name="AccesswireMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TotalIntangibleAssetsMember" name="TotalIntangibleAssetsMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ChangeMember" name="ChangeMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_RedOakMember" name="RedOakMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_PlatformAndTechnologyMember" name="PlatformAndTechnologyMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TrademarksDefiniteMember" name="TrademarksDefiniteMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_CanadaMember" name="CanadaMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_ServicesMember" name="ServicesMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TwentyFourteenPlanMember" name="TwentyFourteenPlanMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_TwentyTenPlanMember" name="TwentyTenPlanMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_InterwestTransferCompanyIncMember" name="InterwestTransferCompanyIncMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_FilingServicesCanadaIncMember" name="FilingServicesCanadaIncMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_DistributionPartnerRelationshipsMember" name="DistributionPartnerRelationshipsMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_VisualWebcasterPlatformMember" name="VisualWebcasterPlatformMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_AsAdjustedMember" name="AsAdjustedMember" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_Range1Member" name="Range1Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_Range2Member" name="Range2Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_Range3Member" name="Range3Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_Range4Member" name="Range4Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_Range5Member" name="Range5Member" nillable="true" xbrli:periodType="duration" type="nonnum:domainItemType" substitutionGroup="xbrli:item" />
    <element id="ISDR_DocumentAndEntityInformationAbstract" name="DocumentAndEntityInformationAbstract" abstract="true" nillable="true" xbrli:periodType="duration" type="xbrli:stringItemType" substitutionGroup="xbrli:item" />
</schema>
`

func TestDecodeSchemaFile(t *testing.T) {
	stimulus := []byte(testSchema)
	decoded, err := serializables.DecodeSchemaFile(stimulus)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	if len(decoded.Import) != 9 {
		t.Fatalf("expected 9 Import; outcome %d;\n", len(decoded.Import))
	}
	if len(decoded.Annotation) != 1 {
		t.Fatalf("expected 1 Annotation; outcome %d;\n", len(decoded.Annotation))
	}
	if len(decoded.Element) != 59 {
		t.Fatalf("expected 59 Element; outcome %d;\n", len(decoded.Element))
	}
}

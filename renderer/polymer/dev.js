console.log("mounting concept network browser development page")

const app = window["telefacts-renderer"]
app.changeLanguage = (v) => {
    console.log(v)
}
app.changeEntity = (v) => {
    console.log(v)
    app.selectedEntity = v
}
app.changeRelationshipSet = (v) => {
    console.log(v)
    app.selectedRelationshipSet = v
}
app.isLoading = true
setTimeout(
    () => {
        app.entities = [ "me", "myself", "I"]
        app.relationshipSets = [
            "http://test",
            "http://foo",
            "http://bar",
            "http://baq",
            "http://baz",
            "http://face-roo0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0olled",
            "http://helloworld"
        ]
        app.isLoading = false
        app.pGrid= {
            "IndentedLabels": [
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementOfStockholdersEquityAbstract",
                    "Indentation": 0
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementTable",
                    "Indentation": 1
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementEquityComponentsAxis",
                    "Indentation": 2
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_PreferredStockMember",
                    "Indentation": 3
                },
                {
                    "Href": "fizz-20200502.xsd#fizz_CommonStockOutstandingMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AdditionalPaidInCapitalMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RetainedEarningsMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AccumulatedOtherComprehensiveIncomeMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_TreasuryStockMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_TreasuryStockCommonMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_EquityComponentDomain",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementClassOfStockAxis",
                    "Indentation": 2
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_SeriesCPreferredStockMember",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ClassOfStockDomain",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StatementLineItems",
                    "Indentation": 2
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_SharesOutstanding",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StockholdersEquity",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StockIssuedDuringPeriodSharesStockOptionsExercised",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StockIssuedDuringPeriodValueStockOptionsExercised",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AdjustmentsToAdditionalPaidInCapitalSharebasedCompensationRequisiteServicePeriodRecognitionValue",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_NetIncomeLoss",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_DividendsCommonStockCash",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OtherComprehensiveIncomeLossDerivativesQualifyingAsHedgesNetOfTax",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OtherComprehensiveIncomeOtherNetOfTax",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_TreasuryStockSharesAcquired",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_TreasuryStockValueAcquiredCostMethod",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_SharesOutstanding",
                    "Indentation": 3
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StockholdersEquity",
                    "Indentation": 3
                }
            ],
            "MaxIndentation": 4,
            "RelevantContexts": [
                {
                    "ContextRef": "d_2012-05-01_2020-05-02",
                    "PeriodHeader": "2012-05-01/2020-05-02",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "i_2017-04-29_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2017-04-29_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2017-04-29_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2017-04-29_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2017-04-29_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2017-04-29_StatementClassOfStockAxis-SeriesCPreferredStockMember_StatementEquityComponentsAxis-PreferredStockMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "us-gaap:PreferredStockMember<us-gaap:StatementEquityComponentsAxis<segment",
                        "us-gaap:SeriesCPreferredStockMember<us-gaap:StatementClassOfStockAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2017-04-29_StatementClassOfStockAxis-SeriesCPreferredStockMember_StatementEquityComponentsAxis-TreasuryStockMember",
                    "PeriodHeader": "2017-04-29",
                    "DomainMemberHeaders": [
                        "us-gaap:SeriesCPreferredStockMember<us-gaap:StatementClassOfStockAxis<segment",
                        "us-gaap:TreasuryStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2017-04-30_2018-04-28",
                    "PeriodHeader": "2017-04-30/2018-04-28",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2017-04-30_2018-04-28_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2017-04-30/2018-04-28",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2017-04-30_2018-04-28_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2017-04-30/2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2017-04-30_2018-04-28_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2017-04-30/2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2017-04-30_2018-04-28_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2017-04-30/2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2017-04-30_2018-04-28_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2017-04-30/2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "i_2018-04-28_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28_StatementClassOfStockAxis-SeriesCPreferredStockMember_StatementEquityComponentsAxis-PreferredStockMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:PreferredStockMember<us-gaap:StatementEquityComponentsAxis<segment",
                        "us-gaap:SeriesCPreferredStockMember<us-gaap:StatementClassOfStockAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2018-04-28_StatementClassOfStockAxis-SeriesCPreferredStockMember_StatementEquityComponentsAxis-TreasuryStockMember",
                    "PeriodHeader": "2018-04-28",
                    "DomainMemberHeaders": [
                        "us-gaap:SeriesCPreferredStockMember<us-gaap:StatementClassOfStockAxis<segment",
                        "us-gaap:TreasuryStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2018-04-29_2018-07-28",
                    "PeriodHeader": "2018-04-29/2018-07-28",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2018-04-29_2019-04-27",
                    "PeriodHeader": "2018-04-29/2019-04-27",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2018-04-29_2019-04-27_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2018-04-29/2019-04-27",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2018-04-29_2019-04-27_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2018-04-29/2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2018-04-29_2019-04-27_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2018-04-29/2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2018-04-29_2019-04-27_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2018-04-29/2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2018-04-29_2019-04-27_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2018-04-29/2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2018-07-29_2018-10-27",
                    "PeriodHeader": "2018-07-29/2018-10-27",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2018-10-28_2019-01-26",
                    "PeriodHeader": "2018-10-28/2019-01-26",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2019-01-27_2019-04-27",
                    "PeriodHeader": "2019-01-27/2019-04-27",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "i_2019-04-27",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "i_2019-04-27_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2019-04-27_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2019-04-27_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2019-04-27_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2019-04-27_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2019-04-27_StatementClassOfStockAxis-SeriesCPreferredStockMember_StatementEquityComponentsAxis-PreferredStockMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:PreferredStockMember<us-gaap:StatementEquityComponentsAxis<segment",
                        "us-gaap:SeriesCPreferredStockMember<us-gaap:StatementClassOfStockAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2019-04-27_StatementClassOfStockAxis-SeriesCPreferredStockMember_StatementEquityComponentsAxis-TreasuryStockMember",
                    "PeriodHeader": "2019-04-27",
                    "DomainMemberHeaders": [
                        "us-gaap:SeriesCPreferredStockMember<us-gaap:StatementClassOfStockAxis<segment",
                        "us-gaap:TreasuryStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2019-04-28_2019-07-27",
                    "PeriodHeader": "2019-04-28/2019-07-27",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2019-04-28_2020-05-02",
                    "PeriodHeader": "2019-04-28/2020-05-02",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2019-04-28_2020-05-02_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2019-04-28/2020-05-02",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2019-04-28_2020-05-02_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2019-04-28/2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2019-04-28_2020-05-02_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2019-04-28/2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2019-04-28_2020-05-02_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2019-04-28/2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2019-04-28_2020-05-02_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2019-04-28/2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "d_2019-07-28_2019-10-26",
                    "PeriodHeader": "2019-07-28/2019-10-26",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2019-10-27_2020-01-25",
                    "PeriodHeader": "2019-10-27/2020-01-25",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "d_2020-01-26_2020-05-02",
                    "PeriodHeader": "2020-01-26/2020-05-02",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "i_2020-05-02",
                    "PeriodHeader": "2020-05-02",
                    "DomainMemberHeaders": []
                },
                {
                    "ContextRef": "i_2020-05-02_StatementEquityComponentsAxis-CommonStockOutstandingMember",
                    "PeriodHeader": "2020-05-02",
                    "DomainMemberHeaders": [
                        "fizz:CommonStockOutstandingMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2020-05-02_StatementEquityComponentsAxis-AccumulatedOtherComprehensiveIncomeMember",
                    "PeriodHeader": "2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2020-05-02_StatementEquityComponentsAxis-AdditionalPaidInCapitalMember",
                    "PeriodHeader": "2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2020-05-02_StatementEquityComponentsAxis-RetainedEarningsMember",
                    "PeriodHeader": "2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                },
                {
                    "ContextRef": "i_2020-05-02_StatementEquityComponentsAxis-TreasuryStockCommonMember",
                    "PeriodHeader": "2020-05-02",
                    "DomainMemberHeaders": [
                        "us-gaap:TreasuryStockCommonMember<us-gaap:StatementEquityComponentsAxis<segment"
                    ]
                }
            ],
            "MaxDepth": 2,
            "FactualQuadrant": [
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null",
                    "null"
                ],
                [
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share"
                ],
                [
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD"
                ],
                [
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share",
                    " 125000125000 Share"
                ],
                [
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD",
                    " 10001000 USD"
                ],
                [
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD",
                    " 125000125000 USD"
                ],
                [
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD",
                    " 129972000129972000 USD"
                ],
                [
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD",
                    " -0-0 USD"
                ],
                [
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD",
                    " -3673000-3673000 USD"
                ],
                [
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD",
                    " -204000-204000 USD"
                ],
                [
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share",
                    " 154000154000 Share"
                ],
                [
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD",
                    " 62330006233000 USD"
                ],
                [
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share",
                    " 150000150000 Share"
                ],
                [
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD",
                    " 452337000452337000 USD"
                ]
            ]
        }
        app.cGrid = {
            "SummationItems": [
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_Assets",
                    "RelevantContexts": [
                        {
                            "ContextRef": "ifc7bef0caea14d3ba6592c2ac3e05544_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": []
                        }
                    ],
                    "MaxDepth": 0,
                    "ContributingConcepts": [
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AssetsCurrent",
                            "IsSummationItem": true
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_PropertyPlantAndEquipmentAndFinanceLeaseRightOfUseAssetAfterAccumulatedDepreciationAndAmortization",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CapitalizedContractCostNetNoncurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_IntangibleAssetsNetExcludingGoodwill",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OtherAssetsNoncurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OperatingLeaseRightOfUseAsset",
                            "IsSummationItem": false
                        }
                    ],
                    "FactualQuadrant": [
                        [
                            "-3 609376000 usd",
                            "-3 609376000 usd"
                        ],
                        [
                            "-3 38231000 usd",
                            "-3 38231000 usd"
                        ],
                        [
                            "-3 17572000 usd",
                            "-3 17572000 usd"
                        ],
                        [
                            "-3 1633000 usd",
                            "-3 1633000 usd"
                        ],
                        [
                            "-3 4059000 usd",
                            "-3 4059000 usd"
                        ],
                        [
                            "-3 16631000 usd",
                            "-3 16631000 usd"
                        ],
                        [
                            "-3 687502000 usd",
                            "-3 687502000 usd"
                        ]
                    ]
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AssetsCurrent",
                    "RelevantContexts": [
                        {
                            "ContextRef": "ifc7bef0caea14d3ba6592c2ac3e05544_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i218d7f51e3104d6489e967663c8bf577_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CashAndCashEquivalentsMember<us-gaap:BalanceSheetLocationAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "iaa15d3881a0447a1865f6774aa91009a_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:MoneyMarketFundsMember<us-gaap:CashAndCashEquivalentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i2875e88fb17947d980c2d1ca93398cf1_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CashAndCashEquivalentsMember<us-gaap:BalanceSheetLocationAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "iaf05bb46951a418584487af3e8413a6d_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:MoneyMarketFundsMember<us-gaap:CashAndCashEquivalentsAxis<segment"
                            ]
                        }
                    ],
                    "MaxDepth": 1,
                    "ContributingConcepts": [
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CashAndCashEquivalentsAtCarryingValue",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AvailableForSaleSecuritiesDebtSecuritiesCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AccountsReceivableNetCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CapitalizedContractCostNetCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OtherReceivablesNetCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_PrepaidExpenseAndOtherAssetsCurrent",
                            "IsSummationItem": false
                        }
                    ],
                    "FactualQuadrant": [
                        [
                            "-3 426132000 usd",
                            "-3 426132000 usd",
                            "-3 426132000 usd",
                            "-3 426132000 usd",
                            "-3 426132000 usd",
                            "-3 426132000 usd"
                        ],
                        [
                            "-3 97750000 usd",
                            "-3 97750000 usd",
                            "-3 97750000 usd",
                            "-3 97750000 usd",
                            "-3 97750000 usd",
                            "-3 97750000 usd"
                        ],
                        [
                            "-3 55571000 usd",
                            "-3 55571000 usd",
                            "-3 55571000 usd",
                            "-3 55571000 usd",
                            "-3 55571000 usd",
                            "-3 55571000 usd"
                        ],
                        [
                            "-3 17919000 usd",
                            "-3 17919000 usd",
                            "-3 17919000 usd",
                            "-3 17919000 usd",
                            "-3 17919000 usd",
                            "-3 17919000 usd"
                        ],
                        [
                            "-3 2417000 usd",
                            "-3 2417000 usd",
                            "-3 2417000 usd",
                            "-3 2417000 usd",
                            "-3 2417000 usd",
                            "-3 2417000 usd"
                        ],
                        [
                            "-3 9587000 usd",
                            "-3 9587000 usd",
                            "-3 9587000 usd",
                            "-3 9587000 usd",
                            "-3 9587000 usd",
                            "-3 9587000 usd"
                        ],
                        [
                            "-3 609376000 usd",
                            "-3 609376000 usd",
                            "-3 609376000 usd",
                            "-3 609376000 usd",
                            "-3 609376000 usd",
                            "-3 609376000 usd"
                        ]
                    ]
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_Liabilities",
                    "RelevantContexts": [
                        {
                            "ContextRef": "ifc7bef0caea14d3ba6592c2ac3e05544_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": []
                        }
                    ],
                    "MaxDepth": 0,
                    "ContributingConcepts": [
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_LiabilitiesCurrent",
                            "IsSummationItem": true
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ContractWithCustomerLiabilityNoncurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OtherLiabilitiesNoncurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_FinanceLeaseLiabilityNoncurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_OperatingLeaseLiabilityNoncurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ConvertibleDebtNoncurrent",
                            "IsSummationItem": false
                        }
                    ],
                    "FactualQuadrant": [
                        [
                            "-3 251224000 usd",
                            "-3 251224000 usd"
                        ],
                        [
                            "-3 34756000 usd",
                            "-3 34756000 usd"
                        ],
                        [
                            "-3 1748000 usd",
                            "-3 1748000 usd"
                        ],
                        [
                            "-3 14805000 usd",
                            "-3 14805000 usd"
                        ],
                        [
                            "-3 18229000 usd",
                            "-3 18229000 usd"
                        ],
                        [
                            "-3 287242000 usd",
                            "-3 287242000 usd"
                        ],
                        [
                            "-3 608004000 usd",
                            "-3 608004000 usd"
                        ]
                    ]
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_LiabilitiesAndStockholdersEquity",
                    "RelevantContexts": [
                        {
                            "ContextRef": "i47ea73bcb5754cbca7c6b0d052259770_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i949280f97b46447d808448a4d135dcfe_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifbc749d136f54a79b1128e3a8dfc4dcd_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "id72b6a17891b468182b4b758d64cb038_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ic02e4353c3574bfca095888ed2932385_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i16c168bd1ca5441bbe2e70e5d7a67663_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i05f300f5f5484007933ae7327dab243e_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i125f51feb6f44f3aa14b2e0c367d7422_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ib3f4b5d3cdd6414aad02ec2cc506984a_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i47f2c8fd7c484cb3870db10862a386c8_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifadb885b7fb24117943cb0923ad675ae_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ide03cc14690e4e28b7b9170453347920_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i52db3004417f4630bcc992f878399e3f_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i4016c0abfdfe4631b74e499355e2d3e3_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "id0a26532604f4685a596711053c2a439_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i6ffea30d3d684b7cbccceebb467c5e3d_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ie12a90785b9340bb9995e44da1dfb677_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i5063c56ba41840f49f816ba577786e8e_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i4331450962814ab5b7b3feef4e4693ff_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i9791e892964647a9938fd24a0c5a6b7a_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifc7bef0caea14d3ba6592c2ac3e05544_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "if41369f6c74649ceb006505fe6609812_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ie9bfe5e3c7fe4c96a1788b87e5f94060_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i16760d69ed714519ae8b641bbb6bf1cc_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i59974f39622748c79c61c913495d8d8f_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ia81481a1ea214d95b604253db58f43d2_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ifc7a8aa9085e43168782edd84bf5279e_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "if816f3ade512446a8e7734182f9a94b1_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i57fff9a4738742189c82d9329a8de8f3_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i8a9240abd7674271a4e9a8666df32fcf_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ie171d0036a744014a783d2b9b3236286_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ied76562fd9f64fb0b1a117b44ef36e0e_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i1a9a359e28dc4727987058b24b6cd3e6_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i083aff72d71f48cfa1cde2b90e7527e2_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "iad72809fe36c47eaa61653938393b803_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i51c3ad1bf7bd427dad0d404b25450566_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifcc12ac90bc24c4e9e35102f04d525f4_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "id8a2923d3ead4743a92b2d3c20b3c33f_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ia30f60012fba4454ae6bc2c3588ae312_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        }
                    ],
                    "MaxDepth": 1,
                    "ContributingConcepts": [
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_Liabilities",
                            "IsSummationItem": true
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StockholdersEquity",
                            "IsSummationItem": true
                        }
                    ],
                    "FactualQuadrant": [
                        [
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd",
                            "-3 608004000 usd"
                        ],
                        [
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd"
                        ],
                        [
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd",
                            "-3 687502000 usd"
                        ]
                    ]
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_LiabilitiesCurrent",
                    "RelevantContexts": [
                        {
                            "ContextRef": "ifc7bef0caea14d3ba6592c2ac3e05544_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": []
                        }
                    ],
                    "MaxDepth": 0,
                    "ContributingConcepts": [
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AccountsPayableCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AccruedLiabilitiesCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_ContractWithCustomerLiabilityCurrent",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_FinanceLeaseLiabilityCurrent",
                            "IsSummationItem": false
                        }
                    ],
                    "FactualQuadrant": [
                        [
                            "-3 3651000 usd",
                            "-3 3651000 usd"
                        ],
                        [
                            "-3 63444000 usd",
                            "-3 63444000 usd"
                        ],
                        [
                            "-3 182700000 usd",
                            "-3 182700000 usd"
                        ],
                        [
                            "-3 1429000 usd",
                            "-3 1429000 usd"
                        ],
                        [
                            "-3 251224000 usd",
                            "-3 251224000 usd"
                        ]
                    ]
                },
                {
                    "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_StockholdersEquity",
                    "RelevantContexts": [
                        {
                            "ContextRef": "i47ea73bcb5754cbca7c6b0d052259770_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i949280f97b46447d808448a4d135dcfe_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifbc749d136f54a79b1128e3a8dfc4dcd_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "id72b6a17891b468182b4b758d64cb038_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ic02e4353c3574bfca095888ed2932385_I20181231",
                            "PeriodHeader": "2018-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i16c168bd1ca5441bbe2e70e5d7a67663_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i05f300f5f5484007933ae7327dab243e_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i125f51feb6f44f3aa14b2e0c367d7422_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ib3f4b5d3cdd6414aad02ec2cc506984a_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i47f2c8fd7c484cb3870db10862a386c8_I20190331",
                            "PeriodHeader": "2019-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifadb885b7fb24117943cb0923ad675ae_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ide03cc14690e4e28b7b9170453347920_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i52db3004417f4630bcc992f878399e3f_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i4016c0abfdfe4631b74e499355e2d3e3_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "id0a26532604f4685a596711053c2a439_I20190630",
                            "PeriodHeader": "2019-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i6ffea30d3d684b7cbccceebb467c5e3d_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ie12a90785b9340bb9995e44da1dfb677_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i5063c56ba41840f49f816ba577786e8e_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i4331450962814ab5b7b3feef4e4693ff_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i9791e892964647a9938fd24a0c5a6b7a_I20190930",
                            "PeriodHeader": "2019-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifc7bef0caea14d3ba6592c2ac3e05544_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "if41369f6c74649ceb006505fe6609812_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ie9bfe5e3c7fe4c96a1788b87e5f94060_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i61ca66652fe34feaa9ae9d9f06b895a8_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonClassAMember<us-gaap:StatementClassOfStockAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i420f5189459e413f8a02774d5de26f2c_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonClassBMember<us-gaap:StatementClassOfStockAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i16760d69ed714519ae8b641bbb6bf1cc_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i59974f39622748c79c61c913495d8d8f_I20191231",
                            "PeriodHeader": "2019-12-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ia81481a1ea214d95b604253db58f43d2_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ifc7a8aa9085e43168782edd84bf5279e_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "if816f3ade512446a8e7734182f9a94b1_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i57fff9a4738742189c82d9329a8de8f3_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i8a9240abd7674271a4e9a8666df32fcf_I20200331",
                            "PeriodHeader": "2020-03-31",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ie171d0036a744014a783d2b9b3236286_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "ied76562fd9f64fb0b1a117b44ef36e0e_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i1a9a359e28dc4727987058b24b6cd3e6_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i083aff72d71f48cfa1cde2b90e7527e2_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "iad72809fe36c47eaa61653938393b803_I20200630",
                            "PeriodHeader": "2020-06-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ic239d0e3c5dc4af2b6820cb7d18e3dac_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": []
                        },
                        {
                            "ContextRef": "i51c3ad1bf7bd427dad0d404b25450566_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AccumulatedOtherComprehensiveIncomeMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ifcc12ac90bc24c4e9e35102f04d525f4_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:AdditionalPaidInCapitalMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "i954976e5852948c59359b0616385aa77_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonClassAMember<us-gaap:StatementClassOfStockAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "if31567373f814614a49ba40eae4d3ee7_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonClassBMember<us-gaap:StatementClassOfStockAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "id8a2923d3ead4743a92b2d3c20b3c33f_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:CommonStockMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        },
                        {
                            "ContextRef": "ia30f60012fba4454ae6bc2c3588ae312_I20200930",
                            "PeriodHeader": "2020-09-30",
                            "DomainMemberHeaders": [
                                "us-gaap:RetainedEarningsMember<us-gaap:StatementEquityComponentsAxis<segment"
                            ]
                        }
                    ],
                    "MaxDepth": 1,
                    "ContributingConcepts": [
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_CommonStockValue",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_PreferredStockValue",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AdditionalPaidInCapital",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_RetainedEarningsAccumulatedDeficit",
                            "IsSummationItem": false
                        },
                        {
                            "Sign": "+",
                            "Scale": "1.0",
                            "Href": "http://xbrl.fasb.org/us-gaap/2020/elts/us-gaap-2020-01-31.xsd#us-gaap_AccumulatedOtherComprehensiveIncomeLossNetOfTax",
                            "IsSummationItem": false
                        }
                    ],
                    "FactualQuadrant": [
                        [
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd",
                            "-3 40000 usd"
                        ],
                        [
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd",
                            "-3 0 usd"
                        ],
                        [
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd",
                            "-3 474969000 usd"
                        ],
                        [
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd",
                            "-3 -396006000 usd"
                        ],
                        [
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd",
                            "-3 487000 usd"
                        ],
                        [
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd",
                            "-3 79498000 usd"
                        ]
                    ]
                }
            ]
        }
    },
    2000
)
package renderables

import "ecksbee.com/telefacts/pkg/hydratables"

type Expressable struct {
	Href    string
	Labels  LabelPack
	Context struct {
		Period LanguagePack
		VoidQuadrant
		ContextualMemberGrid
	}
	Measurement string
	Precision   string
	Footnotes   []string
}

func getExpressions(h *hydratables.Hydratable, conceptFinder ConceptFinder, measurementFinder MeasurementFinder, factFinder FactFinder) (map[string]Expressable, error) {
	ret := make(map[string]Expressable)
	source := h.Document
	if source != nil {
		for _, item := range source.Expressions {
			id := item.ID
			href, _, err := conceptFinder.NameQuery(item.Name.Space, item.Name.Local)
			if err != nil {
				return nil, err
			}
			if id == "" {
				continue
			}
			var extracted *hydratables.Instance
			for _, instance := range h.Instances {
				extracted = &instance
				break
			}
			context := getContext(extracted, item.ContextRef)
			period := periodString(context)
			entity := stringify(&Entity{
				Scheme:   context.Entity.Identifier.Scheme,
				CharData: context.Entity.Identifier.CharData,
			})
			relevantContexts, segment, scenario, _ := getRelevantContexts(entity, h, []string{
				href,
			})
			expressedContexts := make([]relevantContext, 0)
			for _, relevantContext := range relevantContexts {
				if relevantContext.ContextRef == item.ContextRef {
					expressedContexts = append(expressedContexts, relevantContext)
					break
				}
			}
			memberGrid, voidQuadrant := getMemberGridAndVoidQuadrant(expressedContexts, segment, scenario)
			numerator, denominator := measurementFinder.FindMeasurement(item.UnitRef)
			var measurementExpression string
			if numerator != nil {
				if numerator.Symbol != "" {
					measurementExpression = numerator.Symbol
				} else {
					measurementExpression = numerator.UnitName
				}
				if denominator != nil {
					measurementExpression += "/"
					if denominator.Symbol != "" {
						measurementExpression += denominator.Symbol
					} else {
						measurementExpression += denominator.UnitName
					}
				}
			}
			hydratedFact := factFinder.FindFact(href, item.ContextRef)
			footnotes := factFinder.GetFootnotes(hydratedFact)
			footnoteTexts := make([]string, len(footnotes))
			for _, footnote := range footnotes {
				footnoteTexts = append(footnoteTexts, footnote.CharData)
			}
			ret[id] = Expressable{
				Href:   href,
				Labels: GetLabel(h, href),
				Context: struct {
					Period LanguagePack
					VoidQuadrant
					ContextualMemberGrid
				}{
					Period:               period,
					VoidQuadrant:         voidQuadrant,
					ContextualMemberGrid: memberGrid,
				},
				Measurement: measurementExpression,
				Precision:   expressPrecision(item.Precision),
				Footnotes:   footnoteTexts,
			}
		}
	}
	return ret, nil
}

func expressPrecision(p hydratables.Precision) string {
	pmap := map[hydratables.Precision]string{
		hydratables.Exact:             "Exact",
		hydratables.Precisionless:     "Precisionless",
		hydratables.Trillions:         "Trillions (-12)",
		hydratables.HundredBillions:   "Hundred Billions (-11)",
		hydratables.TenBillions:       "Ten Billions (-10)",
		hydratables.Billions:          "Billions (-9)",
		hydratables.HundredMillions:   "Hundred Millions (-8)",
		hydratables.TenMillions:       "Ten Millions (-7)",
		hydratables.Millions:          "Millions (-6)",
		hydratables.HundredThousands:  "Hundred Thousands (-5)",
		hydratables.TenThousands:      "Ten Thousands (-4)",
		hydratables.Thousands:         "Thousands (-3)",
		hydratables.Hundreds:          "Hundreds (-2)",
		hydratables.Tens:              "Tens (-1)",
		hydratables.Oneth:             "Ones (0)",
		hydratables.Tenth:             "Tenth (1)",
		hydratables.Hundredth:         "Hundredth (2)",
		hydratables.Thousandth:        "Thousandth (3)",
		hydratables.TenThousandth:     "Ten Thousandth (4)",
		hydratables.HundredThousandth: "Hundred Thousandth (5)",
		hydratables.Millionth:         "Millionth (6)",
		hydratables.TenMillionth:      "Ten Millionth (7)",
		hydratables.HundredMillionth:  "Hundred Millionth (8)",
		hydratables.Billionth:         "Billionth (9)",
		hydratables.TenBillionth:      "Ten Billionth (10)",
		hydratables.HundredBillionth:  "Hundred Billionth (11)",
		hydratables.Trillionth:        "Trillionth (12)",
	}
	if ret, ok := pmap[p]; ok {
		return ret
	}
	return ""
}

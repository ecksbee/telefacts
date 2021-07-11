package renderables

import (
	"strconv"
	"time"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
	"github.com/klauspost/lctime"
)

func formatSpanishDates(start string, end string) string {
	isDuration := start != ""
	endTime, err := time.Parse(iso8601, end)
	if err != nil {
		return defaultDate(start, end)
	}
	date, err := lctime.StrftimeLoc("es_ES", "%d %B %Y", endTime)
	if err != nil {
		return defaultDate(start, end)
	}
	if isDuration {
		startTime, err := time.Parse(iso8601, start)
		if err != nil {
			return defaultDate(start, end)
		}
		months := monthCount(startTime, endTime)
		if months < 1 {
			return defaultDate(start, end)
		} else if months == 1 {
			return "un mes terminado al " + date
		} else {
			m := strconv.Itoa(months)
			return m + " meses terminados al " + date
		}
	}
	return "al " + date
}

func renderSpanishFact(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder) *FactExpression {
	_, concept, err := cf.HashQuery(fact.Href)
	if err != nil {
		return &FactExpression{
			Core: "error",
		}
	}
	textBlock := renderTextBlock(fact, cf, mf)
	if textBlock != nil {
		return textBlock
	}
	if fact.Precision == hydratables.Precisionless {
		return &FactExpression{
			Core: "NaN",
		}
	}
	isPercent := concept.Type.Space == attr.NUM &&
		concept.Type.Local == attr.PercentItemType
	numerator, denominator := mf.FindMeasurement(fact.UnitRef)
	sigFig, err := SigFigs(fact.XMLInner, fact.Precision, concept, ' ')
	if err != nil {
		return &FactExpression{
			Core: "NaN",
		}
	}
	if numerator != nil {
		if numerator.Symbol != "" {
			sigFig.Head = numerator.Symbol + " " + sigFig.Head
		}
	}
	if isPercent {
		sigFig.Tail += "%"
	}
	if numerator != nil {
		sigFig.Tail += " "
		if numerator.Symbol == "" && !isPercent {
			sigFig.Tail += numerator.UnitName
		}
		if denominator != nil {
			sigFig.Tail += "/"
			if denominator.Symbol != "" {
				sigFig.Tail += denominator.Symbol
			} else {
				sigFig.Tail += denominator.UnitName
			}
		}
	}

	return &FactExpression{
		Head: sigFig.Head,
		Core: sigFig.Core,
		Tail: sigFig.Tail,
	}
}

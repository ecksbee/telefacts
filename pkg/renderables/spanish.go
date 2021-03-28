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

func renderSpanishFact(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder, labelRole LabelRole) FactExpression { //todo spanish fact expression
	_, concept, err := cf.HashQuery(fact.Href)
	if err != nil {
		return FactExpression{
			Head: "",
			Core: "error",
			Tail: "",
		}
	}
	isPercent := concept.Type.Space == attr.NUM &&
		concept.Type.Local == attr.PercentItemType
	head := ""
	numerator, denominator := mf.FindMeasurement(fact.UnitRef)
	if numerator != nil {
		if numerator.Symbol != "" {
			head += numerator.Symbol + " "
		}
	}
	core, tail := SigFigs(fact.XMLInner, fact.Precision, concept)
	if isPercent {
		tail += "%"
	}
	if numerator != nil {
		tail += " "
		if numerator.Symbol == "" {
			tail += numerator.UnitName
		}
		if denominator != nil {
			tail += "/"
			if denominator.Symbol != "" {
				tail += denominator.Symbol
			} else {
				tail += denominator.UnitName
			}
		}
	}

	return FactExpression{
		Head: head,
		Core: core,
		Tail: tail,
	}
}

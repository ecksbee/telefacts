package renderables

import "ecksbee.com/telefacts/pkg/hydratables"

type FactFinder interface {
	FindFact(href string, contextRef string) *hydratables.Fact
}

type MeasurementFinder interface {
	FindMeasurement(unitRef string) (*hydratables.Measurement, *hydratables.Measurement)
}

func render(fact *hydratables.Fact, mf MeasurementFinder, labelRoles []LabelRole, langs []Lang) MultilingualFact { //todo render for percentItemType concepts
	ret := MultilingualFact{}
	ret[Default] = make(map[Lang]FactExpression)
	if fact == nil {
		ret[Default][PureLabel] = FactExpression{
			Head: "",
			Core: "",
			Tail: "",
		}
		return ret
	}
	if fact.IsNil {
		ret[Default][PureLabel] = FactExpression{
			Head: "",
			Core: "nil",
			Tail: "",
		}
		return ret
	}
	if mf == nil {
		ret[Default][PureLabel] = FactExpression{
			Head: "",
			Core: "error",
			Tail: "",
		}
		return ret
	}
	var precision string
	if fact.Decimals != "" {
		precision = fact.Decimals
	} else {
		precision = fact.Precision
	}

	ret[Default][PureLabel] = FactExpression{
		Head: precision,
		Core: fact.XMLInner,
		Tail: fact.UnitRef,
	}
	tail := ""
	numerator, denominator := mf.FindMeasurement(fact.UnitRef)
	if numerator != nil {
		if numerator.Symbol != "" {
			tail += numerator.Symbol
		} else {
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
	for _, labelRole := range labelRoles {
		if _, found := ret[labelRole]; !found {
			ret[labelRole] = make(map[Lang]FactExpression)
		}
		for _, lang := range langs {
			if lang == PureLabel {
				continue
			}
			ret[labelRole][lang] = FactExpression{
				Head: precision,
				Core: fact.XMLInner,
				Tail: tail,
			} //todo
		}
	}
	return ret
}

package renderables

import "ecksbee.com/telefacts/hydratables"

type FactFinder interface {
	FindFact(href string, contextRef string) *hydratables.Fact
}

type MeasurementFinder interface {
	FindMeasurement(unitRef string) (*hydratables.Measurement, *hydratables.Measurement)
}

func render(fact *hydratables.Fact, mf MeasurementFinder, labelRoles []LabelRole, langs []Lang) LabelPack {
	ret := LabelPack{}
	ret[Default] = make(LanguagePack)
	if fact == nil {
		ret[Default][PureLabel] = ""
		return ret
	}
	if fact.IsNil {
		ret[Default][PureLabel] = "nil"
		return ret
	}
	if mf == nil {
		ret[Default][PureLabel] = "error"
		return ret
	}
	var precision string
	if fact.Decimals != "" {
		precision = fact.Decimals
	} else {
		precision = fact.Precision
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

	ret[Default][PureLabel] = precision + " " + fact.XMLInner + " " + fact.UnitRef
	for _, labelRole := range labelRoles {
		if _, found := ret[labelRole]; !found {
			ret[labelRole] = make(LanguagePack)
		}
		for _, lang := range langs {
			if lang == PureLabel {
				continue
			}
			ret[labelRole][lang] = precision + " " + fact.XMLInner + " " + tail //todo
		}
	}
	return ret
}

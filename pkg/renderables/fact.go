package renderables

import (
	"math"
	"math/big"
	"strconv"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
)

type FactFinder interface {
	FindFact(href string, contextRef string) *hydratables.Fact
}

type MeasurementFinder interface {
	FindMeasurement(unitRef string) (*hydratables.Measurement, *hydratables.Measurement)
}

func render(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder, labelRoles []LabelRole, langs []Lang) MultilingualFact {
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
	if fact.Precision != hydratables.Precisionless {
		if fact.Precision == hydratables.Exact {
			precision = "INF"
		} else {
			precision = strconv.Itoa(int(fact.Precision))
		}
	}

	ret[Default][PureLabel] = FactExpression{
		Head: precision,
		Core: fact.XMLInner,
		Tail: fact.UnitRef,
	}
	for _, labelRole := range labelRoles {
		if _, found := ret[labelRole]; !found {
			ret[labelRole] = make(map[Lang]FactExpression)
		}
		for _, lang := range langs {
			switch lang {
			case English:
				ret[labelRole][lang] = renderEnglishFact(fact, cf, mf, labelRole)
			case Español:
				ret[labelRole][lang] = renderEnglishFact(fact, cf, mf, labelRole) //todo spanish fact expression
			case PureLabel:
			default:
				continue
			}
		}
	}
	return ret
}

func SigFigs(value string, precision hydratables.Precision, concept *hydratables.Concept) (string, string) {
	f, _, err := big.ParseFloat(value, 10, 106, big.ToZero)
	if err != nil {
		return value, ""
	}
	original := f.Text('f', -1)
	point := strings.IndexRune(original, '.')
	isPercent := concept.Type.Space == attr.NUM &&
		concept.Type.Local == attr.PercentItemType
	n := int(precision)
	if isPercent {
		n -= 2
		if point > -1 {
			for len(original)-1-point < 2 {
				original += "0"
			}
			original = original[:point] + original[point+1:point+3] + "." + original[point+3:]
			f, _, err = big.ParseFloat(original, 10, 106, big.ToZero)
			if err != nil {
				return value, ""
			}
			original = f.Text('f', -1)
			point = strings.IndexRune(original, '.')
		} else {
			original += "00"
		}
	}
	if precision == hydratables.Exact || precision == hydratables.Precisionless {
		return original, ""
	}
	if point < 0 {
		if n < 1 {
			n = int(math.Abs(float64(n)))
			if n >= len(original) {
				return "", original
			}
			return original[:len(original)-n], original[len(original)-n:]
		} else {
			ret := original
			ret += "."
			n--
			for n > 0 {
				ret += "0"
				n--
			}
			return ret, ""
		}
	} else {
		if point == 0 {
			if n == 0 {
				return "0.", original[1:]
			}
			if n > 0 {
				n = int(math.Abs(float64(n)))
				if n >= len(original)-1 {
					return "", original
				}
				return "0" + original[:n+1], original[n+1:]
			}
			return "", "0" + original
		}
		if n == 0 {
			return original[:point], original[point:]
		}
		if n > 0 {
			n = int(math.Abs(float64(n)))
			if n >= len(original)-point-1 {
				return "", original
			}
			return original[:point+n], original[point+n:]
		}
		n = int(math.Abs(float64(n)))
		if n >= point {
			return "", original
		}
		return original[:point-n], original[point-n:]
	}
}

func renderEnglishFact(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder, labelRole LabelRole) FactExpression {
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

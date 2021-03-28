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

func render(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder, langs []Lang) MultilingualFact {
	ret := MultilingualFact{}
	if fact == nil {
		ret[PureLabel] = FactExpression{}
		return ret
	}
	if fact.IsNil {
		ret[PureLabel] = FactExpression{
			Core: "nil",
		}
		return ret
	}
	if mf == nil {
		ret[PureLabel] = FactExpression{
			Core: "error",
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
	core := fact.XMLInner
	if len(core) > 44 {
		core = core[:44]
	}
	ret[PureLabel] = FactExpression{
		Head: precision,
		Core: core,
		Tail: fact.UnitRef,
	}
	for _, lang := range langs {
		switch lang {
		case English:
			ret[lang] = renderEnglishFact(fact, cf, mf)
		case Español:
			ret[lang] = renderSpanishFact(fact, cf, mf)
		case PureLabel:
		default:
			continue
		}
	}
	return ret
}

func SigFigs(value string, precision hydratables.Precision, concept *hydratables.Concept) (string, string) {
	//todo decouple as its own numeric expression library
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
			return original[:point+1], original[point+1:]
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
		return original[:point+1-n], original[point+1-n:]
	}
}

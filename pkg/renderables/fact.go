package renderables

import (
	"fmt"
	"math/big"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"
	"github.com/joshuanario/digits"
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
	core := fact.XMLInner
	if len(core) > 44 {
		core = core[:44]
	}
	for _, lang := range langs {
		switch lang {
		case English:
			ret[lang] = renderEnglishFact(fact, cf, mf)
		case Español:
			ret[lang] = renderSpanishFact(fact, cf, mf)
		case PureLabel:
			ret[lang] = renderEnglishFact(fact, cf, mf)
		default:
			continue
		}
	}
	return ret
}

func SigFigs(value string, precision hydratables.Precision, concept *hydratables.Concept, g rune) (*FactExpression, error) {
	isPercent := concept.Type.Space == attr.NUM &&
		concept.Type.Local == attr.PercentItemType
	if isPercent {
		return renderPercent(value, precision, g)
	}
	return renderNumeric(value, precision, g)
}

func renderPercent(value string, precision hydratables.Precision, g rune) (*FactExpression, error) {
	f, _, err := big.ParseFloat(value, 10, digits.PREC_BITS, big.ToZero)
	if err != nil {
		return nil, fmt.Errorf("failed to convert fact to numeric expression")
	}
	percent := f.Mul(f, big.NewFloat(100).SetPrec(digits.PREC_BITS)).Text('f', -1)
	return renderNumeric(percent, precision, g)
}

func renderNumeric(value string, precision hydratables.Precision, g rune) (*FactExpression, error) {
	p := digits.Precision(int(precision))
	digit, err := digits.New(p, value, g, digits.PreserveUpToHundredth)
	if err != nil {
		return nil, fmt.Errorf("failed to convert fact to numeric expression")
	}
	return &FactExpression{
		Head: digit.Head(),
		Core: digit.Core(),
		Tail: digit.Tail(),
	}, nil
}

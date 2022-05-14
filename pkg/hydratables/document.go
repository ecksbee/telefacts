package hydratables

import (
	"strconv"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
	"github.com/antchfx/xmlquery"
)

type Expressable struct {
	Href       string
	ID         string
	ContextRef string
	UnitRef    string
	Precision  Precision
	IsNil      bool
	XMLInner   string
}

type Document struct {
	Expressions []Expressable
}

func HydrateDocument(folder *serializables.Folder) (*Document, error) {
	source := folder.Document
	if source == nil {
		return nil, nil
	}
	ret := make([]Expressable, 0)
	for _, nonNumeric := range source.NonNumerics {
		expressable := HydrateExpressable(nonNumeric)
		ret = append(ret, *expressable)
	}
	for _, nonFraction := range source.NonFractions {
		expressable := HydrateExpressable(nonFraction)
		ret = append(ret, *expressable)
	}
	return &Document{
		Expressions: ret,
	}, nil
}

func HydrateExpressable(node *xmlquery.Node) *Expressable {
	var idVal, unitRefVal string
	decimals := Precisionless
	href := "//todo"
	name := attr.FindXpathAttr(node.Attr, "name")
	contextRef := attr.FindXpathAttr(node.Attr, "contextRef")
	unitRef := attr.FindXpathAttr(node.Attr, "unitRef")
	decimalsAttr := attr.FindXpathAttr(node.Attr, "decimals")
	if decimalsAttr != nil {
		if decimalsAttr.Value == "INF" {
			decimals = Exact
		} else {
			decimalsInt, err := strconv.Atoi(decimalsAttr.Value)
			if err != nil {
				decimals = Precisionless
			} else {
				decimals = Precision(decimalsInt)
			}
		}
	}
	id := attr.FindXpathAttr(node.Attr, "id")
	if id == nil {
		if name != nil && contextRef != nil {
			idVal = name.Value + "_" + contextRef.Value
		}
	} else {
		idVal = id.Value
	}
	if unitRef != nil {
		unitRefVal = unitRef.Value
	}
	return &Expressable{
		Href:       href,
		ID:         idVal,
		ContextRef: contextRef.Value,
		UnitRef:    unitRefVal,
		Precision:  Precision(decimals),
		IsNil:      false, //todo
		XMLInner:   "",    //todo
	}
}

package hydratables

import (
	"encoding/hex"
	"encoding/xml"
	"hash/fnv"
	"strconv"
	"sync"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
	"github.com/antchfx/xmlquery"
)

type Expressable struct {
	Name       xml.Name
	ID         string
	ContextRef string
	UnitRef    string
	Precision  Precision
	IsNil      bool
}

type Document struct {
	Expressions []Expressable
}

func HydrateDocument(folder *serializables.Folder) (*Document, error) {
	source := folder.Document
	if source == nil {
		return nil, nil
	}
	np, err := attr.NewNameProvider(source.Html.Attr)
	if err != nil {
		return nil, err
	}
	ret := make([]Expressable, 0)
	eerrChan := make(chan error)
	var wg1 sync.WaitGroup
	wg1.Add(len(source.NonNumerics))
	for _, nnonNumeric := range source.NonNumerics {
		go func(nonNumeric *xmlquery.Node, errChan chan error) {
			expressable, err := HydrateExpressable(nonNumeric, np)
			if err != nil {
				errChan <- err
				return
			}
			ret = append(ret, *expressable)
			errChan <- nil
		}(nnonNumeric, eerrChan)
	}
	wg1.Wait()
	err = <-eerrChan
	if err != nil {
		return nil, err
	}
	var wg2 sync.WaitGroup
	wg2.Add(len(source.NonFractions))
	for _, nnonFraction := range source.NonFractions {
		go func(nonFraction *xmlquery.Node, errChan chan error) {
			expressable, err := HydrateExpressable(nonFraction, np)
			if err != nil {
				errChan <- err
				return
			}
			ret = append(ret, *expressable)
		}(nnonFraction, eerrChan)
	}
	wg2.Wait()
	err = <-eerrChan
	if err != nil {
		return nil, err
	}
	return &Document{
		Expressions: ret,
	}, nil
}

func HydrateExpressable(node *xmlquery.Node, np *attr.NameProvider) (*Expressable, error) {
	var idVal, unitRefVal string
	decimals := Precisionless
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
			h := fnv.New128a()
			h.Write([]byte(name.Value + "_" + contextRef.Value))
			idVal = hex.EncodeToString(h.Sum([]byte{}))
		}
	} else {
		idVal = id.Value
	}
	if unitRef != nil {
		unitRefVal = unitRef.Value
	}
	xmlName, err := np.ProvideXmlName(name.Value)
	if err != nil {
		return nil, err
	}
	nilAttr := attr.FindXpathAttr(node.Attr, "nil")
	isNil := false
	if nilAttr != nil {
		isNil, err = strconv.ParseBool(nilAttr.Value)
		if err != nil {
			isNil = false
		}
	}
	return &Expressable{
		Name:       *xmlName,
		ID:         idVal,
		ContextRef: contextRef.Value,
		UnitRef:    unitRefVal,
		Precision:  Precision(decimals),
		IsNil:      isNil,
	}, nil
}

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
	ret1 := make([]Expressable, 0)
	var lock1 sync.Mutex
	var goErr error
	var wg1 sync.WaitGroup
	wg1.Add(len(source.NonNumerics))
	for _, nnonNumeric := range source.NonNumerics {
		go func(nonNumeric *xmlquery.Node) {
			defer wg1.Done()
			if goErr != nil {
				return
			}
			expressable, err := HydrateExpressable(nonNumeric, np)
			if err != nil {
				goErr = err
				return
			}
			lock1.Lock()
			ret1 = append(ret1, *expressable)
			lock1.Unlock()
		}(nnonNumeric)
	}
	wg1.Wait()
	if goErr != nil {
		return nil, goErr
	}
	ret2 := make([]Expressable, 0)
	var lock2 sync.Mutex
	var goErr2 error
	var wg2 sync.WaitGroup
	wg2.Add(len(source.NonFractions))
	for _, nnonFraction := range source.NonFractions {
		go func(nonFraction *xmlquery.Node) {
			defer wg2.Done()
			if goErr2 != nil {
				return
			}
			expressable, err := HydrateExpressable(nonFraction, np)
			if err != nil {
				goErr2 = err
				return
			}
			lock2.Lock()
			ret2 = append(ret2, *expressable)
			lock2.Unlock()
		}(nnonFraction)
	}
	wg2.Wait()
	if goErr2 != nil {
		return nil, goErr2
	}
	return &Document{
		Expressions: append(ret1, ret2...),
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

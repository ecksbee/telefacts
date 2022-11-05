package hydratables

import (
	"encoding/xml"
	"sync"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
	"github.com/antchfx/xmlquery"
)

type Document struct {
	NamespaceMap  map[string]xml.Name
	ContextRefMap map[string]bool
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
	nret := make(map[string]xml.Name)
	cret := make(map[string]bool)
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
			name := attr.FindXpathAttr(nonNumeric.Attr, "name")
			contextRef := attr.FindXpathAttr(nonNumeric.Attr, "contextRef")
			xmlName, err := np.ProvideXmlName(name.Value)
			if err != nil {
				goErr = err
				return
			}
			lock1.Lock()
			nret[name.Value] = *xmlName
			cret[contextRef.Value] = true
			lock1.Unlock()
		}(nnonNumeric)
	}
	wg1.Wait()
	if goErr != nil {
		return nil, goErr
	}
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
			name := attr.FindXpathAttr(nonFraction.Attr, "name")
			contextRef := attr.FindXpathAttr(nonFraction.Attr, "contextRef")
			xmlName, err := np.ProvideXmlName(name.Value)
			if err != nil {
				goErr2 = err
				return
			}
			lock2.Lock()
			nret[name.Value] = *xmlName
			cret[contextRef.Value] = true
			lock2.Unlock()
		}(nnonFraction)
	}
	wg2.Wait()
	if goErr2 != nil {
		return nil, goErr2
	}
	return &Document{
		NamespaceMap:  nret,
		ContextRefMap: cret,
	}, nil
}

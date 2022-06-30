package serializables

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"strings"
	"sync"

	"ecksbee.com/telefacts/internal/actions"
	"ecksbee.com/telefacts/pkg/attr"
	"github.com/antchfx/xmlquery"
	"github.com/beevik/etree"
)

var INDENT bool

type Document struct {
	Bytes                 []byte
	Root                  *xmlquery.Node
	Html                  *xmlquery.Node
	Excludes              []*xmlquery.Node
	SchemaRefs            []*xmlquery.Node
	Contexts              []*xmlquery.Node
	Units                 []*xmlquery.Node
	NonFractions          []*xmlquery.Node
	NonNumerics           []*xmlquery.Node
	factMap               map[string](*xmlquery.Node)
	footnoteRelationships []*xmlquery.Node
}

func DecodeIxbrlFile(xmlData []byte) *Document {
	doc, err := xmlquery.Parse(bytes.NewReader(xmlData))
	if err != nil {
		fmt.Printf("Error: " + err.Error())
		return nil
	}
	var html *xmlquery.Node
	var goErr error
	htmlDone := make(chan bool)
	go func() {
		defer func() { htmlDone <- true }()
		html, goErr = xmlquery.Query(doc, "//*[local-name()='html']")
		if html == nil {
			fmt.Println("Error: missing html")
		}
		if goErr != nil {
			html = nil
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var schemaRefs []*xmlquery.Node
	schemaRefsDone := make(chan bool)
	go func() {
		defer func() { schemaRefsDone <- true }()
		schemaRefs, goErr = xmlquery.QueryAll(doc, "//*[local-name()='header' and namespace-uri()='"+attr.IX+"']//*[local-name()='schemaRef' and namespace-uri()='"+attr.LINK+"']")
		if goErr != nil {
			schemaRefs = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var contexts []*xmlquery.Node
	contextsDone := make(chan bool)
	go func() {
		defer func() { contextsDone <- true }()
		contexts, err = xmlquery.QueryAll(doc, "//*[local-name()='header' and namespace-uri()='"+attr.IX+"']//*[local-name()='resources' and namespace-uri()='"+attr.IX+"']//*[local-name()='context' and namespace-uri()='"+attr.XBRLI+"']")
		if goErr != nil {
			contexts = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var units []*xmlquery.Node
	unitsDone := make(chan bool)
	go func() {
		defer func() { unitsDone <- true }()
		units, err = xmlquery.QueryAll(doc, "//*[local-name()='header' and namespace-uri()='"+attr.IX+"']//*[local-name()='resources' and namespace-uri()='"+attr.IX+"']//*[local-name()='unit' and namespace-uri()='"+attr.XBRLI+"']")
		if goErr != nil {
			units = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var nonFractions []*xmlquery.Node
	nonFractionsDone := make(chan bool)
	go func() {
		defer func() { nonFractionsDone <- true }()
		nonFractions, err = xmlquery.QueryAll(doc, "//*[local-name()='header' and namespace-uri()='"+attr.IX+"']//*[local-name()='resources' and namespace-uri()='"+attr.IX+"']//*[local-name()='unit' and namespace-uri()='"+attr.XBRLI+"']")
		if goErr != nil {
			nonFractions = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var nonNumerics []*xmlquery.Node
	nonNumericsDone := make(chan bool)
	go func() {
		defer func() { nonNumericsDone <- true }()
		nonNumerics, err = xmlquery.QueryAll(doc, "//*[local-name()='nonNumeric' and namespace-uri()='"+attr.IX+"']")
		if goErr != nil {
			nonNumerics = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var excludes []*xmlquery.Node
	excludesDone := make(chan bool)
	go func() {
		defer func() { excludesDone <- true }()
		excludes, err = xmlquery.QueryAll(doc, "//*[local-name()='exclude' and namespace-uri()='"+attr.IX+"']")
		if goErr != nil {
			excludes = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	var footnoteRelationships []*xmlquery.Node
	footnoteRelationshipsDone := make(chan bool)
	go func() {
		defer func() { footnoteRelationshipsDone <- true }()
		footnoteRelationships, err = xmlquery.QueryAll(doc, "//*[local-name()='relationship' and namespace-uri()='"+attr.IX+"' and @arcrole='"+attr.FactFootnoteArcrole+"']")
		if goErr != nil {
			footnoteRelationships = make([]*xmlquery.Node, 0)
			fmt.Printf("Error: " + goErr.Error())
		}
	}()
	<-htmlDone
	<-schemaRefsDone
	<-contextsDone
	<-unitsDone
	<-nonFractionsDone
	<-nonNumericsDone
	<-excludesDone
	<-footnoteRelationshipsDone
	if html == nil {
		return nil
	}
	factMap := make(map[string](*xmlquery.Node))
	for _, nonFraction := range nonFractions {
		id := attr.FindXpathAttr(nonFraction.Attr, "id")
		if id == nil || len(id.Value) < 1 {
			continue
		}
		factMap[id.Value] = nonFraction
	}
	for _, nonNumeric := range nonNumerics {
		id := attr.FindXpathAttr(nonNumeric.Attr, "id")
		if id == nil || len(id.Value) < 1 {
			continue
		}
		factMap[id.Value] = nonNumeric
	}
	return &Document{
		Bytes:                 xmlData,
		Root:                  doc,
		Html:                  html,
		SchemaRefs:            schemaRefs,
		Contexts:              contexts,
		NonFractions:          nonFractions,
		NonNumerics:           nonNumerics,
		Excludes:              excludes,
		footnoteRelationships: footnoteRelationships,
		Units:                 units,
		factMap:               factMap,
	}
}

func (doc *Document) Extract(destination string) error {
	b, err := doc.Convert()
	if err != nil {
		return err
	}
	return actions.WriteFile(destination, b)
}

func (doc *Document) Convert() ([]byte, error) {
	np, err := attr.NewNameProvider(doc.Html.Attr)
	if err != nil {
		return nil, err
	}
	eDoc := etree.NewDocument()
	eDoc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	eDoc.CreateComment("  ecksbee.com/telefacts  ")
	xbrlName := np.ProvideName(attr.XBRLI, "xbrl")
	xbrl := eDoc.CreateElement(xbrlName)
	eDoc, err = doc.classicSchemaRef(eDoc, np)
	if err != nil {
		return nil, err
	}
	eDoc, err = doc.classicFacts(eDoc, np)
	if err != nil {
		return nil, err
	}
	xbrl.Attr = np.NsAttrs()
	if INDENT {
		eDoc.Indent(2)
	}
	var b bytes.Buffer
	_, err = eDoc.WriteTo(&b)
	return b.Bytes(), err
}

func (doc *Document) classicSchemaRef(eDoc *etree.Document, np *attr.NameProvider) (*etree.Document, error) {
	xbrlName := np.ProvideName(attr.XBRLI, "xbrl")
	xbrl := eDoc.SelectElement(xbrlName)
	if xbrl == nil {
		return nil, fmt.Errorf("no root xbrl element")
	}
	for _, schemaRef := range doc.SchemaRefs {
		schemaRefName := np.ProvideName(schemaRef.NamespaceURI, schemaRef.Data)
		curr := xbrl.CreateElement(schemaRefName)
		for _, a := range schemaRef.Attr {
			curr.CreateAttr(a.Name.Local, a.Value)
		}
	}
	return eDoc, nil
}

func (doc *Document) classicFacts(eDoc *etree.Document, np *attr.NameProvider) (*etree.Document, error) {
	xbrlName := np.ProvideName(attr.XBRLI, "xbrl")
	xbrl := eDoc.SelectElement(xbrlName)
	if xbrl == nil {
		return nil, fmt.Errorf("no root xbrl element")
	}
	var wg1 sync.WaitGroup
	wg1.Add(len(doc.NonFractions))
	for _, nnonFraction := range doc.NonFractions {
		go func(nonFraction *xmlquery.Node) {
			defer wg1.Done()
			nameAttr := attr.FindXpathAttr(nonFraction.Attr, "name")
			if nameAttr == nil {
				return
			}
			factName := np.ProvideConceptName(nameAttr.Value)
			classicFact := xbrl.CreateElement(factName)
			contextRef := attr.FindXpathAttr(nonFraction.Attr, "contextRef")
			if contextRef == nil {
				return
			} else {
				n := np.ProvideName(attr.XBRLI, contextRef.Name.Local)
				classicFact.CreateAttr(n, contextRef.Value)
			}
			idAttr := attr.FindXpathAttr(nonFraction.Attr, "id")
			if idAttr == nil {
				h := fnv.New128a()
				h.Write([]byte(nameAttr.Value + "_" + contextRef.Value))
				idVal := hex.EncodeToString(h.Sum([]byte{}))
				classicFact.CreateAttr("id", idVal)
			} else {
				classicFact.CreateAttr("id", idAttr.Value)
			}
			decimals := attr.FindXpathAttr(nonFraction.Attr, "decimals")
			if decimals != nil {
				n := np.ProvideName(attr.XBRLI, decimals.Name.Local)
				classicFact.CreateAttr(n, decimals.Value)
			}
			unitRef := attr.FindXpathAttr(nonFraction.Attr, "unitRef")
			if unitRef != nil {
				n := np.ProvideName(attr.XBRLI, unitRef.Name.Local)
				classicFact.CreateAttr(n, unitRef.Value)
			}
			// format := attr.FindXpathAttr(nonFraction.Attr, "format")
			// scale := attr.FindXpathAttr(nonFraction.Attr, "scale")
			classicFact.CreateText(nonFraction.InnerText())
		}(nnonFraction)
	}
	wg1.Wait()
	var wg2 sync.WaitGroup
	wg2.Add(len(doc.NonNumerics))
	for _, nnonNumeric := range doc.NonNumerics {
		go func(nonNumeric *xmlquery.Node) {
			defer wg2.Done()
			nameAttr := attr.FindXpathAttr(nonNumeric.Attr, "name")
			if nameAttr == nil {
				return
			}
			factName := np.ProvideConceptName(nameAttr.Value)
			classicFact := xbrl.CreateElement(factName)
			contextRef := attr.FindXpathAttr(nonNumeric.Attr, "contextRef")
			if contextRef == nil {
				return
			} else {
				n := np.ProvideName(attr.XBRLI, contextRef.Name.Local)
				classicFact.CreateAttr(n, contextRef.Value)
			}
			idAttr := attr.FindXpathAttr(nonNumeric.Attr, "id")
			if idAttr == nil {
				h := fnv.New128a()
				h.Write([]byte(nameAttr.Value + "_" + contextRef.Value))
				idVal := hex.EncodeToString(h.Sum([]byte{}))
				classicFact.CreateAttr("id", idVal)
			} else {
				classicFact.CreateAttr("id", idAttr.Value)
			}
			// format := attr.FindXpathAttr(nonFraction.Attr, "format")
			doc.completeTextNode(classicFact, nonNumeric, np)
		}(nnonNumeric)
	}
	wg2.Wait()
	for _, unit := range doc.Units {
		classicUnit := etree.NewDocument()
		err := classicUnit.ReadFromString(np.ProvideOutputXml(unit, true))
		if err != nil {
			return nil, err
		}
		xbrl.AddChild(classicUnit.Root())
	}
	for _, context := range doc.Contexts {
		classicContext := etree.NewDocument()
		err := classicContext.ReadFromString(np.ProvideOutputXml(context, true))
		if err != nil {
			return nil, err
		}
		xbrl.AddChild(classicContext.Root())
	}
	classicFootnoteLink, err := doc.classicFootnoteLink(np)
	if err != nil {
		return nil, err
	}
	xbrl.AddChild(classicFootnoteLink)
	return eDoc, nil
}

func (doc *Document) completeTextNode(classicFact *etree.Element, nonNumeric *xmlquery.Node, np *attr.NameProvider) {
	go func() {
		acc := ""
		sourceNode := nonNumeric
		for {
			if sourceNode != nil {
				acc = acc + np.ProvideOutputXml(sourceNode, false)
			}
			continuedAt := attr.FindXpathAttr(sourceNode.Attr, "continuedAt")
			if continuedAt == nil {
				break
			}
			sourceNode = doc.findContinuation(continuedAt)
			if sourceNode == nil {
				break
			}
		}
		textNode := etree.NewDocument()
		err := textNode.ReadFromString(acc)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
		if textNode.Root() == nil {
			classicFact.CreateText(acc)
		} else {
			excluded := exclude(textNode.Root())
			if excluded == nil {
				fmt.Println("Error: nil text node")
				return
			}
			stripped := stripIx(excluded)
			if stripped == nil {
				fmt.Println("Error: nil resultant text node")
				return
			}
			classicFact.AddChild(*stripped)
		}
	}()
}

func (doc *Document) findContinuation(continuedAt *xmlquery.Attr) *xmlquery.Node {
	continuation, _ := xmlquery.Query(doc.Root, "//*[local-name()='continuation' and namespace-uri()='"+attr.IX+"' and @id='"+continuedAt.Value+"']")
	return continuation
}

func (doc *Document) classicFootnoteLink(np *attr.NameProvider) (*etree.Element, error) {
	footnoteLinkName := np.ProvideName(attr.LINK, "footnoteLink")
	link := etree.NewElement(footnoteLinkName)
	roleName := np.ProvideName(attr.XLINK, "role")
	linkType := np.ProvideName(attr.XLINK, "type")
	link.CreateAttr(roleName, attr.ROLELINK)
	link.CreateAttr(linkType, "extended")
	for _, footnoteRelationship := range doc.footnoteRelationships {
		toRefs := attr.FindXpathAttr(footnoteRelationship.Attr, "toRefs")
		if toRefs == nil || len(strings.TrimSpace(toRefs.Value)) < 1 {
			continue
		}
		toIds := strings.Split(toRefs.Value, " ")
		fromRefs := attr.FindXpathAttr(footnoteRelationship.Attr, "fromRefs")
		if fromRefs == nil || len(strings.TrimSpace(fromRefs.Value)) < 1 {
			continue
		}
		fromIds := strings.Split(fromRefs.Value, " ")
		locs := make([]*etree.Element, 0)
		arcs := make([]*etree.Element, 0)
		for _, toId := range toIds {
			footnote, _ := xmlquery.Query(doc.Root, "//*[local-name()='footnote' and namespace-uri()='"+attr.IX+"' and @id='"+toId+"']")
			if footnote == nil {
				continue
			}
			for _, fromId := range fromIds {
				if _, found := doc.factMap[fromId]; !found {
					continue
				}
				footnoteArcName := np.ProvideName(attr.LINK, "footnoteArc")
				arc := etree.NewElement(footnoteArcName)
				arceRoleName := np.ProvideName(attr.XLINK, "arcrole")
				arc.CreateAttr(arceRoleName, attr.ROLELINK)
				arc.CreateAttr(linkType, attr.FactFootnoteArcrole)
				toName := np.ProvideName(attr.XLINK, "to")
				arc.CreateAttr(toName, toId)
				fromName := np.ProvideName(attr.XLINK, "from")
				arc.CreateAttr(fromName, fromId)
				arcs = append(arcs, arc)
				locName := np.ProvideName(attr.LINK, "loc")
				loc := etree.NewElement(locName)
				hrefName := np.ProvideName(attr.XLINK, "href")
				loc.CreateAttr(hrefName, "#"+fromId)
				loc.CreateAttr(linkType, "locator")
				labelName := np.ProvideName(attr.XLINK, "label")
				loc.CreateAttr(labelName, fromId)
				locs = append(locs, loc)
			}
			if len(locs) > 0 && len(arcs) > 0 {
				for _, loc := range locs {
					link.AddChild(loc)
				}
				for _, arc := range arcs {
					link.AddChild(arc)
				}
				classicFootnote := link.CreateElement("footnote")
				classicFootnote.Space = "link"
				doc.completeTextNode(classicFootnote, footnote, np)
				for _, fattr := range footnote.Attr {
					classicFootnote.CreateAttr(fattr.Name.Space+":"+fattr.Name.Local, fattr.Value)
				}
				classicFootnote.CreateAttr(roleName, attr.ROLEFOOTNOTE)
				classicFootnote.CreateAttr(linkType, "resource")
				link.AddChild(classicFootnote)
			}
		}
	}
	return link, nil
}

func exclude(elem *etree.Element) *etree.Element {
	ret := elem.Copy()
	if len(elem.Child) > 0 {
		if elem.Tag == "exclude" {
			return nil
		} else {
			oldChild := ret.Child
			newChild := make([]etree.Token, 0)
			elem.Child = []etree.Token{}
			for _, i := range oldChild {
				child, ok := i.(*etree.Element)
				if !ok {
					newChild = append(newChild, i)
					continue
				} else {
					excluded := exclude(child)
					if excluded == nil {
						continue
					}
					newChild = append(newChild, excluded)
				}
			}
			ret.Child = newChild
			return ret
		}
	}
	if elem.Tag == "exclude" {
		return nil
	} else {
		return ret
	}
}

func stripIx(elem *etree.Element) *etree.Token {
	ret := elem.Copy()
	if len(elem.Child) > 0 {
		oldChild := ret.Child
		newChild := []etree.Token{}
		for c := 0; c < len(oldChild); c++ {
			i := oldChild[c]
			child, ok := i.(*etree.Element)
			if !ok {
				newChild = append(newChild, i)
			} else {
				if blacklist(child.Tag) {
					temp := oldChild[:c]
					temp = append(temp, child.Child...)
					temp = append(temp, oldChild[c:]...)
					oldChild = temp
					c--
					continue
				} else {
					stripped := stripIx(child)
					newChild = append(newChild, *stripped)
				}
			}
		}
		ret.Child = newChild
		var token etree.Token
		token = etree.Token(ret)
		return &token
	}
	var token etree.Token
	token = etree.Token(ret)
	return &token
}

func blacklist(tag string) bool {
	list := []string{"footnote", "nonNumeric", "nonFraction"}
	for _, i := range list {
		if i == tag {
			return true
		}
	}
	return false
}

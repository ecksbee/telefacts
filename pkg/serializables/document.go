package serializables

import (
	"bytes"
	"fmt"

	"ecksbee.com/telefacts/pkg/attr"
	"github.com/antchfx/xmlquery"
	// parsexml "github.com/jbowtie/gokogiri/xml"
)

var INDENT bool
var XSLT string

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
		nonFractions, err = xmlquery.QueryAll(doc, "//*[local-name()='nonFraction' and namespace-uri()='"+attr.IX+"']")
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

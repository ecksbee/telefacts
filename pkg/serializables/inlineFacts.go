package serializables

import (
	"bytes"
	"encoding/xml"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html/charset"
)

const hiddenNonfractionXPath = "//html/body//*[local-name()='header' and namespace-uri()='" + attr.IX + "']//*[local-name()='hidden' and namespace-uri()='" + attr.IX + "']//*[local-name()='nonFraction' and namespace-uri()='" + attr.IX + "']"
const renderedNonfractionXPath = "//html/body//*[local-name()='nonFraction' and namespace-uri()='" + attr.IX + "' and not(ancestor::*[local-name()='hidden' and namespace-uri()='" + attr.IX + "'])]"
const hiddenNonnumericXPath = "//html/body//*[local-name()='header' and namespace-uri()='" + attr.IX + "']//*[local-name()='hidden' and namespace-uri()='" + attr.IX + "']//*[local-name()='nonNumeric' and namespace-uri()='" + attr.IX + "']"
const renderedNonnumericXPath = "//html/body//*[local-name()='nonNumeric' and namespace-uri()='" + attr.IX + "' and not(ancestor::*[local-name()='hidden' and namespace-uri()='" + attr.IX + "'])]"
const hiddenFootnoteXPath = "//html/body//*[local-name()='header' and namespace-uri()='" + attr.IX + "']//*[local-name()='hidden' and namespace-uri()='" + attr.IX + "']//*[local-name()='footnote' and namespace-uri()='" + attr.IX + "']"
const renderedFootnoteXPath = "//html/body//*[local-name()='footnote' and namespace-uri()='" + attr.IX + "' and not(ancestor::*[local-name()='hidden' and namespace-uri()='" + attr.IX + "'])]"
const continuationXPath = "//html/body//*[local-name()='continuation' and namespace-uri()='" + attr.IX + "' and not(ancestor::*[local-name()='hidden' and namespace-uri()='" + attr.IX + "'])]"

func DecodeInlineFacts(byteArray []byte) (*IxHiddenFacts, *IxRenderedFacts, error) {
	ixHidden := IxHiddenFacts{
		Nonfractions: make([]IxbrlNonfraction, 0),
		Nonnumerics:  make([]IxbrlNonnumeric, 0),
		Footnotes:    make([]IxbrlFootnote, 0),
	}
	ixRenderedFacts := IxRenderedFacts{
		Nonfractions:  make([]IxbrlNonfraction, 0),
		Nonnumerics:   make([]IxbrlNonnumeric, 0),
		Footnotes:     make([]IxbrlFootnote, 0),
		Continuations: make([]IxbrlContinuation, 0),
	}
	hiddenNonfractions, err := decodeHiddenNonfractions(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixHidden.Nonfractions = hiddenNonfractions
	hiddenNonnumerics, err := decodeHiddenNonnumerics(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixHidden.Nonnumerics = hiddenNonnumerics
	hiddenFootnotes, err := decodeHiddenFootnotes(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixHidden.Footnotes = hiddenFootnotes

	renderedNonfractions, err := decodeRenderedNonfractions(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixRenderedFacts.Nonfractions = renderedNonfractions
	renderedNonnumerics, err := decodeRenderedNonnumerics(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixRenderedFacts.Nonnumerics = renderedNonnumerics
	renderedFootnotes, err := decodeRenderedFootnotes(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixRenderedFacts.Footnotes = renderedFootnotes
	continuations, err := decodeContinuations(byteArray)
	if err != nil {
		return nil, nil, err
	}
	ixRenderedFacts.Continuations = continuations
	return &ixHidden, &ixRenderedFacts, nil
}

func decodeHiddenNonfractions(byteArray []byte) ([]IxbrlNonfraction, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, hiddenNonfractionXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlNonfraction, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlNonfraction{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeRenderedNonfractions(byteArray []byte) ([]IxbrlNonfraction, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, renderedNonfractionXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlNonfraction, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlNonfraction{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeHiddenNonnumerics(byteArray []byte) ([]IxbrlNonnumeric, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, hiddenNonnumericXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlNonnumeric, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlNonnumeric{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeRenderedNonnumerics(byteArray []byte) ([]IxbrlNonnumeric, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, renderedNonnumericXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlNonnumeric, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlNonnumeric{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeHiddenFootnotes(byteArray []byte) ([]IxbrlFootnote, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, hiddenFootnoteXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlFootnote, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlFootnote{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeRenderedFootnotes(byteArray []byte) ([]IxbrlFootnote, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, renderedFootnoteXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlFootnote, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlFootnote{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeContinuations(byteArray []byte) ([]IxbrlContinuation, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, continuationXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]IxbrlContinuation, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := IxbrlContinuation{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

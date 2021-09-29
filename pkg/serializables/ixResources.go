package serializables

import (
	"bytes"
	"encoding/xml"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html/charset"
)

const inlineContextXPath = "//html/body//*[local-name()='header' and namespace-uri()='" + attr.IX + "']//*[local-name()='resources' and namespace-uri()='" + attr.IX + "']//*[local-name()='context' and namespace-uri()='" + attr.XBRLI + "']"
const inlineUnitXPath = "//html/body//*[local-name()='header' and namespace-uri()='" + attr.IX + "']//*[local-name()='resources' and namespace-uri()='" + attr.IX + "']//*[local-name()='unit' and namespace-uri()='" + attr.XBRLI + "']"

func DecodeIxResources(byteArray []byte) (*IxResources, error) {
	decoded := IxResources{}
	contexts, err := decodeInlineContexts(byteArray)
	if err != nil {
		return nil, err
	}
	decoded.Contexts = contexts
	units, err := decodeInlineUnits(byteArray)
	if err != nil {
		return nil, err
	}
	decoded.Units = units
	return &decoded, nil
}

func decodeInlineContexts(byteArray []byte) ([]CommonContext, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, inlineContextXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]CommonContext, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := CommonContext{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

func decodeInlineUnits(byteArray []byte) ([]CommonUnit, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	list, err := xmlquery.QueryAll(doc, inlineUnitXPath)
	if err != nil {
		return nil, err
	}
	ret := make([]CommonUnit, 0, len(list))
	for _, item := range list {
		xmlData := item.OutputXML(true)
		reader := strings.NewReader(xmlData)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		decoded := CommonUnit{}
		err := decoder.Decode(&decoded)
		if err != nil {
			return nil, err
		}
		ret = append(ret, decoded)
	}
	return ret, nil
}

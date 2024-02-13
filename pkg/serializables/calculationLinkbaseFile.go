package serializables

import (
	"bytes"
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

type CalculationLinkbaseFile struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"roleRef"`
	CalculationLink []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Loc      []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"loc"`
		CalculationArc []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"calculationArc"`
	} `xml:"calculationLink"`
}

func DecodeCalculationLinkbaseFile(xmlData []byte) (*CalculationLinkbaseFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := CalculationLinkbaseFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadCalculationLinkbaseFile(filepath string) (*CalculationLinkbaseFile, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeCalculationLinkbaseFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

package serializables

import (
	"bytes"
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

type LabelLinkbaseFile struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"roleRef"`
	LabelLink []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Loc      []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"loc"`
		Label []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			CharData string     `xml:",chardata"`
		} `xml:"label"`
		LabelArc []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"labelArc"`
	} `xml:"labelLink"`
}

func DecodeLabelLinkbaseFile(xmlData []byte) (*LabelLinkbaseFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := LabelLinkbaseFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadLabelLinkbaseFile(filepath string) (*LabelLinkbaseFile, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeLabelLinkbaseFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

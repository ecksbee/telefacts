package serializables

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

type DefinitionLinkbaseFile struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"roleRef"`
	DefinitionLink []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Loc      []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"loc"`
		DefinitionArc []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"definitionArc"`
	} `xml:"definitionLink"`
}

func DecodeDefinitionLinkbaseFile(xmlData []byte) (*DefinitionLinkbaseFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := DefinitionLinkbaseFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadDefinitionLinkbaseFile(filepath string) (*DefinitionLinkbaseFile, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeDefinitionLinkbaseFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

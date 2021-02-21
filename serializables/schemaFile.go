package serializables

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

type SchemaFile struct {
	XMLName  xml.Name   `xml:"schema"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	Include  []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"include"`
	Import []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"import"`
	Annotation []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Appinfo  []struct {
			XMLName     xml.Name
			XMLAttrs    []xml.Attr `xml:",any,attr"`
			LinkbaseRef []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
			} `xml:"linkbaseRef"`
			RoleType []struct {
				XMLName    xml.Name
				XMLAttrs   []xml.Attr `xml:",any,attr"`
				Definition []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					CharData string     `xml:",chardata"`
				} `xml:"definition"`
				UsedOn []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					CharData string     `xml:",chardata"`
				} `xml:"usedOn"`
			} `xml:"roleType"`
		} `xml:"appinfo"`
	} `xml:"annotation"`
	Element []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		CharData string     `xml:",chardata"`
	} `xml:"element"`
}

func DecodeSchemaFile(xmlData []byte) (*SchemaFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := SchemaFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadSchemaFile(filepath string) (*SchemaFile, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeSchemaFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

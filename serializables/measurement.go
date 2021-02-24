package serializables

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

type UnitTypeRegistry struct {
	XMLName  xml.Name   `xml:"utr"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	Units    []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Unit     []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			UnitID   []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"unitId"`
			UnitName []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"unitName"`
			NSUnit []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"nsUnit"`
			ItemType []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"itemType"`
			ItemTypeDate []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"itemTypeDate"`
			Symbol []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"symbol"`
			Definition []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"definition"`
			BaseStandard []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"baseStandard"`
			ConversionPresentation []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				XMLInner string     `xml:",innerxml"`
			} `xml:"conversionPresentation"`
			ConversionContent []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				XMLInner string     `xml:",innerxml"`
			} `xml:"conversionContent"`
			Status []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"status"`
			VersionDate []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"versionDate"`
		} `xml:"unit"`
	} `xml:"units"`
}

func DecodeUnitTypeRegistry(xmlData []byte) (*UnitTypeRegistry, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := UnitTypeRegistry{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadUnitTypeRegistry(filepath string) (*UnitTypeRegistry, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeUnitTypeRegistry(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

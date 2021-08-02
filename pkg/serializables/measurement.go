package serializables

import (
	"bytes"
	"encoding/xml"

	"ecksbee.com/telefacts/pkg/attr"
	"golang.org/x/net/html/charset"
)

type UnitTypeRegistryFile struct {
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
			NSItemType []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"nsItemType"`
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
			NumeratorItemType []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"numeratorItemType"`
			NSNumeratorItemType []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"nsNumeratorItemType"`
			DenominatorItemType []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"denominatorItemType"`
			NSDenominatorItemType []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"nsDenominatorItemType"`
		} `xml:"unit"`
	} `xml:"units"`
}

func DecodeUnitTypeRegistry(xmlData []byte) (*UnitTypeRegistryFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := UnitTypeRegistryFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func DiscoverUnitTypeRegistry() (*UnitTypeRegistryFile, error) {
	data, err := DiscoverGlobalFile(attr.UTR)
	if err != nil {
		return nil, err
	}
	return DecodeUnitTypeRegistry(data)
}

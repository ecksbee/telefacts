package serializables

import (
	"bytes"
	"encoding/xml"
	"os"

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
		XMLName    xml.Name
		XMLAttrs   []xml.Attr `xml:",any,attr"`
		Definition []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			CharData string     `xml:",chardata"`
		} `xml:"definition"`
		Appinfo []struct {
			XMLName    xml.Name
			XMLAttrs   []xml.Attr `xml:",any,attr"`
			Definition []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"definition"`
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
			EmbeddedLinkbase []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				RoleRef  []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
				} `xml:"roleRef"`
				PresentationLink []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					Loc      []struct {
						XMLName  xml.Name
						XMLAttrs []xml.Attr `xml:",any,attr"`
					} `xml:"loc"`
					PresentationArc []struct {
						XMLName  xml.Name
						XMLAttrs []xml.Attr `xml:",any,attr"`
					} `xml:"presentationArc"`
				} `xml:"presentationLink"`
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
			} `xml:"linkbase"`
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
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeSchemaFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
